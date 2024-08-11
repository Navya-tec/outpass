package main

import (
	"goprojects/outpass/db"
	"goprojects/outpass/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()

	server := gin.Default()

	server.POST("/signup", routes.CreateUser)

	server.POST("login", routes.Login)

	server.POST("request",routes.CreateRequest)

	server.GET("/users",routes.GetAllUser)

	server.GET("/requests",routes.GetAllRequests)

	server.GET("/request",routes.GetRequestByUserID)

	server.PATCH("/request/status",routes.UpdateStatus)

	server.GET("/request/status",routes.GetAllRequestsByStatus)

	server.Run(":8080")

}
