package main

import "github.com/SommerEngineering/Ocean/System"
import "github.com/SommerEngineering/Ocean/Shutdown"
import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

func main() {

	Log.LogShort(senderName, LM.CategoryAPP, LM.LevelINFO, LM.MessageNameSTARTUP, `The application is starting.`)
	System.InitHandlers()
	registerAllAppHandler()

	Shutdown.AddShutdownHandler(shutdownFunction{})
	System.StartAndBlockForever()
}

type shutdownFunction struct {
}

func (s shutdownFunction) Shutdown() {
	Log.LogShort(senderName, LM.CategoryAPP, LM.LevelWARN, LM.MessageNameSHUTDOWN, `The application is shutting down.`)
}
