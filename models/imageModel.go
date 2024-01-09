package models

type Image struct {
	PublicID string `json:"publicId"`
	Tags     string `json:"tags"`
	Source   string `json:"src"`
	Filename string `json:"filename"`
}
