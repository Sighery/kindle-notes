package main

import (
	"fmt"
	"time"
	"syscall"

	"github.com/holoplot/go-evdev"
)

func CurrentTimeval() (syscall.Timeval) {
	now := time.Now()
	return syscall.Timeval{
		Sec: int32(now.Unix()),
		Usec: int32(now.UnixNano() / 1000 % 1000),
	}
}

func main() {
	devs, err := evdev.ListDevicePaths()
	fmt.Println(devs)

	d, err := evdev.Open("/dev/input/event1")
	if err != nil {
		fmt.Printf("Cannot read input device")
		return
	}

	events := []evdev.InputEvent {
		evdev.InputEvent{
			Time: CurrentTimeval(),
			Type: evdev.EV_ABS,
			Code: evdev.ABS_MT_TRACKING_ID,
			Value: 0,
		},
		evdev.InputEvent{
			Time: CurrentTimeval(),
			Type: evdev.EV_ABS,
			Code: evdev.ABS_MT_POSITION_X,
			Value: 153,
		},
		evdev.InputEvent{
			Time: CurrentTimeval(),
			Type: evdev.EV_ABS,
			Code: evdev.ABS_MT_POSITION_Y,
			Value: 825,
		},
		evdev.InputEvent{
			Time: CurrentTimeval(),
			Type: evdev.EV_ABS,
			Code: evdev.ABS_MT_PRESSURE,
			Value: 100,
		},
		evdev.InputEvent{
			Time: CurrentTimeval(),
			Type: evdev.EV_SYN,
			Code: evdev.SYN_REPORT,
			Value: 0,
		},
		evdev.InputEvent{
			Time: CurrentTimeval(),
			Type: evdev.EV_ABS,
			Code: evdev.ABS_MT_POSITION_Y,
			Value: 829,
		},
		evdev.InputEvent{
			Time: CurrentTimeval(),
			Type: evdev.EV_ABS,
			Code: evdev.ABS_MT_PRESSURE,
			Value: 75,
		},
		evdev.InputEvent{
			Time: CurrentTimeval(),
			Type: evdev.EV_ABS,
			Code: evdev.ABS_MT_TOUCH_MAJOR,
			Value: 126,
		},
		evdev.InputEvent{
			Time: CurrentTimeval(),
			Type: evdev.EV_SYN,
			Code: evdev.SYN_REPORT,
			Value: 0,
		},
		evdev.InputEvent{
			Time: CurrentTimeval(),
			Type: evdev.EV_ABS,
			Code: evdev.ABS_MT_TRACKING_ID,
			Value: -1,
		},
		evdev.InputEvent{
			Time: CurrentTimeval(),
			Type: evdev.EV_SYN,
			Code: evdev.SYN_REPORT,
			Value: 0,
		},
	}

	for _, event := range events {
		err := d.WriteOne(&event)
		if err != nil {
			fmt.Println("Cannot send event")
			fmt.Println(event)
			return
		}
	}
}
