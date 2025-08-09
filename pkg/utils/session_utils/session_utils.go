package session_utils

import (
	"github.com/asutosh29/amx-restro/pkg/utils/config"
	"github.com/gorilla/sessions"
)

var session_secret = []byte(config.SessionSecret)
var Store = sessions.NewCookieStore([]byte(session_secret))
