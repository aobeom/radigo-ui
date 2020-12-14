# radigo-ui
Desktop Version For [radigo](https://github.com/aobeom/radigo)

## Usage

1. [Font](https://fonts.google.com/) Both Regular & Bold
2. [fyne](https://github.com/fyne-io/fyne)
3. Bundle Font
```go
go get fyne.io/fyne/cmd/fyne
fyne bundle -package theme your-font-regular.ttf > theme/bundle.go
fyne bundle -append your-font-bold.ttf >> theme/bundle.go
```
4. Generate theme
```go
type MyTheme struct{}

func (MyTheme) TextFont() fyne.Resource { return resourceYourFontRegularTtf }
func (MyTheme) TextBoldFont() fyne.Resource { return resourceYourFontBoldTtf }

func (MyTheme) BackgroundColor() color.Color      { return theme.LightTheme().BackgroundColor() }
func (MyTheme) ButtonColor() color.Color          { return theme.LightTheme().ButtonColor() }
func (MyTheme) DisabledButtonColor() color.Color  { return theme.LightTheme().DisabledButtonColor() }
func (MyTheme) IconColor() color.Color            { return theme.LightTheme().IconColor() }
func (MyTheme) DisabledIconColor() color.Color    { return theme.LightTheme().DisabledIconColor() }
func (MyTheme) HyperlinkColor() color.Color       { return theme.LightTheme().HyperlinkColor() }
func (MyTheme) TextColor() color.Color            { return theme.LightTheme().TextColor() }
func (MyTheme) DisabledTextColor() color.Color    { return theme.LightTheme().DisabledTextColor() }
func (MyTheme) HoverColor() color.Color           { return theme.LightTheme().HoverColor() }
func (MyTheme) PlaceHolderColor() color.Color     { return theme.LightTheme().PlaceHolderColor() }
func (MyTheme) PrimaryColor() color.Color         { return theme.LightTheme().PrimaryColor() }
func (MyTheme) FocusColor() color.Color           { return theme.LightTheme().FocusColor() }
func (MyTheme) ScrollBarColor() color.Color       { return theme.LightTheme().ScrollBarColor() }
func (MyTheme) ShadowColor() color.Color          { return theme.LightTheme().ShadowColor() }
func (MyTheme) TextSize() int                     { return theme.LightTheme().TextSize() }
func (MyTheme) TextItalicFont() fyne.Resource     { return theme.LightTheme().TextItalicFont() }
func (MyTheme) TextBoldItalicFont() fyne.Resource { return theme.LightTheme().TextBoldItalicFont() }
func (MyTheme) TextMonospaceFont() fyne.Resource  { return theme.LightTheme().TextMonospaceFont() }
func (MyTheme) Padding() int                      { return theme.LightTheme().Padding() }
func (MyTheme) IconInlineSize() int               { return theme.LightTheme().IconInlineSize() }
func (MyTheme) ScrollBarSize() int                { return theme.LightTheme().ScrollBarSize() }
func (MyTheme) ScrollBarSmallSize() int           { return theme.LightTheme().ScrollBarSmallSize() }
```
5. Package
```go
fyne package -icon radigo-ui.png
```
---
## Reference
[fyne-font-example](https://github.com/lusingander/fyne-font-example)