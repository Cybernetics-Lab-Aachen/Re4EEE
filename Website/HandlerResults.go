package Website

import (
	"github.com/SommerEngineering/Ocean/ConfigurationDB"
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

// HandlerResults displays the results of the questionnaire
func HandlerResults(response http.ResponseWriter, request *http.Request) {
	session := request.FormValue(`session`)
	amountText := request.FormValue(`amount`)
	lang := Tools.GetRequestLanguage(request)[0]
	answers, loadAnswersError := DB.LoadAnswers(session)
	groups := XML.GetData()
	amountValue := -1
	resultSet := Scheme.Recommendation{}

	// Check if session exists, otherwise redirect to start page
	if loadAnswersError {
		http.Redirect(response, request, `/start`, 302)
		return
	}

	// Validate input
	if amountText != `` && len(amountText) > 2 {
		response.WriteHeader(http.StatusNotFound)
		return
	}
	if value, errConv := strconv.Atoi(amountText); errConv != nil {
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityMiddle, LM.ImpactNone, LM.MessageNameREQUEST, `Cannot read the amount value!`, amountText, errConv.Error())
	} else {
		amountValue = value
	}

	// Load/calculate recommendations
	if !DB.CheckRecommendation(session) {
		resultSet.ProductGroups = Algorithm.ExecuteAnswers(answers)
		resultSet.CreateTimeUTC = time.Now().UTC()
		resultSet.Session = session
		resultSet.SchemeVersion = Scheme.CURRENT_VERSION
		DB.StoreRecommendation(resultSet)

	} else {
		resultSet = DB.LoadRecommendation(session)

		// Check for old session. Can't work with outdated data.
		if resultSet.SchemeVersion < Scheme.CURRENT_VERSION {
			http.Redirect(response, request, `/start`, 302)
			return
		}
	}

	// Reduce number of shown product groups, if requested
	if amountValue >= 1 {
		resultSet.ProductGroups = resultSet.ProductGroups[0:amountValue]
	}

	// Prepare localized strings
	data := PageResults{}
	data.Basis.Version = VERSION
	data.Basis.Lang = lang.Language
	data.Basis.Session = session
	data.Basis.SiteVerificationToken = ConfigurationDB.Read("SiteVerificationToken")
	data.Groups = groups.ProductsCollection.Products
	data.Questions = groups.QuestionsCollection.Questions
	data.Recommendation = resultSet
	data.AmountCurrent = amountValue
	data.Strings = groups.ResultStrings
	data.Answers = answers

	if strings.Contains(lang.Language, `de`) {

		if amountValue > 0 {
			data.TextAllGroups = `Alle Gruppen anzeigen`
			data.AmountToggle = -1
		} else {
			data.TextAllGroups = `Top 6 anzeigen`
			data.AmountToggle = 6
		}

		data.Basis.Name = NAME_DE
		data.Basis.Logo = LOGO_DE
		data.LangPos = LANG_DE
	} else {

		if amountValue > 0 {
			data.TextAllGroups = `Show all groups`
			data.AmountToggle = -1
		} else {
			data.TextAllGroups = `Show top 6`
			data.AmountToggle = 6
		}

		data.Basis.Name = NAME_EN
		data.Basis.Logo = LOGO_UK
		data.LangPos = LANG_EN
	}

	// Finally, execute the template
	Tools.SendChosenLanguage(response, lang)
	Templates.ProcessHTML(`results`, response, data)
}

// GetProgressState returns the css class representing the progress.
func (data PageResults) GetProgressState(influence int8) string {
	if influence > 0 {
		return ` progressitemdone`
	} else if influence < 0 {
		return ` progressitemundone`
	} else {
		return ``
	}
}

// GetGroupName returns the localized name of a product group by its index.
func (data PageResults) GetGroupName(xmlIndex int) string {
	return data.Groups[xmlIndex].GroupName.Names[data.LangPos].Text
}

// Lang returns the localized string using the language id.
func (data PageResults) Lang(strings []XML.String) string {
	return strings[data.LangPos].Text
}

// FormatAnswer returns localized string for a given answer by it's internal representation.
func (data PageResults) FormatAnswer(answer string) string {
	switch answer {
	case `1`:
		return data.Lang(data.Strings.AnswerYes)
	case `0`:
		return data.Lang(data.Strings.AnswerNo)
	case `*`:
		return data.Lang(data.Strings.AnswerSkipped)
	case `studentCount1`:
		return data.Lang(data.Strings.AnswerStudentCount1)
	case `studentCount2`:
		return data.Lang(data.Strings.AnswerStudentCount2)
	case `studentCount3`:
		return data.Lang(data.Strings.AnswerStudentCount3)
	case `studentCount4`:
		return data.Lang(data.Strings.AnswerStudentCount4)
	case `support4lecture`:
		return data.Lang(data.Strings.AnswerSupportLecture)
	case `replace`:
		return data.Lang(data.Strings.AnswerReplaceLecture)
	}

	Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityMiddle, LM.ImpactNone, LM.MessageNameREQUEST, `Unknown answer!`, answer)
	return ``
}
