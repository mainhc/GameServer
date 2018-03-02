package NetServer

import (	
    "container/list"  
)

var nouseid *list.List
var lastuseid int

func initUserIdMgr(){
    nouseid = list.New()
    lastuseid = 100001
}

func getCanUserID() int {
    if nouseid.Len() > 0 {
        res := nouseid.Front().Value.(int)
        nouseid.Remove(nouseid.Front())
        return res;
    } else {
        lastuseid +=1
        return lastuseid
    }
}












