package Game

import (	   
   // "log"
    "container/list" 
)

var m_iMapWidth int
var m_iMapHeight int

type ObjGrid struct {
    posX            int
    posY            int
    objIndex        int
}

var m_akMapGrid *list.List




func InitMap(mapWidth int , mapHeight int){
    m_akMapGrid = list.New()
    m_iMapWidth = mapWidth
    m_iMapHeight = mapHeight
    for iloop:=0;iloop<m_iMapHeight;iloop++ {
        for jloop:=0;jloop<m_iMapWidth;jloop++ {
            indexint := iloop * m_iMapWidth + jloop
            tempGrid :=ObjGrid{posX:jloop,posY:iloop,objIndex:indexint}
            m_akMapGrid.PushFront(tempGrid)
        }
    }
}




















