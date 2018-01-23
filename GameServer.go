package main

import (
	//"./NetServer"

	"fmt"
	"./NetServer"
	
)


func main() {
	fmt.Println("Game Server Start......")
	NetServer.InitMsg()
	NetServer.NetServerMain()
	
}
