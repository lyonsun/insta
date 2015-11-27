package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	// "github.com/leonmaia/requests"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	API_URL = "https://api.instagram.com/v1"
)

func UserInfoHandler(c *gin.Context) {

	session, err := store.Get(c.Request, "something-very-secret")
	if err != nil {
		http.Error(c.Writer, err.Error(), 500)
		return
	}

	token := session.Values["token"]

	fmt.Println("token[user]: ", token)

	request_uri := API_URL + "/users/self/?access_token=" + token.(string)

	resp, err := http.Get(request_uri)

	if err != nil {
		fmt.Println("error: ", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var data UserInfo

	err = json.Unmarshal(body, &data)

	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println("data: ", data.Data)

	c.HTML(200, "account.tmpl", gin.H{
		"title": "Main Page",
		"data":  data.Data,
	})

	return
}

func RecentHandler(c *gin.Context) {

	session, err := store.Get(c.Request, "something-very-secret")
	if err != nil {
		http.Error(c.Writer, err.Error(), 500)
		return
	}

	token := session.Values["token"]

	fmt.Println("token[user]: ", token)

	request_uri := API_URL + "/users/self/media/recent/?access_token=" + token.(string)

	resp, err := http.Get(request_uri)

	if err != nil {
		fmt.Println("error: ", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var dat map[string]interface{}

	if err := json.Unmarshal(body, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// c.HTML(200, "account.tmpl", gin.H{
	// 	"title": "Main Page",
	// 	"data":  data.Data,
	// })

	return
}
