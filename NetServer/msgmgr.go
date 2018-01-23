package NetServer

import (	
    "encoding/json"  
    "io/ioutil"
    "strconv"
	//"log"	
    //"../msgconfig"
    //"github.com/golang/protobuf/proto"    
)

var m_msgstr map[string]int
var m_msgid map[int]string

func registerMsg( msgstr string,msgid int){
    m_msgstr[msgstr] = msgid
    m_msgid[msgid] = msgstr
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






