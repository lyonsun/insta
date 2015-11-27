package main

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"handler"
	"net/http"
	"time"
)

func main() {

	/*
	 * Configurations
	 */

	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	/*
	 * Routers
	 */

	r.GET("/", handler.IndexHandler)
	r.GET("/ping", pong)
	r.GET("/callback", handler.CallbackHandler)
	r.GET("/account", handler.UserInfoHandler)
	r.GET("/recent", handler.RecentHandler)

	/*
	 * Run server
	 */

	s := &http.Server{
		Addr:           ":80",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}

func pong(c *gin.Context) {
	if false {
		c.JSON(200, gin.H{"result": "pong"})
		return
	}
	c.JSON(404, gin.H{"error": "something wrong"})
	// c.String(200, "pong")
}
