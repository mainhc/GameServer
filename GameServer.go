package main

import (
	//"./NetServer"

	"fmt"
	"./NetServer"
)
//import("./NetServer")

func main() {
	fmt.Println("Game Server Start......")
	NetServer.NetServerMain()
}
