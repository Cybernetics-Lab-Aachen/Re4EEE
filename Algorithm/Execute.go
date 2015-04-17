package Algorithm

import (
	"fmt"
	"github.com/SommerEngineering/Re4EEE/DB/Scheme"
	"github.com/SommerEngineering/Re4EEE/XML"
	"sort"
)

func ExecuteAnswers(answers Scheme.Answers) (result Scheme.ProductGroups) {

	data := XML.GetData()
	groups := make(Scheme.ProductGroups, 16)

	// Algorithm:
	for n, productGroup := range data.ProductsCollection.Products {

		/*  1 */ groups[n].Points = groups[n].Points + kindConditionalPresence(answers.A1Data, productGroup.SharedProperties.VideoContent)
		/*  2 */ groups[n].Points = groups[n].Points + kindConditionalPossibility(answers.A2Data, productGroup.SharedProperties.Assistant)
		/*  3 */ groups[n].Points = groups[n].Points + kindConditionalPossibility(answers.A3Data, productGroup.SharedProperties.UserComments)
		/*  4 */ groups[n].Points = groups[n].Points + kindConditionalPresence(answers.A4Data, productGroup.SharedProperties.LiveCollaboration)
		/*  5 */ groups[n].Points = groups[n].Points + kindConditionalPresence(answers.A5Data, productGroup.SharedProperties.CommunityCollaboration)
		/*  6 */ groups[n].Points = groups[n].Points + kindAppropriateCountStudents(answers.A6Data, productGroup.SharedProperties.AppropriateCountStudents)
		/*  7 */ groups[n].Points = groups[n].Points + kindConditionalPresence(answers.A7Data, productGroup.SharedProperties.Downloads)
		/*  8 */ groups[n].Points = groups[n].Points + kindConditionalYesNo(answers.A8Data, productGroup.SharedProperties.Possibility2DeclareLearningObjectives)
		/*  9 */ groups[n].Points = groups[n].Points + kindOperationType(answers.A9Data, productGroup.SharedProperties.OperationType)
		/* 10 */ groups[n].Points = groups[n].Points + kindConditionalYesNo(answers.A10Data, productGroup.SharedProperties.CloudBased)
		/* 11 */ groups[n].Points = groups[n].Points + kindConditionalYesNo(answers.A11Data, productGroup.SharedProperties.Intranet)
		/* 12 */ groups[n].Points = groups[n].Points + kindConditionalPresence(answers.A12Data, productGroup.SharedProperties.Exam)
		/* 13 */ groups[n].Points = groups[n].Points + kindConditionalPresence(answers.A13Data, productGroup.SharedProperties.StudentRoles)
		/* 14 */ groups[n].Points = groups[n].Points + kindConditionalYesNo(answers.A14Data, productGroup.SharedProperties.DisplayEquations)
		/* 15 */ groups[n].Points = groups[n].Points + kindConditionalYesNo(answers.A15Data, productGroup.SharedProperties.WriteEquations)
		/* 16 */ groups[n].Points = groups[n].Points + kindConditionalPossibility(answers.A16Data, productGroup.SharedProperties.TeachingTypePresentation)
		/* 17 */ groups[n].Points = groups[n].Points + kindConditionalPossibility(answers.A17Data, productGroup.SharedProperties.TeachingTypeDevelopment)
		/* 18 */ groups[n].Points = groups[n].Points + kindConditionalPossibility(answers.A18Data, productGroup.SharedProperties.TeachingTypeExplorative)

		groups[n].Name = productGroup.InternalName
		groups[n].XMLIndex = n
	}

	//
	// Re-align the results to respect the range from 0-100%:
	//
	sort.Sort(groups)
	worstPoints := groups[len(groups)-1].Points
	correctionPoints := worstPoints * -1
	bestPointsCorrected := float64(18 + correctionPoints)
	bestPointsNormal := float64(18)

	for n, _ := range groups {
		if worstPoints < 0 {
			groups[n].Points += correctionPoints
			result := (float64(groups[n].Points) / bestPointsCorrected) * 100.0
			groups[n].Percent = fmt.Sprintf("%.f", result)
		} else {
			result := (float64(groups[n].Points) / bestPointsNormal) * 100.0
			groups[n].Percent = fmt.Sprintf("%.f", result)
		}
	}

	result = groups //[0-6]
	return
}
