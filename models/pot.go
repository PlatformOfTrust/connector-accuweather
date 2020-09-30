package models

type POTContext struct {
	Context struct {
		Version float32 `json:"@version"`
		Schema  string  `json:"@schema"`
	} `json:"@context"`
}
