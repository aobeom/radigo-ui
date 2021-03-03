# radigo-ui
Desktop Version For [radigo](https://github.com/aobeom/radigo)

## Usage

1. [Font](https://fonts.google.com/)
2. [fyne](https://github.com/fyne-io/fyne)
3. Bundle Font
```go
go get fyne.io/fyne/cmd/fyne
fyne bundle -package theme your-font-regular.ttf > theme/bundle.go
fyne bundle -append radigo-ui.png >> theme/bundle.go
```
4. Generate theme
```go
type MyTheme struct{}

func (m MyTheme) Font(s fyne.TextStyle) fyne.Resource {
	return resourceYourFontRegularTtf
}
func (*MyTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(n, v)
}

func (*MyTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (*MyTheme) Size(n fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(n)
}
```
5. Package
```go
fyne package -icon radigo-ui.png
```
---
## Reference
[fyne-font-example](https://github.com/lusingander/fyne-font-example)