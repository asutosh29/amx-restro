package session_utils

import "github.com/gorilla/sessions"

// TODO: add this to ENV
var Store = sessions.NewCookieStore([]byte("session-secret"))
