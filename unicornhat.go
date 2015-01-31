package unicornhat

// #include "ws2812-RPi.h"
// #include "unicornhat-bridge.h"
import "C"
import "fmt"

func InitHardware() {
	C.initHardware()
}

