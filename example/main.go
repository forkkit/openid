package main

import (
	"github.com/ShaleApps/openid"
	"log"
	"net/http"
	"os"
)

func main() {
	config, err := openid.NewConfig(&openid.Opts{
		DiscoveryUrl:    os.Getenv("OPENID_TEST_DISCOVERY_URL"),
		ClientID:        os.Getenv("OPENID_TEST_CLIENT_ID"),
		ClientSecret:    os.Getenv("OPENID_TEST_CLIENT_SECRET"),
		Redirect:        os.Getenv("OPENID_TEST_REDIRECT"),
		Scopes:          openid.DefaultScopes,
		SkipIssuerCheck: true,
		SessionSecret:   os.Getenv("OPENID_TEST_SESSION_SECRET"),
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	mux := http.NewServeMux()
	///login/authorization redirects the user to login to the identity provider
	mux.HandleFunc("/login/authorization", config.AuthorizationRedirect())
	///mock home
	//mux.HandleFunc("/", config.Middleware(func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("hello!"))
	//}))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello!"))
	})
	mux.HandleFunc("/login", config.HandleLogin(func(w http.ResponseWriter, r *http.Request, usr *openid.User) error {
		log.Print(usr.String())
		return nil
	}, "/"))
	http.ListenAndServe(":8080", mux)
}
