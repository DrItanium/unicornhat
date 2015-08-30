// Interface functions with the pimoroni UnicornHat
package unicornhat

//#include "ws2811.h"
//#include "board_info.h"
//
//int getTargetFreq() {
//	return WS2811_TARGET_FREQ;
//}
//ws2811_t ledstring =
//{
//	.freq = WS2811_TARGET_FREQ,
//	.dmanum = 5,
//	.channel =
//	{
//		[0] =
//		{
//			.gpionum = 18,
//			.count = (8 * 8),
//			.invert = 0,
//			.brightness = 55,
//		}
//	}
//};
//void setBrightness(int b) {
//	ledstring.channel[0].brightness = b;
//}
//int getBrightness(void) {
// return ledstring.channel[0].brightness;
//}
//void setPixelColorRGB(int pixel, int r, int g, int b) {
// ledstring.channel[0].leds[pixel] = (r << 16) | (g << 8) | b;
//}
//
//void clearLEDBuffer(void) {
//	int i;
//	for (i = 0; i < 64; i++) {
//		setPixelColorRGB(i, 0, 0, 0);
//	}
//}
//void show(void) {
// ws2811_render(&ledstring);
//}
//int initializeUnicornhat(void) {
//	if (board_info_init() < 0) {
//		return 0;
//	}
//	if (ws2811_init(&ledstring)) {
//		return 0;
//	}
//
//	clearLEDBuffer();
//	return 1;
//}
//void shutdownUnicornhat(void) {
//	clearLEDBuffer();
// ws2811_render(&ledstring);
// ws2811_fini(&ledstring);
//}
//ws2811_led_t getPixel(int index) {
//	return ledstring.channel[0].leds[index];
//}
import "C"
import "fmt"

const (
	Width             = 8
	Height            = 8
	PixelCount        = Width * Height
	DefaultBrightness = 55
)

var pixelPos = [][]int{
	{7, 6, 5, 4, 3, 2, 1, 0},
	{8, 9, 10, 11, 12, 13, 14, 15},
	{23, 22, 21, 20, 19, 18, 17, 16},
	{24, 25, 26, 27, 28, 29, 30, 31},
	{39, 38, 37, 36, 35, 34, 33, 32},
	{40, 41, 42, 43, 44, 45, 46, 47},
	{55, 54, 53, 52, 51, 50, 49, 48},
	{56, 57, 58, 59, 60, 61, 62, 63},
}

func PixelPosition(x, y int) (int, error) {
	if x > Width || x < 0 {
		return 0, fmt.Errorf("X coordinate is out of range (%d)", x)
	} else if y > Height || y < 0 {
		return 0, fmt.Errorf("Y coordinate is out of range (%d)", y)
	} else {
		return pixelPos[x][y], nil
	}
}

type Pixel struct {
	R, G, B byte
}

func NewPixel(r, g, b byte) *Pixel {
	return &Pixel{R: r, G: g, B: b}
}

// initialization of hardware
func Initialize() error {
	if result := C.initializeUnicornhat(); result != 1 {
		return fmt.Errorf("Couldn't initalize the unicornhat!")
	} else {
		return nil
	}
}
func Shutdown() {
	C.shutdownUnicornhat()
}

// Led updates
func Show() {
	C.show()
}

func GetBrightness() byte {
	return byte(C.getBrightness())
}
func SetBrightness(brightness byte) {
	C.setBrightness(C.int(brightness))
}

func ClearLEDBuffer() {
	for i := 0; i < PixelCount; i++ {
		C.setPixelColorRGB(C.int(i), 0, 0, 0)
	}
}

func GetPixelColor(pixel int) *Pixel {
	var color C.ws2811_led_t
	color = C.getPixel(C.int(pixel))
	return NewPixel(byte((color&0x00FF0000)>>16), byte((color&0x0000FF00)>>8), byte(color&0x000000FF))
}

func SetPixelColor(pixel int, r, g, b byte) {
	C.setPixelColorRGB(C.int(pixel), C.int(r), C.int(g), C.int(b))
}
