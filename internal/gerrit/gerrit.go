package gerrit

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"os"
	log "github.com/sirupsen/logrus"
)

// func GetProjects() (map[string]Repositorio, error) {
// 	repos, err := GetRequest("a/projects/")
// 	return repos, err
// }

func GetChanges(query string) (change_list []Change) {
	data, _ := GetRequest("/a/changes/?q=" + query)

	var changes []Change
	if err := json.Unmarshal([]byte(data), &changes); err != nil {
		log.Info("Error parsing JSON:", err)
		return
	}

	// for _, change := range changes {
	// 	fmt.Println(change)
	// }

	return changes
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

	log.Debug(resp.Body)
	log.Debug(resp.StatusCode)
	log.Debug(resp)

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	bodyStr := string(bodyBytes)
	bodyStr = strings.TrimPrefix(bodyStr, ")]}'\n")

	log.Debug("HTTP response body: " + bodyStr)

	return bodyStr, err
}
