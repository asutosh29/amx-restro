package session_utils

import (
	"github.com/asutosh29/amx-restro/pkg/utils/config"
	"github.com/gorilla/sessions"
)

// For SSR will be removed while CSR
var session_secret = []byte(config.SessionSecret)
var Store = sessions.NewCookieStore(session_secret)

// Store = sessions.NewCookieStore(session_secret)
