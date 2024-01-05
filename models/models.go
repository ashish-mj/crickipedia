package models

type Player struct {
	Id string         `json:"id"`
	Name string       `json:"name"`
	Age int           `json:"age"`
	Role string       `json:"role"`
	Contract string   `json:"contract"`
	Jersey int        `json:"jersey"`
}