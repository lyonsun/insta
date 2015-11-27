package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(c *gin.Context) {

	session, err := store.Get(c.Request, "something-very-secret")
	if err != nil {
		http.Error(c.Writer, err.Error(), 500)
		return
	}

	token := session.Values["token"]

	fmt.Println("token[index]: ", token)

	c.HTML(200, "index.tmpl", gin.H{
		"title": "Main Page",
		"token": token,
	})
}
