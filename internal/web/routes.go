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
)

type Change struct {
    Title string
    Done  bool
}
type TemplateData struct {
	GerritURL string
	Changes   interface{}
}

type Board struct {
    ID int
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
	for i, board := range strings.Split(boards_env, ",") {
		description, exists_description := os.LookupEnv("DESCRIPTION_" + board)
		query, exists_query := os.LookupEnv("QUERY_" + board)
		custom_field_name, exists_custom_field_name := os.LookupEnv("CUSTOM_FIELD_NAME_" + board)
		custom_full_field_value, exists_custom_field_value := os.LookupEnv("CUSTOM_FIELD_VALUE_" + board)

		if !exists_description || !exists_query {
			log.Warn("Missing description on query for board "+ board)
			continue
		}

		// if (exists_custom_field_name || exists_custom_field_value) && !(exists_custom_field_name && exists_custom_field_value){
		// 	log.Warn("xor with custom values")
		// 	boards = append(boards, Board{
		// 		ID:               i,
		// 		Description:      description,
		// 		Query:            query,
		// 	})
		// } else{
		// 	boards = append(boards, Board{
		// 	ID:               i,
		// 	Description:      description,
		// 	Query:            query,
		// 	CustomFieldName:  &custom_field_name,
		// 	CustomFieldValue: &custom_full_field_value,
		// 	})
		// }
		boards = append(boards, Board{
			ID:               i,
			Description:      description,
			Query:            query,
			CustomFieldName:  &custom_field_name,
			CustomFieldValue: &custom_full_field_value,
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
		tmpl := template.Must(template.ParseFiles("templates/board.html"))

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