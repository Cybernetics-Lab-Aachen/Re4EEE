package Website

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/ConfigurationDB"
	"github.com/SommerEngineering/Ocean/Templates"
	"github.com/SommerEngineering/Ocean/Tools"
	"github.com/SommerEngineering/Re4EEE/DB"
	"github.com/SommerEngineering/Re4EEE/DB/Scheme"
	"net/http"
	"strings"
	"time"
)

//HandlerStart displays the start page
func HandlerStart(response http.ResponseWriter, request *http.Request) {
	readSession := request.FormValue(`session`)
	lang := Tools.GetRequestLanguage(request)[0]
	data := PageStart{}
	data.Basis.Version = VERSION
	data.Basis.Lang = lang.Language
	data.Basis.SiteVerificationToken = ConfigurationDB.Read("SiteVerificationToken")

	// Check if sesion is valid
	if readSession != `` {
		answers, loadAnswersError := DB.LoadAnswers(readSession)
		if loadAnswersError || !answers.StartTimeQ1.IsZero() { // Session doesn't exist or questionnaire was already started
			readSession = `` // Forfeit session
		}
	}

	//Check if a session exists
	if readSession != `` {
		data.Basis.Session = readSession
	} else {
		//Create new session
		data.Basis.Session = Tools.RandomGUID()
		answers := Scheme.Answers{}
		answers.SchemeVersion = Scheme.CURRENT_VERSION
		answers.Session = data.Basis.Session
		answers.CreateTimeUTC = time.Now().UTC()
		DB.StoreNewAnswers(answers)

		//Redirect so created session shows up in address bar
		http.Redirect(response, request, fmt.Sprintf(`/start?lang=%s&session=%s`, data.Basis.Lang, data.Basis.Session), 302)
		return
	}

	if strings.Contains(lang.Language, `de`) {
		data.Basis.Name = NAME_DE
		data.Basis.Logo = LOGO_DE
		data.TextStartButton = `Fragebogen beginnen`
		data.TextProject = `Projekt`
		data.TextExecuted = `Durchgeführt von`
		data.TextPromoted = `Gefördert von`
		data.TextWelcome = `Herzlich willkommen beim ` + NAME_DE_PLAIN + `. Der ` + NAME_DE_PLAIN + ` ist ein E-Learning-Empfehlungssystem
		mit Fokus auf die MINT-Wissenschaften. Ziel ist es, dass Sie einen Überblick über geeignete Formate erhalten,
		die Sie in Ihrer Veranstaltung einsetzen können. Wenn Sie mit der Schaltfläche unten den Fragebogen starten,
		werden Ihnen einige Fragen gestellt. Basierend auf Ihren Antworten wird dann eine geeignete Empfehlung
		ausgegeben. Der ` + NAME_DE_PLAIN + ` wird realisiert im Projekt ELLI – Exzellentes Lehren und Lernen in den
		Ingenieurwissenschaften – welches im Rahmen des "Qualitätspakt Lehre" vom Bundesministerium für Bildung und Forschung gefördert wird.`
	} else {
		data.TextStartButton = `Start Questionnaire`
		data.Basis.Name = NAME_EN
		data.Basis.Logo = LOGO_UK
		data.TextProject = `Project`
		data.TextExecuted = `Executed by`
		data.TextPromoted = `Promoted by`
		data.TextWelcome = `Welcome to ` + NAME_EN + `. The ` + NAME_EN_PLAIN + ` is an e-learning recommendation system
		with focus on STEM sciences. Our aim is to provide you with an overview about suitable tools which you can use
		e.g. in your courses, lectures, etc. As soon as you start the questionnaire by pressing the button below, you
		will receive a few questions. Afterwards, the ` + NAME_EN_PLAIN + ` will be able to recommend e-learning formats for
		you, based on your answers. The ` + NAME_EN_PLAIN + ` is part of the project ELLI – Excellent Teaching and Learning in Engineering
		Sciences. This project is supported by the Federal Ministry of Education and Research, Germany.`
	}

	Tools.SendChosenLanguage(response, lang)
	Templates.ProcessHTML(`start`, response, data)
}
