package VUFPL

import "net/http"
import "io/ioutil"
import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

func WriteXMLHandler(response http.ResponseWriter, request *http.Request) {
	Log.LogShort(senderName, LM.CategoryAPP, LM.LevelINFO, LM.MessageNameINPUT, `A new VUFPL XML file was received.`)
	if bytes, err := ioutil.ReadAll(request.Body); err != nil {
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityHigh, LM.ImpactHigh, LM.MessageNameINPUT, `Was not able to read the XML data from the stream.`)
		http.NotFound(response, request)
	} else {
		writeXML(bytes)
	}
}
