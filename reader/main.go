package main

import (
	"fmt"

	"github.com/holoplot/go-evdev"
)

func main() {
	d, err := evdev.Open("/dev/input/event1")
	if err != nil {
		fmt.Printf("cannot read input device")
		fmt.Println(err)
		return
	}

	for {
		event, err := d.ReadOne()
		if err != nil {
			fmt.Printf("error reading event")
			fmt.Println(err)
			return
		}

		fmt.Println("event")
		fmt.Println(event)
		fmt.Printf("%+v\n", event)
	}
}
