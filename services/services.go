package services

import (
	"go-todolist/databases"
	"go-todolist/models"
	"go-todolist/utils"

	"github.com/gin-gonic/gin"
)

type ServicesInterface interface {
	GetAll(c *gin.Context)
	GetOne(c *gin.Context, id int)
	Create(c *gin.Context)
	Update(c *gin.Context, id int)
	Delete(c *gin.Context, id int)
}

func GetAll(c *gin.Context) []models.TaskModel {
	db, err := databases.DBConnect()
	utils.PanicIfError(err)
	defer db.Close()

	sql := "SELECT id, title, email FROM tasks"
	result, err := db.Query(sql)
	utils.PanicIfError(err)

	var tasks []models.TaskModel

	for result.Next() {
		var task models.TaskModel
		err := result.Scan(&task.Id, &task.Title, &task.Email)
		utils.PanicIfError(err)
		tasks = append(tasks, task)
	}
	return tasks

}

func GetOne(c *gin.Context, taskId int) models.TaskModel {
	// TODO
	// Create exception for out of bound index
	db, err := databases.DBConnect()
	utils.PanicIfError(err)

	stmt, err := db.Prepare("SELECT * FROM tasks WHERE id=?")
	defer stmt.Close()
	utils.PanicIfError(err)

	var task models.TaskModel
	err = stmt.QueryRow(taskId).Scan(&task.Id, &task.Title, &task.Email, &task.CreatedAt, &task.UpdatedAt)
	utils.PanicIfError(err)
	return task
}
