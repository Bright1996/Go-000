package main

import (
	"fmt"
	"geek/Week02/controller"
)

func main() {
	fmt.Println(controller.Home{}.GetHomeInfo(1))
	fmt.Println(controller.Home{}.GetHomeInfo(2))
}
