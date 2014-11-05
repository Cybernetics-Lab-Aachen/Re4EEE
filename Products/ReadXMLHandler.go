package Products

import "fmt"
import "net/http"
import "github.com/SommerEngineering/Ocean/MimeTypes"
import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

func ReadXMLHandler(response http.ResponseWriter, request *http.Request) {
	Log.LogShort(senderName, LM.CategoryAPP, LM.LevelINFO, LM.MessageNameOUTPUT, `The current products XML file was requested.`)
	MimeTypes.Write2HTTP(response, MimeTypes.TypeXML)
	fmt.Fprintf(response, `%s`, readXML())
}
