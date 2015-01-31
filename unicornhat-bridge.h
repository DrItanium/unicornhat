#ifndef _UNICORNHAT_BRIDGE_H
#define _UNICORNHAT_BRIDGE_H
enum {
	BaseAddresses_DMA_BASE = DMA_BASE,
	BaseAddresses_DMA_LEN = DMA_LEN,
	BaseAddresses_PWM_BASE		= 0x2020C000,
	BaseAddresses_PWM_LEN			= 0x28,
	BaseAddresses_CLK_BASE	    = 0x20101000,
	BaseAddresses_CLK_LEN			= 0xA8,
	BaseAddresses_GPIO_BASE		= 0x20200000,
	BaseAddresses_GPIO_LEN		= 0xB4,
};
#endif
