package shared

import "github.com/gorilla/sessions"

var Store = sessions.NewCookieStore(key)
