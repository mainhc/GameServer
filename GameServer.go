package main

import (
	//"./NetServer"

	"fmt"
	"./NetServer"
	"./Game"
	
)


func main() {
	fmt.Println("Game Server Start......")
	NetServer.InitMsg()
	go NetServer.NetServerMain()
	Game.InitGame()
	
}
