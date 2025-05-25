package web

import (
	//"fmt"
	"html/template"
	//"io"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/victorposada/gerrit-wui/internal/gerrit"
)

type Change struct {
    Title string
    Done  bool
}
type TemplateData struct {
	GerritURL string
	Changes   interface{}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	GERRIT_URL := os.Getenv("GERRIT_URL")
	tmpl := template.Must(template.ParseFiles("templates/main.html"))
	log.Info("got / request")
	changes := gerrit.GetChanges(os.Getenv("DEFAULT_QUERY"))

	data := TemplateData{
		GerritURL: GERRIT_URL,
		Changes:   changes,
	}

	tmpl.Execute(w, data)
}

func getBoard(board string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		GERRIT_URL := os.Getenv("GERRIT_URL")
		tmpl := template.Must(template.ParseFiles("templates/main.html"))

		log.Info("got /" + board+ " request")
		query := os.Getenv("QUERY_" + board)
		log.Info("Query: " + query)
		changes := gerrit.GetChanges(query)
		data := TemplateData{
			GerritURL: GERRIT_URL,
			Changes:   changes,
		}
		tmpl.Execute(w, data)
	}
}