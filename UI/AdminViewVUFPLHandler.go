package UI

import "net/http"
import "github.com/SommerEngineering/Ocean/Templates"

func AdminViewVUFPLHandler(response http.ResponseWriter, request *http.Request) {
	data := UIData{}
	data.AdditionalHeaders = `<script src='staticFiles/javascript/common.js' type='text/javascript'></script>
							  <script src='staticFiles/javascript/editor.js' type='text/javascript'></script>`
	data.Version = `1.0`
	data.DropdownActions = ``
	data.DropdownActionItems = `<li><a href="#" onclick="Block.save();"><span class="glyphicon glyphicon-floppy-save"></span> Regelwerk speichern</a></li>
            					<li role="presentation" class="divider"></li>
            					<li><a href="#" onclick="new Block('START');"><span class="glyphicon glyphicon-plus"></span> Regelwerk: Startpunkt hinzufügen</a></li>
            					<li><a href="#" onclick="new Block('END');"><span class="glyphicon glyphicon-plus"></span> Regelwerk: Ende hinzufügen</a></li>
					            <li><a href="#" onclick="new Block('YESNO');"><span class="glyphicon glyphicon-plus"></span> Regelwerk: Ja/Nein-Frage hinzufügen</a></li>
					            <li><a href="#" onclick="new Block('NUMBER');"><span class="glyphicon glyphicon-plus"></span> Regelwerk: Wertabfrage hinzufügen</a></li>
					            <li role="presentation" class="divider"></li>
					            <li><a href="#" onclick="Block.deleteCurrentBlock();"><span class="glyphicon glyphicon-trash"></span> Markierten Block löschen</a></li>
					            <li><a href="#" onclick="Block.editCurrentBlock();"><span class="glyphicon glyphicon-pencil"></span> Markierten Block bearbeiten</a></li>`

	Templates.ProcessHTML(`admin`, response, data)
}
