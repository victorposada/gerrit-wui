package gerrit

import (
	//"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"os"
	log "github.com/sirupsen/logrus"
)

type Project struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Change struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}


// func GetProjects() (map[string]Repositorio, error) {
// 	repos, err := GetRequest("a/projects/")
// 	return repos, err
// }

func GetChanges(){
	changes, _ := GetRequest("a/changes/?q=status:open+limit:1")
	fmt.Println(changes)
}

func GetRequest(path string) (body string, err error) {

	gerrit_url := os.Getenv("GERRIT_URL") + path
	gerrit_user := os.Getenv("GERRIT_USER")
	gerrit_token := os.Getenv("GERRIT_TOKEN")

	log.Debug("Get request to " + gerrit_url)
	req, err := http.NewRequest("GET", gerrit_url, nil)
	if err != nil {
		return "", err
	}

	req.SetBasicAuth(gerrit_user, gerrit_token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	bodyStr := string(bodyBytes)
	bodyStr = strings.TrimPrefix(bodyStr, ")]}'\n")
	return bodyStr, err
}
