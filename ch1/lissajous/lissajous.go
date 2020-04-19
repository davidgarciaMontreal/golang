package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"reflect"
)

var palettes = []color.Color{}

func init() {
	// color := color.RGBA{0x00, 0xff, 0x00, 0xff}
	// v := reflect.ValueOf(&color).Elem()
	// typeOfS := v.Type()
	// Fprintf(w io.Writer, format string, a ...interface{})
	// fmt.Fprintf(os.Stderr, "i=%d,r=%d,cindex=%d\n", i, rvalue, cindex)
	// mv.SetMapIndex(reflect.ValueOf("a"), reflect.ValueOf("11"))
	// see https://stackoverflow.com/questions/49147706/reflect-accessed-map-does-not-provide-a-modifiable-value/49147910#49147910
	// v.Field(0).Set(reflect.ValueOf(0xff))
	// mv := reflect.ValueOf(&color).Elem()
	// mv.Field(0).Set(reflect.ValueOf(uint8(0xff)))
	// fmt.Fprintf(os.Stderr, "Field: %v\n", mv)
	// v.Field(0).Set(reflect.ValueOf(uint8(0xff)))
	// for i := 0; i < v.NumField(); i++ {
	// 	fmt.Fprintf(os.Stderr, "Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
	// }
	palettes = append(palettes, color.White)
	palettes = append(palettes, color.Black)
	c := color.RGBA{0x00, 0xff, 0x00, 0xff}
	for i := 2; i < numberOfElements; i++ {
		v := reflect.ValueOf(&c).Elem()                                  // create new elem
		colorvalue := uint8(rand.Intn(255))                              // chose a ranmdom RGBA value
		colorfieldvalue := rand.Intn(3)                                  // chose a random field to be updated: R, G, or B
		v.Field(colorfieldvalue).Set(reflect.ValueOf(uint8(colorvalue))) // update the color.RGBA value
		// typeOfS := v.Type()
		// fmt.Fprintf(os.Stderr, "Field: %s RGBValue: %v ", typeOfS.Field(colorfieldvalue).Name, v.Field(colorfieldvalue).Interface())
		// fmt.Fprintf(os.Stderr, "Pointer: %p \n", &v)
		palettes = append(palettes, c)
		c = color.RGBA{0x00, 0xff, 0x00, 0xff}
	}
	// for index, element := range palettes {
	// 	fmt.Fprintf(os.Stderr, "index=%d value=%v p=%p\n", index, element, &palettes[index])
	// }
}

// var palette = []color.Color{color.White,
// 	color.Black,
// 	color.RGBA{0x00, 0xff, 0x00, 0xff}} // composite literals

const (
	whiteIndex       = 0
	blackIndex       = 1
	numberOfElements = 100
)

func main() {
	lissajous(os.Stdout)
}
func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of coplete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0        // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes} //struct git.GIF
	phase := 0.0                        // phase difference
	// var cindex uint8
	// max, min := 99, 1
	// rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palettes)
		// rvalue := rand.Intn(max)     // need to remove whiteIndex
		// cindex = uint8(rvalue + min) //
		// Fprintf(w io.Writer, format string, a ...interface{})
		// fmt.Fprintf(os.Stderr, "i=%d,r=%d,cindex=%d\n", i, rvalue, cindex)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(i+2))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ingoring econding errors :'(
}
