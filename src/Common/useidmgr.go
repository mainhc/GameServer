package Common

import (
	"container/list"
	"log"
)

//玩家连接的ID
var nouseid *list.List
var lastuseid int

//局内单位的ID
var noObjuseID *list.List
var lastobjuseid int

func InitUserIdMgr() {
	nouseid = list.New()
	lastuseid = 100001
	noObjuseID = list.New()
	lastobjuseid = 0
}

func OnDeleClientID(clientId int) {
	log.Printf("onDeleClientID+++++++++")
	log.Print(clientId)
	nouseid.PushFront(clientId)
}

func GetCanUserID() int {
	if nouseid.Len() > 0 {
		res := nouseid.Front().Value.(int)
		nouseid.Remove(nouseid.Front())
		return res
	} else {
		lastuseid += 1
		return lastuseid
	}
}

func OnDeleObjUseClientID(objID int) {
	noObjuseID.PushFront(objID)
}

func GetCanObjuseID() int {
	if noObjuseID.Len() > 0 {
		res := noObjuseID.Front().Value.(int)
		noObjuseID.Remove(noObjuseID.Front())
		return res

	} else {
		lastobjuseid += 1
		return lastobjuseid
	}
}
