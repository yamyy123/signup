package routes

import ("github.com/gin-gonic/gin"
"signup/controller")

func Router(router *gin.Engine, controller controller.Signupcontroller)  {
	router.Static("/static", "./")
	router.POST("/signup", controller.Signup)
}
