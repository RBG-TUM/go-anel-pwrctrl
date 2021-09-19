package main

import (
	"github.com/RBG-TUM/go-anel-pwrctrl"
	"log"
)

func main()  {
	client := go_anel_pwrctrl.New("http://device-addr.de", "user:password")
	isOn, err := client.IsOn(0)
	log.Print(isOn, err)
}
