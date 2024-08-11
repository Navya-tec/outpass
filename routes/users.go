package routes

import (
	"goprojects/outpass/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(context *gin.Context) {

	var user models.User
	err:=context.ShouldBindJSON(&user)
	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Error in parsing user data"})
		return 
	}

	err=user.SaveUSer()
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Error in saving user!"})
		return
	}

	context.JSON(http.StatusOK,gin.H{"message":"User created successfully","user":user})
}

func Login(context *gin.Context){

	var user models.User
	err:=context.ShouldBindJSON(&user)
	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Error in parsing user data"})
		return 
	}

	err=user.ValidateCredentials()
	if err!=nil{
		context.JSON(http.StatusUnauthorized,gin.H{"message":"INVALID CREDENTIALS"})
		return
	}

	context.JSON(http.StatusOK,gin.H{"message":"LOGIN successfull"})
}

func GetAllUser(context *gin.Context){
	// var users []models.User

	users,err:=models.GetUsers()
	if err!=nil{
       context.JSON(http.StatusInternalServerError,gin.H{"message":"Couldn't fetch users!!"})
	   return
	}

	context.JSON(http.StatusOK,gin.H{"message":"successfull","users":users})
}