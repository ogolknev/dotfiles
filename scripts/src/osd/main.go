package main

import (
	"fmt"
	"os"
)

const (
	Black   = "\033[30m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
	Reset   = "\033[0m"
)

func main() {
	mode := os.Args[1]
	option := os.Args[2]

	fmt.Println(Blue+"OSD:"+Reset, Green+mode, option+Reset)

	var err error
	switch mode {
	case "fn":
		if option == "lock" {
			err = notify("Fn ", NotifyOptions{App: "osd-fn-lock", Category: "osd", Expire: "1000"})
		} else if option == "unlock" {
			err = notify("Fn ", NotifyOptions{App: "osd-fn-lock", Category: "osd", Expire: "1000"})
		}
	case "volume":
		value := option
		err = notify(value+"%", NotifyOptions{App: "osd-volume", Category: "osd", Expire: "1000", Value: "int:value:" + value})
	}

	if err != nil {

		fmt.Println(Red+"Error:"+Reset, err)
	}
}
