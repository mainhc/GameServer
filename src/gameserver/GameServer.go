package main

import (
	//"./NetServer"

	"Game"
	"NetServer"
	"fmt"
)

func main() {
	fmt.Println("Game Server Start......")
	NetServer.InitMsg()
	go NetServer.NetServerMain()
	Game.InitGame()

}
