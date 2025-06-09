package main

import (
	"net/http"
	"os"
	"runtime"
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	r.GET("/", func (ctx *gin.Context){
		ctx.String(http.StatusOK, "hello");
	})
	r.GET("/greeting", func (ctx *gin.Context){
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hello user",
		})
	})
	r.GET("/os", func (ctx *gin.Context){
		ctx.JSON(http.StatusOK, gin.H{
			"os": runtime.GOOS,
		})
	})
	r.Run(":" + os.Getenv("PORT"))
}