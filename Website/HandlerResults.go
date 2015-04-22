package Website

import (
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Templates"
	"github.com/SommerEngineering/Ocean/Tools"
	"github.com/SommerEngineering/Re4EEE/Algorithm"
	"github.com/SommerEngineering/Re4EEE/DB"
	"github.com/SommerEngineering/Re4EEE/DB/Scheme"
	"github.com/SommerEngineering/Re4EEE/XML"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func HandlerResults(response http.ResponseWriter, request *http.Request) {

	session := request.FormValue(`session`)
	amountText := request.FormValue(`amount`)
	lang := Tools.GetRequestLanguage(request)[0]
	answers := DB.LoadAnswers(session)
	groups := XML.GetData()
	amountValue := -1
	resultSet := Scheme.Recommendation{}

	if !DB.CheckRecommendation(session) {

		assessedGroups := Algorithm.ExecuteAnswers(answers)

		resultSet.ProductGroups = assessedGroups
		resultSet.CreateTimeUTC = time.Now().UTC()
		resultSet.Session = session
		DB.StoreRecommendation(resultSet)

	} else {
		resultSet = DB.LoadRecommendation(session)
	}

	if value, errConv := strconv.Atoi(amountText); errConv != nil {
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityMiddle, LM.ImpactNone, LM.MessageNameREQUEST, `Cannot read the amount value!`, amountText, errConv.Error())
	} else {
		amountValue = value
	}

	if amountValue >= 1 {
		resultSet.ProductGroups = resultSet.ProductGroups[0:amountValue]
	}

	data := PageResults{}
	data.Basis.Version = VERSION
	data.Basis.Lang = lang.Language
	data.Basis.Session = session
	data.Groups = groups.ProductsCollection.Products
	data.Recommendation = resultSet
	data.AmountCurrent = amountValue

	if strings.Contains(lang.Language, `de`) {

		if amountValue > 0 {
			data.TextAllGroups = `Alle Gruppen anzeigen`
			data.AmountToggle = -1
		} else {
			data.TextAllGroups = `Top 6 anzeigen`
			data.AmountToggle = 6
		}

		data.Basis.Name = NAME_DE
		data.LangPos = 0
		data.TextMatch = `Übereinstimmung mit Ihren Antworten`
		data.TextGroup = `Gruppe`
		data.TextExamples = `Beispiele für`
		data.TextOptionen = `Optionen`
		data.TextHeaderPrefix = `Ihre E-Learning-Empfehlung`
		data.TextRestart = `Den Fragebogen erneut starten`

		data.TextHeader1 = `Unten sehen Sie eine Empfehlung für verschiedene Gruppen von E-Learning-Systemen, basierend auf
		Ihren Antworten. Bitte wählen Sie eine Gruppe, um Beispiele sehen zu können. Zunächst sehen Sie nur die Top Sechs
		aller E-Learning-Gruppen: Mit der nachfolgenden Schaltfläche (siehe Optionen) können Sie auch alle
		Gruppenergebnisse einsehen.`

		data.TextHeader2 = `Bitte beachten Sie, dass nur Beispiele in den Gruppen dargestellt werden. Die Aufzählung hat keinen
		Anspruch auf Vollständigkeit. Möglicherweise existiert in Ihrer Hochschule eine E-Learning-Strategie: Bitte erkundigen
		Sie sich, inwieweit Sie die hier vorgeschlagenen Formate einsetzen können und dürfen. Einige der aufgeführten Formate
		benutzen "die Cloud". Das bedeutet, dass sie Dienste von Dritten einsetzen. Bitte halten Sie aus diesen Gründen
		Rücksprache mit Ihrem Datenschutzbeauftragten, der Sie zu diesen Themen beraten kann.`

		data.TextHeader3 = `Einige der hier vorgestellten Formate sind kostenpflichtig, andere Formate sind dagegen ohne Kosten
		frei erhältlich. Für viele kostenpflichtige Formate können kostenlose Alternativen gefunden werden. Bei dem Einsatz
		einiger Formate können Folgekosten entstehen. Beachten Sie außerdem, dass die Studierenden in einigen Formaten anonym
		auftreten können. Sollten Sie dies nicht wünschen, prüfen Sie bitte geeignete Gegenmaßnahmen.`

	} else {

		if amountValue > 0 {
			data.TextAllGroups = `Show all groups`
			data.AmountToggle = -1
		} else {
			data.TextAllGroups = `Show top 6`
			data.AmountToggle = 6
		}

		data.Basis.Name = NAME_EN
		data.LangPos = 1
		data.TextMatch = `match with your answers`
		data.TextGroup = `Group`
		data.TextOptionen = `Options`
		data.TextExamples = `Examples for`
		data.TextHeaderPrefix = `Your E-Learning Recommendation`
		data.TextRestart = `Restart the questionnaire`
		data.TextHeader1 = `Please find below your e-learning recommendation, based on your answers. Please choose a
		group to show the details with the examples. Initially, just the top 8 groups are visible. With the options
		below, you can access also the results of all groups.`

		data.TextHeader2 = `Please consider that the examples at the groups are only a small extract and not a
		complete listing. It is possible that your university provides an e-learning strategy: Please ensure,
		that you can use the formats within the examples. Some formats are using cloud services. Please consider
		the privacy policy of your university to ensure, that you can use these formats. Several formats are only
		available as paid version; others are free. Please enquire, because it is possible that both charges for
		the usage. Finally, it is possible that the students are anonymous within formats. If this is unwanted,
		please plan countermeasures.`

		data.TextHeader3 = ``
	}

	Tools.SendChosenLanguage(response, lang)
	Templates.ProcessHTML(`results`, response, data)
}
