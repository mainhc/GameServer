package NetServer

import (	
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



func init(){
    m_msgstr = make(map[string]int)
    m_msgid  = make(map[int]string)
    registerMsg("Player.cPlayerInfo",10001);

}






