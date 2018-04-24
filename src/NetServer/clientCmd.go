package NetServer

import (
	"log"
	"msgconfig"

	"github.com/golang/protobuf/proto"
)

var playerOk map[int]int

func InitClinetCmd() {
	playerOk = make(map[int]int)
}

func IsPlayerAllOk() bool {
	iPlayerNum := GetPlayerNum()
	if iPlayerNum == len(playerOk) {
		return true
	}
	return false
}

func ProcessC2sClientGameWorldOK(recvData []byte) {
	log.Print("+++++++++++++1")
	clientMsg := &Player.C2SClientGameWorldOK{}
	err := proto.Unmarshal(recvData, clientMsg)
	log.Print(recvData)
	if err == nil {
		log.Print(clientMsg.GetMyClientID())
		log.Print("+++++++++++++2")
		playerOk[int(clientMsg.GetMyClientID())] = 1
	}

}
