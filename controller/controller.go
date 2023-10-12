package controller

import (
	"net/http"
	"signup/interfaces"
	"signup/models"

	"github.com/gin-gonic/gin"
)

type Signupcontroller struct{
	Signupservice interfaces.Isignup
}
func Initsignup(signupservice interfaces.Isignup)(Signupcontroller){
        return Signupcontroller{signupservice}
}

func(s *Signupcontroller)Signup(ctx *gin.Context){
	var user *models.Signup
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newacc,err:= s.Signupservice.SignUP(user)
	if err!=nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newacc})
}