package main

type ErrorResponse struct {
	Status bool   `json:"Status"`
	Msg    string `json:"Msg"`
}
