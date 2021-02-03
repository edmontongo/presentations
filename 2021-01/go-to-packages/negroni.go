package main

import (
	"database/sql"
	"net/http"

	oauth2 "github.com/goincremental/negroni-oauth2"
	sessions "github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"github.com/urfave/negroni"
)

func NewOAuth(db *sql.DB, clientID, secret, callback string) *negroni.Negroni { // OMIT
	n := negroni.New(
		sessions.Sessions("my_session", cookiestore.New([]byte("secret123"))),
		oauth2.Google(&oauth2.Config{
			ClientID:     clientID,
			ClientSecret: secret,
			RedirectURL:  callback,
			Scopes:       []string{"email"},
		}),
		oauth2.LoginRequired(),
		negroni.NewRecovery(),
		auth.UserFilter{database.NewUsers(db)},
		log.NewHTTPLogger(),
	)
	http.ListenAndServe(":8080", n)
	return n // OMIT
} // OMIT
