// Interface functions with the pimoroni UnicornHat
package unicornhat

//#include "ws2812-RPi.h"
//
//double getDefaultBrightnessMacro() {
//	return DEFAULT_BRIGHTNESS;
//}
import "C"

type Pixel struct {
	R, G, B byte
}

// initialization of hardware
func InitHardware() {
	C.initHardware()
}

func Init(numPixels int) {
	C.init(C.int(numPixels))
}

func Initialize(numPixels uint, brightness float64) {
	SetNumLEDs(numPixels)
	SetBrightness(brightness)
	C.initHardware()
}

func StartTransfer() {
	C.startTransfer()
}

// Led updates
func Show() {
	C.show()
}

func GetBrightness() float64 {
	return float64(C.getBrightness())
}
func SetBrightness(brightness float64) bool {
	return (C.setBrightness(C.double(brightness)) == 1)
}

func ClearPWMBuffer() {
	C.clearPWMBuffer()
}

func Clear() {
	C.clear()
}

func ClearLEDBuffer() {
	C.clearLEDBuffer()
}

func GetPixelColor(pixel uint) Pixel {
	return fromNativePixel(C.getPixelColor(C.uint(pixel)))
}
func NewPixel(r, g, b byte) Pixel {
	return Pixel{R: r, G: g, B: b}

}
func Color(r, g, b byte) Pixel {
	return fromNativePixel(C.Color(C.uchar(r), C.uchar(g), C.uchar(b)))
}
func RGB2Color(r, g, b byte) Pixel {
	return fromNativePixel(C.RGB2Color(C.uchar(r), C.uchar(g), C.uchar(b)))
}

func (this Pixel) nativePixel() C.Color_t {
	var v C.Color_t
	v.r = C.uchar(this.R)
	v.g = C.uchar(this.G)
	v.b = C.uchar(this.B)
	return v
}

func fromNativePixel(pixel C.Color_t) Pixel {
	return NewPixel(byte(pixel.r), byte(pixel.g), byte(pixel.b))
}

func NumPixels() uint {
	return uint(C.numPixels())
}

func SetPixelColor(pixel uint, r, g, b byte) bool {
	return (C.setPixelColor(C.uint(pixel), C.uchar(r), C.uchar(g), C.uchar(b)) == 1)
}

func SetPixelColorType(pixel uint, color Pixel) bool {
	return (C.setPixelColorT(C.uint(pixel), color.nativePixel()) == 1)
}

func DefaultBrightness() float64 {
	return float64(C.getDefaultBrightnessMacro())
}

func Shutdown(dummy int) {
	C.shutdown(C.int(dummy))
}

var coordinates [8][8]int = [8][8]int{
	{7, 6, 5, 4, 3, 2, 1, 0},
	{8, 9, 10, 11, 12, 13, 14, 15},
	{23, 22, 21, 20, 19, 18, 17, 16},
	{24, 25, 26, 27, 28, 29, 30, 31},
	{39, 38, 37, 36, 35, 34, 33, 32},
	{40, 41, 42, 43, 44, 45, 46, 47},
	{55, 54, 53, 52, 51, 50, 49, 48},
	{56, 57, 58, 59, 60, 61, 62, 63},
}

// taken from unicorn.c
func CoordinateToPosition(x, y int) int {
	return coordinates[x][y]
}

func SetNumLEDs(count uint) {
	// since this is for the unicorn hat, I'm going to silently change this to 64
	if count > 64 {
		C.numLEDs = C.uint(64)
	} else {
		C.numLEDs = C.uint(count)
	}
}
func GetNumLEDs() uint {
	return uint(C.numLEDs)
}
