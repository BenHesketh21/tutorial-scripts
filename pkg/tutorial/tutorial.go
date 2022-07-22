package tutorial

type Prerequisite struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Version     string `json:"version"`
	Alternative string `json:"alternative"`
}

type Tutorial struct {
	Name          string         `json:"name"`
	Prerequisites []Prerequisite `json:"prerequisites"`
	Steps         []interface{}  `json:"steps"`
}
