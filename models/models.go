package models

import "time"

type TaskModel struct {
	Id        int
	Title     string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type WebResponse struct {
	Message int         `json:"message"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}
