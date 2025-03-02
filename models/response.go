package models

type Success struct {
	Status  int64  `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type Failed struct {
	Status  int64  `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}
