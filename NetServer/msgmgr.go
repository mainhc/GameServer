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
)

var m_msgstr map[string]int
var m_msgid map[int]string

func registerMsg( msgstr string,msgid int){
    m_msgstr[msgstr] = msgid
    m_msgid[msgid] = msgstr

}

func getMsgProById(id int){

}

func InitMsg(){
    m_msgstr = make(map[string]int)
    m_msgid  = make(map[int]string)

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

func talkClientConnectSuc(clientSocket *WsSocket, clientId int){

    log.Printf("talkClientConnectSuc +++ %d",clientId)
    sendMsg := &Player.CPlayerConnect{
        Clinetid:proto.Uint32(uint32(clientId)),
    }
    //sendMsg.Clinetid = clientId;
    senddata,_ := proto.Marshal(sendMsg)
    var sendbuf bytes.Buffer
    xNum :=uint32(10004)
    tempsendbuf := bytes.NewBuffer([]byte{})
        //客户端请求连上服务器
    binary.Write(tempsendbuf,binary.LittleEndian,xNum)
    sendbuf.Write(tempsendbuf.Bytes())
    log.Print(sendbuf.Bytes())
    //binary.Write(sendbuf,binary.LittleEndian,senddata)
    sendbuf.Write(senddata)
    log.Print(sendbuf.Bytes())
    //bytes.Join(sendbuf,senddata);
    clientSocket.SendIframe(sendbuf.Bytes())

}



func recvMsgFromClient(recvData []byte){
    //iLength := len(recvData)   
    msgId := []byte("0000")
    var tempbuf bytes.Buffer
    tempbuf.Write(recvData)    
    tempbuf.Read(msgId);
    msgdata :=tempbuf.Bytes();
    log.Print(msgId);
    mapintID := uint32(msgId[0]) | uint32(msgId[1])<<8 | uint32(msgId[2])<<16 | uint32(msgId[3])<<24
    
   // SendIframe(recvData);        
   
    log.Print(mapintID)
    log.Print(msgdata)
   
    newTest := &Player.CPlayerInfo{}
    err01 := proto.Unmarshal(msgdata,newTest)
    if err01 != nil {
       // log.Printf("not  %s  ",recvdata);
    } else {
        log.Print(newTest);
        log.Printf("palyerdata  %d   %s   %d",newTest.GetId(),newTest.GetName(),newTest.GetEnterTime());
    }

}






