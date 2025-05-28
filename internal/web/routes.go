package web

import (
	//"fmt"
	"html/template"
	//"io"
	"net/http"
	"os"
	"strings"
	log "github.com/sirupsen/logrus"
	"github.com/victorposada/gerrit-wui/internal/gerrit"
	"bytes"
)

type Change struct {
    Title string
    Done  bool
}
type TemplateData struct {
	GerritURL string
	Changes   interface{}
	CustomFieldName *string
	CustomFieldValue *string
}

type Board struct {
    ID string
	Description string
	Query string
	CustomFieldName *string
	CustomFieldValue *string
}

type TemplateBoards struct {
	Boards   []Board
}

func getBoards(w http.ResponseWriter, r *http.Request) {
	boards_env := os.Getenv("BOARDS")
	var boards []Board
	for _, board := range strings.Split(boards_env, ",") {
		description, exists_description := os.LookupEnv("DESCRIPTION_" + board)
		query, exists_query := os.LookupEnv("QUERY_" + board)
		custom_field_name, _ := os.LookupEnv("CUSTOM_FIELD_NAME_" + board)

		if !exists_description || !exists_query {
			log.Warn("Missing description on query for board "+ board)
			continue
		}
		boards = append(boards, Board{
			ID:               board,
			Description:      description,
			Query:            query,
			CustomFieldName:  &custom_field_name,
		})
	}

	log.Debug(boards)
	tmpl := template.Must(template.ParseFiles("templates/main.html"))
	log.Info("got / request")
	data := TemplateBoards{
		Boards:   boards,
	}
	tmpl.Execute(w, data)
}


func getBoard(board string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		GERRIT_URL := os.Getenv("GERRIT_URL")
		custom_field_name := os.Getenv("CUSTOM_FIELD_NAME_" + board)
		custom_field_template := (os.Getenv("CUSTOM_FIELD_VALUE_" + board))
		board_tmpl := template.Must(template.ParseFiles("templates/board.html"))

		var custom_field_value_html string
		tmpl, err := template.New("custom_value").Parse(custom_field_template)
		if err != nil {
			http.Error(w, "Error de plantilla (perfil)", http.StatusInternalServerError)
			return
		}
		buf := new(bytes.Buffer)
		tmpl.Execute(buf, custom_field_template)
		custom_field_value_html = buf.String()

		log.Info("got /" + board+ " request")
		query := os.Getenv("QUERY_" + board)
		log.Info("Query: " + query)
		changes := gerrit.GetChanges(query)
		data := TemplateData{
			GerritURL: GERRIT_URL,
			Changes:   changes,
			CustomFieldName: &custom_field_name,
			CustomFieldValue: &custom_field_value_html,
		}
		board_tmpl.Execute(w, data)
	}
}