package models

type ErrorMessage struct {
	Message string `json:"error"`
}

type Status400 struct {
	ErrorMessage
}

type Status404 struct {
	ErrorMessage
}

type Status500 struct {
	ErrorMessage
}

type SuccessMessage struct {
	Message string `json:"message"`
}
