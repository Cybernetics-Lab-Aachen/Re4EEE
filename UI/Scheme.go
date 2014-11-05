package UI

import "html/template"

type UIData struct {
	Version             string
	Content             string
	DropdownActionItems template.HTML
	DropdownActions     template.HTML
	AdditionalHeaders   template.HTML
}
