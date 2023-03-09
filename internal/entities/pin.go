package entities

type Piin struct {
	Long float64 `json:"lng"`
	Lat  float64 `json:"lat"`
}

type Pin struct {
	ID          string `json:"id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Comment     string `json:"comment,omitempty"`
	UserID      int64  `json:"user_id,omitempty"`
	Coordinates []Piin `json:"coordinates,omitempty"` //two dimensional array for postgresql
}

// ======================ust taraf karsilayan======================alt kisim ise response donduren=============================
//type Cordinate [2]float64

type Pin2 struct {
	ID          string       `json:"id"`
	Name        string       `json:"name,omitempty"`
	Description string       `json:"description,omitempty"`
	Comment     string       `json:"comment,omitempty"`
	UserID      int64        `json:"user_id,omitempty"`
	Coordinates [][2]float64 `json:"coordinates,omitempty"`
}
