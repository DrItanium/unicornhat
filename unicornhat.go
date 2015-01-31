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
	return Pixel { R: r, G: g, B: b }
	
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
