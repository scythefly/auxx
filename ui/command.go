package ui

import (
	"fyne.io/fyne"
	application "fyne.io/fyne/app"
	"fyne.io/fyne/driver/desktop"
	"fyne.io/fyne/widget"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ui",
		Short: "Run ui examples",
		RunE:  uiRun,
	}
	return cmd
}

func uiRun(_ *cobra.Command, _ []string) error {
	// app := application.New()
	// w := app.NewWindow("Hello")
	_ = application.New()
	drv := fyne.CurrentApp().Driver().(desktop.Driver)
	if drv == nil {
		return errors.New("not in desktop model")
	}
	w := drv.CreateSplashWindow()

	hello := widget.NewLabel("Hello Scythefly")
	var sw fyne.Window
	w.SetContent(widget.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			if sw == nil {
				hello.SetText("Show")
				sw = drv.CreateSplashWindow()
				sw.SetContent(widget.NewLabelWithStyle(`func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(widget.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	w.ShowAndRun()
}`, fyne.TextAlignCenter, fyne.TextStyle{Bold: true}))

				sw.Show()
			} else {
				hello.SetText("close")
				sw.Close()
				sw = nil
			}
		}),
	))

	w.ShowAndRun()

	return nil
}
