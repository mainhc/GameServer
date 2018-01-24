package NetServer

import (	
    "encoding/json"
    "io/ioutil"
    "bytes"
    "strconv"
	"log"	
    "../msgconfig"
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

func recvMsgFromClient(recvData []byte){
    //iLength := len(recvData)   
    msgId := []byte("0000")
    var tempbuf bytes.Buffer
    tempbuf.Write(recvData)    
    tempbuf.Read(msgId);
    msgdata :=tempbuf.Bytes();


    //iNum := recvData[0]
    log.Print(msgId);
    mapintID := uint32(msgId[0]) | uint32(msgId[1])<<8 | uint32(msgId[2])<<16 | uint32(msgId[3])<<24

    log.Print(mapintID);
    log.Print(msgdata);
    newTest := &Player.CPlayerInfo{}
    err01 := proto.Unmarshal(msgdata,newTest)
    if err01 != nil {
       // log.Printf("not  %s  ",recvdata);
    } else {
        log.Printf("palyerdata  %d   %s   %d",newTest.GetId(),newTest.GetName(),newTest.GetEnterTime());
    }

}






