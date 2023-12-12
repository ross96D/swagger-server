package models

type Album struct {
	ID       int    `json:"album_id" doc:"id of the album"`
	Title    string `json:"tittle" doc:"title of the album" maxLength:"80"`
	ArtistID int    `json:"artist_id" doc:"id of the artist that create this album"`
}
