package models

type Track struct {
	ID           int     `json:"track_id" doc:"id of the track"`
	Name         string  `json:"name" doc:"name of the track"`
	AlbumID      int     `json:"album_id" doc:"id of the album corresponing to the track"`
	MediaTypeID  int     `json:"media_type_id" doc:"id of the media type corresponing to the track"`
	GenreID      int     `json:"genre_id" doc:"id of the genre corresponing to the track"`
	Composer     string  `json:"composer" doc:"composer?"`
	Milliseconds int     `json:"milliseconds" doc:"time of the track"`
	Bytes        int     `json:"bytes" doc:"bytes of the track"`
	UnitPrice    float64 `json:"unit_price" doc:"price"`
}

type Genre struct {
	ID   int    `json:"genre_id" doc:"id of the genre"`
	Name string `json:"name" doc:"genre name" maxLength:"120"`
}

type MediaType struct {
	ID   int    `json:"media_type_id" doc:"id of the media type"`
	Name string `json:"name" doc:"media type name" maxLength:"120"`
}
