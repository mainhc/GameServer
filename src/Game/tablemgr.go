package Game

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var akTableName [1]string
var akTabelData map[string]interface{}

func InitTableMgr() {
	akTabelData = make(map[string]interface{})

	akTableName[0] = "ObjView"

	akLength := len(akTableName)
	for iloop := 0; iloop < akLength; iloop++ {
		strname := akTableName[iloop]
		allstrname := "table/" + strname + ".json"
		log.Print(allstrname)
		bytes, err := ioutil.ReadFile(allstrname)
		if err != nil {
			continue
		}
		var pValue interface{}
		err2 := json.Unmarshal(bytes, &pValue)
		if err2 != nil {
			continue
		}
		akTabelData[strname] = pValue
	}
}

func GetTableData(strname string) interface{} {
	log.Print("+++++++++++++GetTableData")
	temp, ok := akTabelData[strname]
	if !ok {

		return nil
	}
	return temp
}

func GetTabelDataById(tabelname string, indexID string) interface{} {

	tabledata := GetTableData(tabelname)
	if tabledata == nil {
		return nil
	}
	datamap := tabledata.(map[string]interface{})
	tabelvalue, ok := datamap[indexID]
	if !ok {
		return nil
	}
	testtemp := tabelvalue.(map[string]interface{})
	log.Print(testtemp["ID"])
	return tabelvalue
}
