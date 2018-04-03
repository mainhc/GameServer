package Game

import (
	"bytes"
	"encoding/binary"
	"log"
	"msgconfig"
	"sync"
	"time"

	"NetServer"

	"github.com/golang/protobuf/proto"
)

const iMapWidth = 16
const iMapHeight = 16

var m_iGameState GameState
var m_gamebegintime time.Time

//记录当前的游戏状态

func InitGame() {
	InitTableMgr()
	InitMap(iMapWidth, iMapHeight)

	m_iGameState = ZhunBie
	var gametime time.Time
	gametime = time.Now()
	m_gamebegintime = gametime
	//GetTabelDataById("ObjView","10002")
	for {
		newtime := time.Now()
		timego := newtime.Sub(gametime).Nanoseconds()
		misgo := int(timego / 1e6)
		if misgo >= 50 {
			gametime = newtime
			updateGame(misgo)
		}

	}

}

func GetGameCanStart() bool {
	if m_iGameState == KaiShi {
		return true
	} else {
		return false
	}
}

func s2cZhunbei() {

	sendMsg := &Player.CUiMessage{}
	tempstr := "updataZhunBei"
	sendMsg.UiMsgName = &tempstr
	paramint := []uint32{uint32(3), 12}
	sendMsg.AkMsgParame = paramint
	log.Print("updataZhunBei")

	senddata, _ := proto.Marshal(sendMsg)
	var sendbuf bytes.Buffer
	xNum := uint32(100003)
	tempsendbuf := bytes.NewBuffer([]byte{})
	binary.Write(tempsendbuf, binary.LittleEndian, xNum)
	sendbuf.Write(tempsendbuf.Bytes())
	sendbuf.Write(senddata)
	NetServer.BroadCastMsgToClient(sendbuf.Bytes())
}

func doZhunBei() {
	ticker := time.NewTicker(time.Second * 1)
	wg := sync.WaitGroup{}
	wg.Add(10)
	go func() {
		for _ = range ticker.C {

			s2cZhunbei()
			wg.Done()
		}
	}()
	wg.Wait()
	NetServer.SetGameCanStart()
	ticker.Stop()
}

func updateGame(dt int) {
	//log.Print(dt)

	if m_iGameState == ZhunBie {
		newtime := time.Now()
		timego := int(newtime.Sub(m_gamebegintime).Nanoseconds() / 1e9)
		if timego >= 10 {
			m_iGameState = KaiShi
			doZhunBei()
		}
	} else {
		if m_iGameState == KaiShi {

		}
	}
}
