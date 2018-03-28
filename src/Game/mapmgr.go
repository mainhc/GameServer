package Game

import (
	"Common"
)

var m_iMapWidth int
var m_iMapHeight int

type ObjGrid struct {
	posX     int
	posY     int
	objIndex int
}

var m_akMapGrid map[int]*ObjGrid

func InitMap(mapWidth int, mapHeight int) {
	m_akMapGrid = make(map[int]*ObjGrid)
	m_iMapWidth = mapWidth
	m_iMapHeight = mapHeight
	for iloop := 0; iloop < m_iMapHeight; iloop++ {
		for jloop := 0; jloop < m_iMapWidth; jloop++ {
			indexint := iloop*m_iMapWidth + jloop
			tempGrid := &ObjGrid{posX: jloop, posY: iloop, objIndex: indexint}
			m_akMapGrid[indexint] = tempGrid
		}
	}
}

func CreateObj(x int, y int, playerobjid int) int {
	indexint := y*m_iMapWidth + x
	pGrid, ok := m_akMapGrid[indexint]
	if !ok {
		objIndex := Common.GetCanObjuseID()
		pGrid.objIndex = objIndex
		pGrid.posX = x
		pGrid.posY = y
		return objIndex
	}
	return -1
}
