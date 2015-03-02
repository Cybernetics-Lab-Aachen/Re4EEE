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

func HandlerQuestion7(response http.ResponseWriter, request *http.Request) {
	noQuestion := 7
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

	data.Button1Data = `small`
	data.Button2Data = `mid`
	data.Button3Data = `huge`
	data.Button4Data = `masses`
	data.Button5Data = ``

	data.NoQuestion = fmt.Sprintf(`%d`, noQuestion)
	data.NoQuestions = totalQuestions
	data.Progress = fmt.Sprintf("%d", (int((float32(noQuestion) / float32(TOTAL_QUESTIONS)) * 100.0)))

	if strings.Contains(lang.Language, `de`) {
		data.TextButton1 = `Kleingruppen`
		data.TextButton2 = `Mittel (Dutzend(e))`
		data.TextButton3 = `Groß (Hundert(e))`
		data.TextButton4 = `Massen (Tausend(e)+)`
		data.TextButton5 = ``
		data.TextQuestion = `Frage`
		data.TextQuestionTopic = `Anzahl der gleichzeitgen Zugriffe`
		data.TextQuestionBody = `Wie viele Studierende greifen gleichzeitig auf das E-Learning-System zu?
		Bedenken Sie dabei, dass z.B. bei einer Gruppenarbeit nur die Anzahl der Gruppen zählt. Beispiel:
		In Ihrer Lehrveranstaltung befinden sich 1000 Studierende. Variante 1: Alle sollen zugleich das
		E-Learning nutzen können. Sie wählen "Massen". Variante 2: Sie teilen die Studierende in 100 Gruppen
		mit jeweils 10 Studierende. Pro Tag kommen 25 Gruppen in Ihren Rechnerpool, pro Stunde etwa 6 Gruppen.
		Es greifen somit ca. 60 Studierende auf das E-Learning-System zu, Sie wählen somit "Groß".`
	} else {
		data.TextButton1 = `Small`
		data.TextButton2 = `Mid (dozen(s))`
		data.TextButton3 = `Huge (hundred(s))`
		data.TextButton4 = `Masses (thousand(s)+)`
		data.TextButton5 = ``
		data.TextQuestion = `Question`
		data.TextQuestionTopic = `Count of Concurrent Accesses`
		data.TextQuestionBody = `How many students using the e-learning system concurrently? Please also
		consider in case of e.g. group work, that instead the amount of groups is necessary. Example: At your
		lecture you have 1,000 students. Variation 1: All students can access the e-learning system concurrently.
		In this case, you choose "Masses". Variation 2: You divide the students in 100 groups with 10 students each.
		Each day, 25 groups are going to the lab, around six groups per hour. In this case, around 60 students
		accesses concurrently the e-learning system. You choose "Huge".`
	}

	Tools.SendChosenLanguage(response, lang)
	Templates.ProcessHTML(`question`, response, data)
}
