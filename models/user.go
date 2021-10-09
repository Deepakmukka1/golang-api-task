package models

type User struct {
	Name string `json:"name"`
	City string `json:"city"`
	Age  int    `json:"age"`
}
