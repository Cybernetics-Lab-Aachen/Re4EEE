package main

import "net/http"
import "github.com/SommerEngineering/Re4EEE/Products"
import "github.com/SommerEngineering/Re4EEE/VUFPL"
import "github.com/SommerEngineering/Re4EEE/UI"
import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

func registerAllAppHandler() {

	Log.LogShort(senderName, LM.CategoryAPP, LM.LevelINFO, LM.MessageNameSTARTUP, `Register now all app handlers.`)
	defer Log.LogShort(senderName, LM.CategoryAPP, LM.LevelINFO, LM.MessageNameSTARTUP, `Register now all app handlers done.`)

	http.HandleFunc(`/admin`, UI.AdminViewVUFPLHandler)
	http.HandleFunc(`/admin/VUFPL/show`, UI.AdminViewVUFPLHandler)
	http.HandleFunc(`/data/VUFPL/read`, VUFPL.ReadXMLHandler)
	http.HandleFunc(`/data/VUFPL/write`, VUFPL.WriteXMLHandler)
	http.HandleFunc(`/data/products/read`, Products.ReadXMLHandler)
	http.HandleFunc(`/data/products/write`, Products.WriteXMLHandler)
}
