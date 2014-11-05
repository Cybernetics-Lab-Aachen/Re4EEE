package DB

import "github.com/SommerEngineering/Ocean/Log"
import LM "github.com/SommerEngineering/Ocean/Log/Meta"

func init() {

	Log.LogShort(senderName, LM.CategoryAPP, LM.LevelINFO, LM.MessageNameINIT, `Start init of customer database.`)
	defer Log.LogShort(senderName, LM.CategoryAPP, LM.LevelINFO, LM.MessageNameINIT, `Done init of customer database.`)

	// Get the database:
	//dbSession, db := CustomerDatabase.DB()
}
