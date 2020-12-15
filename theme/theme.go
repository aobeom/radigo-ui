package theme

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/theme"
)

type MyTheme struct{}

func (MyTheme) TextFont() fyne.Resource     { return resourceMPLUSRounded1cRegularTtf }
func (MyTheme) TextBoldFont() fyne.Resource { return resourceMPLUSRounded1cBoldTtf }

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

func MyLogo() fyne.Resource { return resourceRadigoUiPng }
