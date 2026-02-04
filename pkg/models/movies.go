package models

type Movie struct {
	Id          int          `json:"id" db:"id"`
	UserID      int          `json:"-" db:"-"`
	Title       string       `json:"title" binding:"required" db:"title"`
	Type        string       `json:"type" binding:"required,oneof=film serial" db:"type"`
	Year        int          `json:"year" binding:"required" db:"year"`
	Description string       `json:"description" db:"description"`
	Status      string       `json:"status" binding:"required,oneof=watched planned" db:"status"`
	Mark        *int         `json:"mark" binding:"min=1,max=10" db:"mark"`
	Review      *string      `json:"review" db:"review"`
	Actors      []MovieActor `json:"actors"`
	PosterURL   string       `json:"poster_url" db:"poster_url"`
}

type MovieActor struct {
	Id        int    `json:"id" db:"id"`
	FirstName string `json:"first_name" binding:"required" db:"first_name"`
	LastName  string `json:"last_name" binding:"required" db:"last_name"`
	Role      string `json:"role" binding:"required" db:"role"`
	BioUrl    string `json:"bio_url" db:"bio_url"`
}
