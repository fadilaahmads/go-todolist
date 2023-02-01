package models

type TaskModel struct {
	Id        int
	Title     string
	Email     string
	CreatedAt string
	UpdatedAt string
}

type WebResponse struct {
	Message int         `json:"message"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}
