package main

import (
	//"./NetServer"

	"Common"
	"Game"
	"NetServer"
	"fmt"
)

func main() {
	fmt.Println("Game Server Start......")
	NetServer.InitMsg()
	Common.InitUserIdMgr()
	go NetServer.NetServerMain()
	Game.InitGame()

}
