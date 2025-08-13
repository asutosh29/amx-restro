package session_utils

import (
	"encoding/gob"
	"fmt"
	"net/http"

	"github.com/asutosh29/amx-restro/pkg/types"
	"github.com/gorilla/sessions"
)

// For SSR will be removed while CSR
// var session_secret_bytes = []byte(config.SessionSecret)

// REMINDER: Passing secret via config causes Hash key not found
var session_secret_bytes = []byte("Very-secret")
var Store *sessions.CookieStore

// Store = sessions.NewCookieStore(session_secret_bytes)
func InitiateStructSession() {
	Store = sessions.NewCookieStore(session_secret_bytes)

	// TODO: Add Session ID and Order ID via this
	registerGobTypes(
		&types.Popup{},
	)
}

func registerGobTypes(values ...any) {
	for _, v := range values {
		gob.Register(v)
	}
}

func InsertPopupInFlash(w http.ResponseWriter, r *http.Request, object types.Popup) error {
	session, _ := Store.Get(r, "flash")

	session.AddFlash(object)

	err := session.Save(r, w)
	if err != nil {
		return fmt.Errorf("error saving the flash session : %v", err)
	}

	return nil
}

func ExtractPopupFromFlash(w http.ResponseWriter, r *http.Request) (types.Popup, error) {
	var nilPopup = types.Popup{}
	var popup = &types.Popup{}

	session, err := Store.Get(r, "flash")
	if err != nil {
		return nilPopup, fmt.Errorf("error getting data : %v", err)
	}

	if flashes := session.Flashes(); len(flashes) > 0 {
		var ok bool
		popup, ok = flashes[0].(*types.Popup)
		if !ok {
			return nilPopup, fmt.Errorf("error in deserialisation flash : %v", err)
		}
	}

	err = session.Save(r, w)
	if err != nil {
		return nilPopup, fmt.Errorf("error saving the session : %v", err)
	}

	return *popup, nil
}

func FlashMsgErr(w http.ResponseWriter, r *http.Request, Msg string, IsError bool) {
	tempPopup := types.Popup{
		Msg:     Msg,
		IsError: IsError,
	}
	err := InsertPopupInFlash(w, r, tempPopup)
	if err != nil {
		fmt.Println("Error Adding Flash messages:", err)
	}
}
