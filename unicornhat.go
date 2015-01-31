package unicornhat

// #include "ws2812-RPi.h"
// #include "unicornhat-bridge.h"
import "C"
import "fmt"

type Color struct {
	r, g, b byte
}

// initialization of hardware
func InitHardware() {
	C.initHardware()
}

func Initialize(numPixels int) {
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
func SetBrightness(brightness float64) byte {
	return byte(C.setBrightness(C.double(brightness)))
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

func GetPixelColor(pixel uint) Color {
	container := C.getPixelColor(C.uint(pixel))
	return NewColor(byte(container.r), byte(container.g), byte(container.b))
}

func NewColor(r, g, b byte) Color {
	return Color { r: r, g: g, b: b }
}
