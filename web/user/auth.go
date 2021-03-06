package user

import (
	"net/http"

	"github.com/danilbushkov/experimental-web/conf"
)

//Login user
func Login(w http.ResponseWriter, r *http.Request) {
	session, _ := conf.Store.Get(r, "user")

	session.Values["auth"] = true
	session.Values["id"] = 1
	session.Save(r, w)
}

//Logout user
func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := conf.Store.Get(r, "user")

	session.Values["auth"] = false
	delete(session.Values, "id")
	session.Save(r, w)
}
