package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"

	"radigo-ui/radio"
	"radigo-ui/theme"
	"radigo-ui/utils"
)

var stationMaps = map[string]string{
	"TBSラジオ":       "TBS",
	"文化放送":         "QRR",
	"ニッポン放送":       "LFR",
	"InterFM897":   "INT",
	"TOKYO FM":     "FMT",
	"J-WAVE":       "FMJ",
	"ラジオ日本":        "JORF",
	"NHKラジオ第1（東京）": "JOAK",
	"ラジオNIKKEI第1":  "RN1",
	"ラジオNIKKEI第2":  "RN2",
	"NHKラジオ第2":     "JOAB",
	"NHK-FM（東京）":   "JOAK-FM",
}

var stationsList = []string{
	"TBSラジオ",
	"文化放送",
	"ニッポン放送",
	"InterFM897",
	"TOKYO FM",
	"J-WAVE",
	"ラジオ日本",
	"NHKラジオ第1（東京）",
	"ラジオNIKKEI第1",
	"ラジオNIKKEI第2",
	"NHKラジオ第2",
	"NHK-FM（東京）",
}

// StationChecker 检查电台编号
func StationChecker(i string) bool {
	checkStation := radio.RegionXML("id", i)
	if checkStation == "" {
		return false
	}
	return true
}

// DLServer 下载控制器
type DLServer struct {
	WG    sync.WaitGroup
	Gonum chan string
}

// rThread 并发下载
func rThread(url string, ch chan []byte, dl *DLServer) {
	headers := utils.MiniHeaders{
		"User-Agent": radio.UserAgent,
	}
	res := utils.Minireq.Get(url, headers, nil)
	dl.WG.Done()
	<-dl.Gonum
	ch <- res.RawData()
}

// rEngine 下载器
func rEngine(thread bool, urls []string, savePath string, progressBar *widget.ProgressBar) {
	aacFile, _ := os.OpenFile(savePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	total := float64(len(urls))
	part, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", 100.0/total), 64)
	if thread {
		var thread int
		dl := new(DLServer)
		if total < 16 {
			thread = int(total)
		} else {
			thread = 16
		}
		var num int

		ch := make([]chan []byte, 8192)
		dl.Gonum = make(chan string, thread)
		dl.WG.Add(len(urls))

		for i, url := range urls {
			dl.Gonum <- url
			ch[i] = make(chan []byte, 8192)
			go rThread(url, ch[i], dl)
			progress := float64(i) * part
			num := progress / 10
			progressBar.SetValue(num)
		}
		for _, d := range ch {
			if num == int(total) {
				break
			}
			tmp, _ := <-d
			offset, _ := aacFile.Seek(0, os.SEEK_END)
			aacFile.WriteAt(tmp, offset)
			num++
		}
		dl.WG.Wait()
	} else {
		for i, url := range urls {
			headers := utils.MiniHeaders{
				"User-Agent": radio.UserAgent,
			}
			res := utils.Minireq.Get(url, headers, nil)
			offset, _ := aacFile.Seek(0, os.SEEK_END)
			aacFile.WriteAt(res.RawData(), offset)
			progress := float64(i) * part
			num := progress / 10
			progressBar.SetValue(num)
		}
	}
	defer aacFile.Close()
}

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(&theme.MyTheme{})
	myWin := myApp.NewWindow("Radiko")
	myWin.Resize(fyne.NewSize(320, 250))
	myWin.SetFixedSize(true)
	myWin.CenterOnScreen()

	var stationCode string
	stationBox := widget.NewSelect(stationsList, func(value string) {
		stationCode = stationMaps[value]
	})
	startAtEntry := widget.NewEntry()
	endAtEntry := widget.NewEntry()
	proxyEntry := widget.NewEntry()
	progressBar := widget.NewProgressBar()
	dlInfo := widget.NewLabel("Please Enter the Radio Info")

	stationBox.PlaceHolder = "Station"
	startAtEntry.SetPlaceHolder("Start: 20200101060000")
	endAtEntry.SetPlaceHolder("End: 20200101080000")
	proxyEntry.SetPlaceHolder("Socks5 Proxy")

	dlBtn := widget.NewButton("Download", func() {
		radioData := new(radio.Params)
		if StationChecker(stationCode) {
			startText := startAtEntry.Text
			endText := endAtEntry.Text
			_, startErr := time.Parse("20060102150405", startText)
			_, endErr := time.Parse("20060102150405", endText)
			if startErr != nil || endErr != nil {
				dlInfo.SetText("Wrong time format: 20200101060000")
			} else {
				if proxyEntry.Text != "" {
					utils.Minireq.Proxy(proxyEntry.Text)
				}
				radioData.StationID = stationCode
				radioData.StartAt = startText
				radioData.Ft = startText
				radioData.EndAt = endText
				radioData.To = endText
				radioData.L = "15"
				radioData.RType = "b"

				dlInfo.SetText("Checking Your IP...")
				result, ok := radio.IPCheck()
				if ok {
					saveName := fmt.Sprintf("%s.%s.%s.aac", radioData.StationID, radioData.StartAt, radioData.EndAt)
					workDir := utils.FileSuite.LocalPath(radio.Debug)
					savePath := filepath.Join(workDir, saveName)
					// 1.获取 JS Key
					// dlInfo.SetText("Get Key...")
					// authkey := radikoJSKey(client)
					authkey := "bcd151073c03b352e1ef2fd66c32209da9ca0afa"
					// 2.获取认证信息
					dlInfo.SetText("Get Auth Info...")
					token, offset, length := radio.Auth1()
					partialkey := radio.EncodeKey(authkey, offset, length)
					// 3.获取播放地区
					dlInfo.SetText("Get Radio Area...")
					area := radio.Auth2(token, partialkey)
					// 4.开始下载
					dlInfo.SetText("Downloading...")
					aacURLs := radio.GetAACList(token, area, radioData)
					rEngine(true, aacURLs, savePath, progressBar)
					dlInfo.SetText("Finished")
				} else {
					dlInfo.SetText("IP Forbidden: " + result)
				}
			}
		} else {
			dlInfo.SetText("Station does not exist.")
		}
	})

	content := widget.NewVBox(
		stationBox,
		startAtEntry,
		endAtEntry,
		proxyEntry,
		progressBar,
		dlInfo,
		dlBtn,
	)

	myWin.SetContent(content)
	myWin.ShowAndRun()
}
