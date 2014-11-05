package VUFPL

import "io/ioutil"
import "github.com/SommerEngineering/Ocean/StaticFiles"
import "github.com/SommerEngineering/Ocean/CustomerDB"
import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

func readXML() (result string) {
	dbSession, gridFS := CustomerDB.GridFS()
	defer dbSession.Close()

	if gridFile, errFile := gridFS.Open(`VUFPL.xml`); errFile != nil {
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelWARN, LM.SeverityCritical, LM.ImpactUnknown, LM.MessageNameDATABASE, `Was not able to open the VUFPL program file.`, `This is okay for the first run but should not occur later on!`, errFile.Error())
		result = readEmptyXMLTemplate()
		return
	} else {
		defer gridFile.Close()
		if bytes, errRead := ioutil.ReadAll(gridFile); errRead != nil {
			Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `Was not able to read the VUFPL program file.`, errRead.Error())
			result = readEmptyXMLTemplate()
		} else {
			result = string(bytes)
		}
	}

	return
}

func readEmptyXMLTemplate() (result string) {
	bytes := StaticFiles.FindAndReadFile(`xml/VUFPL template.xml`)
	if bytes == nil {
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameFILESYSTEM, `Was not able to read the empty template for the VUFPL.`, `There is no fallback for this issue and the app is not usable.`)
		return
	} else {
		result = string(bytes)
	}
	return
}
