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
	set "github.com/deckarep/golang-set"
	"github.com/spf13/cobra"
)

var option struct {
	x    int
	anti bool
}

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "img",
		Short: "Run image examples",
		RunE:  runImage,
	}
	flags := cmd.PersistentFlags()
	flags.IntVarP(&option.x, "max-x", "x", 50, "max x pixes")
	flags.BoolVarP(&option.anti, "anti-color", "o", false, "anti-color")

	return cmd
}

func runImage(_ *cobra.Command, args []string) error {
	fs := set.NewSet()
	for idx := range args {
		fs.Add(args[idx])
	}

	fs.Each(func(v interface{}) bool {
		imgTransfer(v.(string))
		return false
	})
	return nil
}

func imgTransfer(path string) {
	ff, _ := ioutil.ReadFile(path)
	buf := bytes.NewBuffer(ff)
	m, _, err := image.Decode(buf)
	if err != nil {
		fmt.Printf("decode return err: %s\n", err.Error())
		return
	}
	asciiImage(m, path+".txt")
}
func rectImage(m image.Image, newdx int) *image.RGBA {
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	newRgba := image.NewRGBA(image.Rect(0, 0, newdx*3/2, newdx*dy/dx))
	graphics.Scale(newRgba, m)
	return newRgba
}

func asciiImage(m image.Image, out string) {
	fmt.Printf("dx: %d, dy: %d\n", m.Bounds().Dx(), m.Bounds().Dy())
	if m.Bounds().Dx() > option.x {
		m = rectImage(m, option.x)
	}
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	var arr []string
	if !option.anti {
		arr = []string{".", ";", "-", ":", "!", ">", "7", "?", "C", "O", "$", "Q", "H", "N", "M"}
	} else {
		arr = []string{"M", "N", "H", "Q", "$", "O", "C", "?", "7", ">", "!", ":", "–", ";", "."}
	}

	ff, err := os.Create(out)
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
