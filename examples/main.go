package main

import (
	"fmt"
	"github.com/RBG-TUM/go-anel-pwrctrl"
	"os"
	"time"
)

//usage: ./main http://url-to-device user:password
func main() {
	client := go_anel_pwrctrl.New(os.Args[1], os.Args[2])
	fmt.Println("Turning on")
	err := client.TurnOn(0)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Second * 5)
	fmt.Println("Turning off")
	err = client.TurnOff(0)
	if err != nil {
		fmt.Println(err)
	}
}
