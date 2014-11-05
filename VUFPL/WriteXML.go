package VUFPL

import "github.com/SommerEngineering/Ocean/CustomerDB"
import "github.com/SommerEngineering/Ocean/MimeTypes"
import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

func writeXML(xmlData []byte) {
	dbSession, gridFS := CustomerDB.GridFS()
	defer dbSession.Close()

	if gridFile, errFile := gridFS.Create(`VUFPL.xml`); errFile != nil {
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameDATABASE, `Was not able to create a new GridFS file.`, errFile.Error())
	} else {
		defer gridFile.Close()
		gridFile.SetContentType(MimeTypes.TypeXML.MimeType)
		gridFile.Write(xmlData)
		Log.LogShort(senderName, LM.CategoryAPP, LM.LevelINFO, LM.MessageNameINPUT, `The new VUFPL program was saved.`)
	}
}
