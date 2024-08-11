package routes

import (
	"fmt"
	"goprojects/outpass/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateRequest(context *gin.Context) {

	var request models.Request
	err := context.ShouldBindJSON(&request)
	if err != nil {
		
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't Parse Request Data!"})
		return
	}

	err = request.SaveRequest()
	if err != nil {
		log.Println("Error Occured", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error in saving request data"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Request created successfully"})
}

func GetRequestByUserID(context *gin.Context){

	user_id:=context.Query("id")
	userId,err:=strconv.Atoi(user_id)
	if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }
	
	fmt.Println(userId)
    reqs,err:=models.GetReqByUserId(userId)
	if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching requests"})
        return
    }
 
    context.JSON(http.StatusOK, gin.H{"requests": reqs})

}

func GetAllRequests(context *gin.Context){
	reqs,err:=models.GetAllRequest()
	if err!=nil{
       context.JSON(http.StatusInternalServerError,gin.H{"message":"Couldn't fetch requests!!"})
	   return
	}

	context.JSON(http.StatusOK,gin.H{"message":"successfull","request":reqs})
}

func UpdateStatus(context *gin.Context){

	r_id:=context.Query("id")
	status:=context.Query("status")

	if status!="completed"&&status!="declined"&&status!="approved"{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Invalid Status"})
		return
	}

	rId,err:=strconv.Atoi(r_id)
	if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request ID"})
        return
    }

	err=models.UpdateRequestStatus(rId,status)
	if err!=nil{
       context.JSON(http.StatusInternalServerError,gin.H{"message":"Couldn't Update Request Status"})
	   return
	}

	context.JSON(http.StatusOK,gin.H{"message":"Status Updated!"})
}

func GetAllRequestsByStatus(context *gin.Context){

	status:=context.Query("status")
	if status!="completed"&&status!="declined"&&status!="approved"&&status!="pending"{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Invalid Status"})
		return
	}

	reqs,err:=models.GetRequestByStatus(status)
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Error in fetching all requests by status"})
		return
	}

	context.JSON(http.StatusOK,gin.H{"requests":reqs})
}
