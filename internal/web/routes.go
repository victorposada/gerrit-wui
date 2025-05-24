package web

import (
	//"fmt"
	"html/template"
	"io"
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
	changes := gerrit.GetChanges()

	data := TemplateData{
		GerritURL: GERRIT_URL,
		Changes:   changes,
	}
	tmpl.Execute(w, data)
}

func getInfo(w http.ResponseWriter, r *http.Request) {
	log.Info("got /info request")
	io.WriteString(w, "Hello, HTTP!\n")
}