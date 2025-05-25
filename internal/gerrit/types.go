package gerrit


type Project struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Change struct {
	ID       int    `json:"_number"`
	Project  string `json:"project"`
	Branch   string `json:"branch"`
	ChangeID string `json:"change_id"`
	Subject  string `json:"subject"`
}
