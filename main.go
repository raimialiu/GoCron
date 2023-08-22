package main

import (
	"GoSimpleCron/lib"
	"fmt"
	"time"
)

func main() {
	fmt.Println("GoSimpleCron.....func implementation....")

	goCron := lib.Instance(lib.GoSimpleCronConfig{
		RecoverOnError: false,
	})

	//goCron(2, PrintName)
	//goCron.DoEvery(PrintName).Second()
}

func PrintName() {
	name := " olatunde"
	fmt.Println("My name is " + name)
}

// https://www.golinuxcloud.com/golang-multiply-duration-by-integer/
func DoEvery(timeout int32, taskFunc func()) {
	for range time.Tick(time.Second * time.Duration(timeout)) {
		taskFunc()
	}
}
