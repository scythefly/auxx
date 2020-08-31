package image

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"

	"code.google.com/p/graphics-go/graphics"
	"github.com/spf13/cobra"
)

var option struct {
	x int
}

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "img",
		Short: "Run image examples",
		RunE:  runImage,
	}
	flags := cmd.PersistentFlags()
	flags.IntVarP(&option.x, "max-x", "x", 50, "max x pixes")

	return cmd
}

func runImage(_ *cobra.Command, args []string) error {
	for idx := range args {
		imgTransfer(args[idx])
	}
	return nil
}

func imgTransfer(path string) {
	ff, _ := ioutil.ReadFile("./47.44.png")
	buf := bytes.NewBuffer(ff)
	m, _, err := image.Decode(buf)
	if err != nil {
		fmt.Printf("decode return err: %s\n", err.Error())
		return
	}
	asciiImage(m)
}
func rectImage(m image.Image, newdx int) *image.RGBA {
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	newRgba := image.NewRGBA(image.Rect(0, 0, newdx*3/2, newdx*dy/dx))
	graphics.Scale(newRgba, m)
	return newRgba
}

func asciiImage(m image.Image) {
	fmt.Printf("dx: %d, dy: %d\n", m.Bounds().Dx(), m.Bounds().Dy())
	if m.Bounds().Dx() > option.x {
		m = rectImage(m, option.x)
	}
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	arr := []string{".", "N", "H", "Q", "$", "O", "C", "?", "7", ">", "!", ":", "–", ";", "."}

	ff, err := os.Create("./out.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ff.Close()
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			colorRgb := m.At(j, i)
			_, g, _, _ := colorRgb.RGBA()
			avg := uint8(g >> 8)
			num := avg / 18
			ff.WriteString(arr[num])
			if j == dx-1 {
				ff.WriteString("\n")
			}
		}
	}

}
