package main

import (
	"fmt"
	"strconv"
	"net/http"
	"os"
	"runtime"
	"github.com/gin-gonic/gin"
)

func endPointHandler (ctx *gin.Context){
	ctx.String(http.StatusOK, "%s; %s", ctx.Request.Method, ctx.Request.URL.Path)
}

func add(ctx *gin.Context){
	x, err := strconv.ParseFloat(ctx.Param("x"), 64)
	if err != nil{
		ctx.String(http.StatusBadRequest, "%s is not a number", ctx.Param("x"))
		return
	}

	y, err := strconv.ParseFloat(ctx.Param("y"), 64);
	if err != nil{
		ctx.String(http.StatusBadRequest,"%s is not a number", ctx.Param("y"))
		return
	}

	ctx.String(http.StatusOK, fmt.Sprintf("%f", x+y))
}

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
	r.GET("/info/:id", endPointHandler)
	r.GET("/add/:x/:y", add)
	// group endpoint
	ur := r.Group("/user")
	ur.	GET("/name", func (ctx *gin.Context){
		ctx.String(http.StatusOK, "baldazi")
	})
	r.Run(":" + os.Getenv("PORT"))
}