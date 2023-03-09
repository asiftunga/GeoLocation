package entities

type Coordinate struct {
	Long float64 `json:"lng"`
	Lat  float64 `json:"lat"`
}

type Polygon struct {
	ID          string       `json:"id"`
	Name        string       `json:"name,omitempty"`
	Description string       `json:"description,omitempty"`
	Comment     string       `json:"comment,omitempty"`
	UserID      int64        `json:"user_id,omitempty" swaggerignore:"true"`
	Coordinates []Coordinate `json:"coordinates,omitempty"` //two dimensional array for postgresql
}

// ======================ust taraf karsilayan======================alt kisim ise response donduren=============================
//type Cordinate [2]float64

type Polygon2 struct {
	ID          string       `json:"id"`
	Name        string       `json:"name,omitempty"`
	Description string       `json:"description,omitempty"`
	Comment     string       `json:"comment,omitempty"`
	UserID      int64        `json:"user_id,omitempty"`
	Coordinates [][2]float64 `json:"coordinates,omitempty"`
}

// DeletePolygonOnePoint represents the request body for deleting a polygon.
type DeletePolygonOnePoint struct {
	ID          string       `json:"id"`
	Coordinates []Coordinate `json:"coordinates,omitempty" swaggerignore:"true"`
}

// until solutions find it is the best way to use like this
type DeletePolygonOnePoint1 struct {
	ID          string       `json:"id"`
	Coordinates []Coordinate `json:"coordinates,omitempty"`
}
