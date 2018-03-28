package Game

import (
	// "log"
	"time"
)

const iMapWidth = 16
const iMapHeight = 16

var m_akPlayer map[int]int
var m_iPlayerNum int
var m_iGameState GameState
var m_gamebegintime time.Time

//记录当前的游戏状态

func InitGame() {
	InitTableMgr()
	InitMap(iMapWidth, iMapHeight)
	m_akPlayer = make(map[int]int)
	m_iPlayerNum = 0
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

func AddPlayer(clientid int) {
	m_iPlayerNum++
	m_akPlayer[clientid] = m_iPlayerNum
}

func GetGameCanStart() bool {
	if m_iGameState == KaiShi {
		return true
	} else {
		return false
	}
}

func updateGame(dt int) {
	//log.Print(dt)
	if m_iGameState == ZhunBie {
		newtime := time.Now()
		timego := int(newtime.Sub(m_gamebegintime).Nanoseconds() / 1e9)
		if timego >= 10 {
			m_iGameState = KaiShi
		}
	} else {
		if m_iGameState == KaiShi {

		}
	}
}
