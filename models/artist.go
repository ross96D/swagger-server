package models

type Artist struct {
	ID   int    `json:"artist_id" doc:"id of the artist"`
	Name string `json:"name" doc:"name of the artist"`
}
