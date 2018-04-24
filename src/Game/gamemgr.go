package Game

import (
	"bytes"
	//"container/list"
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

var m_iTimeGo int

//记录当前的游戏状态

func InitGame() {
	InitTableMgr()
	InitMap(iMapWidth, iMapHeight)

	m_iGameState = ZhunBie
	var gametime time.Time
	gametime = time.Now()
	m_gamebegintime = gametime
	m_iTimeGo = 10
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

func s2cZhunbei() {
	m_iTimeGo--
	sendMsg := &Player.CUiMessage{}
	tempstr := "updataZhunBei"
	sendMsg.UiMsgName = &tempstr
	iPlayerNum := NetServer.GetPlayerNum()
	paramint := []uint32{uint32(iPlayerNum), uint32(m_iTimeGo)}
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

	sendMsg := &Player.CGameRaceStart{}
	tempint := uint32(0)
	sendMsg.ClientId = &tempint
	senddata, _ := proto.Marshal(sendMsg)
	var sendbuf bytes.Buffer
	xNum := uint32(100004)
	tempsendbuf := bytes.NewBuffer([]byte{})
	binary.Write(tempsendbuf, binary.LittleEndian, xNum)
	sendbuf.Write(tempsendbuf.Bytes())
	sendbuf.Write(senddata)
	NetServer.BroadCastMsgToClient(sendbuf.Bytes())
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

			if NetServer.IsPlayerAllOk() {
				m_iGameState = KaiShiOK
				//通知玩家创建
				NetServer.CreatePlayer(0)
			}
		} else if m_iGameState == KaiShiOK {
			pTempList := NetServer.GetNextClientCmd()
			iLen := pTempList.Len()
			if iLen > 0 {
				var sendbuf bytes.Buffer
				xNum := uint32(100000)
				tempsendbuf := bytes.NewBuffer([]byte{})
				binary.Write(tempsendbuf, binary.LittleEndian, xNum)
				sendbuf.Write(tempsendbuf.Bytes())
				for pTemp := pTempList.Front(); pTemp != nil; pTemp = pTemp.Next() {
					pData := pTemp.Value.([]byte)
					sendbuf.Write(pData)
				}
				pTempList.Init()
				log.Print(sendbuf.Bytes())
				NetServer.BroadCastMsgToClient(sendbuf.Bytes())
			}

		}
	}
}
