package NetServer

import (
	"Common"
	"bytes"
	"container/list"
	"encoding/binary"
	"encoding/json"
	"io/ioutil"
	"log"
	"msgconfig"
	"strconv"

	"github.com/golang/protobuf/proto"
)

var m_msgstr map[string]int
var m_msgid map[int]string

var m_bIsGameIsStart bool

var akToClientSock map[int]*WsSocket

//缓存帧中的操作指令
var nextClientCmd *list.List

func addClientSock(clientid int, sock *WsSocket) {
	log.Print(clientid)
	_, ok := akToClientSock[clientid]
	if ok {
		return
	}
	akToClientSock[clientid] = sock
}

func deleClientSock(clientid int) {
	_, ok := akToClientSock[clientid]
	if ok {
		delete(akToClientSock, clientid)
		Common.OnDeleClientID(clientid)
	}
}

func getClientSock(clientid int) *WsSocket {
	tempsock, ok := akToClientSock[clientid]
	if ok {
		return tempsock
	}
	return nil
}

func GetNextClientCmd() *list.List {

	return nextClientCmd
}

func registerMsg(msgstr string, msgid int) {

	log.Print(msgstr)
	m_msgstr[msgstr] = msgid
	m_msgid[msgid] = msgstr

}

func getMsgProById(id int) {

}
func isGameCanStart() bool {
	return m_bIsGameIsStart
}

func InitMsg() {
	m_msgstr = make(map[string]int)
	m_msgid = make(map[int]string)
	nextClientCmd = list.New()
	akToClientSock = make(map[int]*WsSocket)
	m_bIsGameIsStart = false

	registerMsgByConfig()

	InitClinetCmd()

}

func SetGameCanStart() {
	m_bIsGameIsStart = true
}

func registerMsgByConfig() {

	bytes, err := ioutil.ReadFile("../msgconfig/MsgRegisterMap.json")
	if err != nil {
		return
	}
	var msgmap map[string]string
	err2 := json.Unmarshal(bytes, &msgmap)
	if err2 != nil {
		return
	}
	for k, v := range msgmap {
		ivalue, err3 := strconv.Atoi(k)
		if err3 != nil {
			continue
		}
		registerMsg(v, ivalue)
	}
}

func BroadCastMsgToClient(sendMsg []byte) {
	for _, v := range akToClientSock {
		v.SendIframe(sendMsg)
	}
}

func CreatePlayer(tempid int) {
	log.Print("CreatePlayer")
	for k, _ := range akToClientSock {
		clienttempid := k
		sendMsg := &Player.CPlayerCreator{
			ClientId: proto.Uint32(uint32(clienttempid)),
		}

		senddata, _ := proto.Marshal(sendMsg)
		var sendbuf bytes.Buffer
		xNum := uint32(100002)
		tempsendbuf := bytes.NewBuffer([]byte{})
		//客户端请求连上服务器
		binary.Write(tempsendbuf, binary.LittleEndian, xNum)
		sendbuf.Write(tempsendbuf.Bytes())

		sendbuf.Write(senddata)

		log.Print("CreatePlayer++++++++")
		BroadCastMsgToClient(sendbuf.Bytes())
	}

}

func GetPlayerNum() int {
	return len(akToClientSock)
}

func talkClientConnectSuc(clientSocket *WsSocket, clientId int) {

	log.Printf("talkClientConnectSuc +++ %d", clientId)
	addClientSock(clientId, clientSocket)

	sendMsg := &Player.CPlayerConnect{
		Clinetid: proto.Uint32(uint32(clientId)),
	}

	senddata, _ := proto.Marshal(sendMsg)
	var sendbuf bytes.Buffer
	xNum := uint32(100001)
	tempsendbuf := bytes.NewBuffer([]byte{})
	//客户端请求连上服务器
	binary.Write(tempsendbuf, binary.LittleEndian, xNum)
	sendbuf.Write(tempsendbuf.Bytes())

	sendbuf.Write(senddata)
	log.Printf("testConnectSuc+++")

	clientSocket.SendIframe(sendbuf.Bytes())
	//createPlayer(clientId)

}

func ClientSockDis(clientId int) {
	deleClientSock(clientId)
}

func doClientCmd(msgId int, recvData []byte) {
	tempmsgId := []byte("0000")
	var tempbuf bytes.Buffer
	tempbuf.Write(recvData)
	tempbuf.Read(tempmsgId)
	log.Print(tempmsgId)

	switch msgId {
	case 200001:
		{
			ProcessC2sClientGameWorldOK(tempbuf.Bytes())
			break
		}
	default:
		{
			break
		}

	}
}

func isPassMsg(msgID int) bool {
	if msgID < 100000 {
		return true
	}
	return false
}

func recvMsgFromClient(clinetid int, recvData []byte) {
	msgId := []byte("0000")
	var tempbuf bytes.Buffer
	tempbuf.Write(recvData)
	tempbuf.Read(msgId)

	mapintID := uint32(msgId[0]) | uint32(msgId[1])<<8 | uint32(msgId[2])<<16 | uint32(msgId[3])<<24
	log.Print(mapintID)

	if isPassMsg(int(mapintID)) {
		log.Printf("+++++++++++Pass")
		log.Print(mapintID)
		if isGameCanStart() {
			//BroadCastMsgToClient(recvData)
			nextClientCmd.PushBack(recvData)
		}
	} else {
		log.Print("recvMsgFromClient")
		doClientCmd(int(mapintID), recvData)
	}
}
