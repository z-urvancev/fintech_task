package models

type Link struct {
	Id    uint   `json:"-"`
	Url   string `json:"url"`
	Short string `json:"short"`
}
