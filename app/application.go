package app

import (
	"github.com/gin-gonic/gin"
	"github.com/helyus1412/todo-list-api/configs"
	projectdb "github.com/helyus1412/todo-list-api/datasource/mysql/project_db"
)

var (
	router = gin.Default()
)

func StartApp() {
	projectdb.Connect()
	port := configs.EnvFile.AppPort
	router.Run(":" + port)
}
