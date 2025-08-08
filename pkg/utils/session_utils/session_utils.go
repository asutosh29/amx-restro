package session_utils

import (
	"os"

	"github.com/gorilla/sessions"
)

var session_secret = []byte(os.Getenv("SESSION_SECRET"))
var Store = sessions.NewCookieStore([]byte(session_secret))
