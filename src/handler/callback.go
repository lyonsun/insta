package handler

import (
	"config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	// "log"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func CallbackHandler(c *gin.Context) {

	session, err := store.Get(c.Request, "something-very-secret")
	if err != nil {
		http.Error(c.Writer, err.Error(), 500)
		return
	}

	token := session.Values["token"]

	fmt.Println("token[callback]: ", token)

	if token != nil {
		c.Redirect(http.StatusMovedPermanently, "/")
		return
	}

	conf := &oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		RedirectURL:  config.RedirectURL,
		// Scopes:       []string{"SCOPE1", "SCOPE2"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  config.AuthURL,
			TokenURL: config.TokenURL,
		},
	}

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL("otrack_v2")
	fmt.Println("Visit the URL for the auth dialog: %v", url)

	code := c.Query("code")

	if code == "" {
		var authorize_url = conf.Endpoint.AuthURL + "?client_id=" + conf.ClientID + "&redirect_uri=" + conf.RedirectURL + "&response_type=" + config.ResponseType

		c.Redirect(http.StatusMovedPermanently, authorize_url)
		return
	}

	fmt.Println("code: ", code)

	tok, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		// log.Fatal(err)
		fmt.Println("error: ", err)
	} else {
		session.Values["token"] = tok.AccessToken

		session.Save(c.Request, c.Writer)
	}

	c.Redirect(http.StatusMovedPermanently, "/")
}
