package Website

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Templates"
	"github.com/SommerEngineering/Ocean/Tools"
	"net/http"
	"strings"
)

func HandlerQuestion20(response http.ResponseWriter, request *http.Request) {
	noQuestion := 20
	readSession := request.FormValue(`session`)
	if readSession == `` {
		defer http.Redirect(response, request, "/", 307)
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelSECURITY, LM.SeverityMiddle, LM.ImpactNone, LM.MessageNameUSER, `A request without session.`)
		return
	}

	lang := Tools.GetRequestLanguage(request)[0]
	data := PageQuestion{}
	data.Basis.Name = NAME
	data.Basis.Version = VERSION
	data.Basis.Lang = lang.Language
	data.Basis.Session = readSession

	data.Button1Status = BUTTON_SHOW
	data.Button2Status = BUTTON_SHOW
	data.Button3Status = BUTTON_SHOW
	data.Button4Status = BUTTON_SHOW
	data.Button5Status = BUTTON_HIDDEN

	data.Button1Data = `static`
	data.Button2Data = `change`
	data.Button3Data = `interact`
	data.Button4Data = `*`
	data.Button5Data = ``

	data.NoQuestion = fmt.Sprintf(`%d`, noQuestion)
	data.NoQuestions = totalQuestions
	data.Progress = fmt.Sprintf("%d", (int((float32(noQuestion) / float32(TOTAL_QUESTIONS)) * 100.0)))

	if strings.Contains(lang.Language, `de`) {
		data.TextButton1 = `Statische Inhalte`
		data.TextButton2 = `Veränderungen möglich`
		data.TextButton3 = `Interaktion möglich`
		data.TextButton4 = `Unbekannt`
		data.TextButton5 = ``
		data.TextQuestion = `Frage`
		data.TextQuestionTopic = `Art der Inhalte`
		data.TextQuestionBody = `Welcher Art sind Ihre Inhalte für das E-Learning-Tool? Statische Inhalte lassen
		sich nur konsumieren (z.B. Präsentationen, Videos, Bücher), während veränderbare Inhalte bearbeitet
		werden können (z.B. Wikis). Inhalte mit möglichen Interaktionen können z.B. eine Simulation oder
		eine virtuelle Umgebung sein.`
	} else {
		data.TextButton1 = `Static`
		data.TextButton2 = `Changes possible`
		data.TextButton3 = `Interaction possible`
		data.TextButton4 = `Unknown`
		data.TextButton5 = ``
		data.TextQuestion = `Question`
		data.TextQuestionTopic = `Kind of Content`
		data.TextQuestionBody = `Which kind of content do you want to provide for the e-learning solution? Static
		content is just consumable (e.g. presentations, videos, books), wherever changeable content can be edit
		by the students (e.g. Wikis). Examples for content with interactions are e.g. a simulation or a
		virtual environment.`
	}

	Tools.SendChosenLanguage(response, lang)
	Templates.ProcessHTML(`question`, response, data)
}
