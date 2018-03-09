package NetServer

import (	
    "encoding/json"
    "io/ioutil"
    "bytes"
    "strconv"
    "log"  
    "../msgconfig"
    "encoding/binary"
    "github.com/golang/protobuf/proto"
    "../Game"     
)

var m_msgstr map[string]int
var m_msgid map[int]string

var akToClientSock map[int]*WsSocket


func addClientSock(clientid int, sock *WsSocket){
    log.Print(clientid)
    _,ok := akToClientSock[clientid]
    if ok{
        return
    }
    akToClientSock[clientid] = sock
}

func deleClientSock(clientid int){
    _,ok := akToClientSock[clientid]
    if ok{
       delete(akToClientSock,clientid)
       onDeleClientID(clientid) 
    }  
}

func getClientSock(clientid int) *WsSocket {
    tempsock,ok := akToClientSock[clientid]
    if ok{
       return tempsock
    }  
    return nil
}

func registerMsg( msgstr string,msgid int){
    m_msgstr[msgstr] = msgid
    m_msgid[msgid] = msgstr

}

func getMsgProById(id int){

}

func InitMsg(){
    m_msgstr = make(map[string]int)
    m_msgid  = make(map[int]string)
    akToClientSock = make(map[int]*WsSocket)

    registerMsgByConfig();
    initUserIdMgr()
    //registerMsg("Player.cPlayerInfo",10001);
}

func registerMsgByConfig(){
    bytes, err := ioutil.ReadFile("msgconfig/MsgRegisterMap.json")
    if err != nil{
        return;
    }
    var msgmap map[string]string
    err2 := json.Unmarshal(bytes,&msgmap)
    if err2 != nil {       
        return;
    }
    for k, v := range msgmap {
        ivalue,err3 := strconv.Atoi(k)
        if err3 != nil{
            continue;
        }
        registerMsg(v,ivalue)
    }
}

func broadCastMsgToClient(sendMsg []byte){
    for _,v := range akToClientSock{
        v.SendIframe(sendMsg)
    }
}

func createPlayer(tempid int){
    for k,_ := range akToClientSock{
        clienttempid := k
        sendMsg := &Player.CPlayerCreator{
            ClientId:proto.Uint32(uint32(clienttempid)),
        }
        //sendMsg.Clinetid = clientId;
        senddata,_ := proto.Marshal(sendMsg)
        var sendbuf bytes.Buffer
        xNum :=uint32(100002)
        tempsendbuf := bytes.NewBuffer([]byte{})
            //客户端请求连上服务器
        binary.Write(tempsendbuf,binary.LittleEndian,xNum)
        sendbuf.Write(tempsendbuf.Bytes())
        //log.Print(sendbuf.Bytes())
        sendbuf.Write(senddata)
        broadCastMsgToClient(sendbuf.Bytes()) 
    }
     
}

func talkClientConnectSuc(clientSocket *WsSocket, clientId int){

    log.Printf("talkClientConnectSuc +++ %d",clientId)
    addClientSock(clientId,clientSocket)
    Game.AddPlayer(clientId)   
    sendMsg := &Player.CPlayerConnect{
        Clinetid:proto.Uint32(uint32(clientId)),
    }
    //sendMsg.Clinetid = clientId;
    senddata,_ := proto.Marshal(sendMsg)
    var sendbuf bytes.Buffer
    xNum :=uint32(100001)
    tempsendbuf := bytes.NewBuffer([]byte{})
        //客户端请求连上服务器
    binary.Write(tempsendbuf,binary.LittleEndian,xNum)
    sendbuf.Write(tempsendbuf.Bytes())
    //log.Print(sendbuf.Bytes())
    //binary.Write(sendbuf,binary.LittleEndian,senddata)
    sendbuf.Write(senddata)
    log.Print(sendbuf.Bytes())
    //bytes.Join(sendbuf,senddata);
    //broadCastMsgToClient(sendbuf.Bytes())
    clientSocket.SendIframe(sendbuf.Bytes())
    createPlayer(clientId)

}

func ClientSockDis(clientId int){
    deleClientSock(clientId)
}

func isPassMsg(msgID int) bool{
    if msgID <100000{
        return true
    }
    return false
}



func recvMsgFromClient(clinetid int,recvData []byte){
    //iLength := len(recvData)   
    msgId := []byte("0000")
    var tempbuf bytes.Buffer
    tempbuf.Write(recvData)    
    tempbuf.Read(msgId);
    //msgdata :=tempbuf.Bytes();
    //log.Print(msgId);
    mapintID := uint32(msgId[0]) | uint32(msgId[1])<<8 | uint32(msgId[2])<<16 | uint32(msgId[3])<<24
    log.Print(mapintID)
   // SendIframe(recvData);  
   if isPassMsg(int(mapintID)){
       log.Printf("+++++++++++Pass")
       log.Print(mapintID)
       if Game.GetGameCanStart() {
            broadCastMsgToClient(recvData)
       }      
        //clientsock :=getClientSock(clinetid)
        //clientsock.SendIframe(recvData)
   } 
}






