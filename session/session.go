package session

import (
	"blog-Go_SR/utils"
	"net/http"
	"time"

	"github.com/go-martini/martini"
)

const (
	COOKIE_NAME = "sessionId"
)

type Session struct {
	Id           string
	Username     string
	IsAuthorized bool
}
type SessionStore struct {
	data map[string]*Session
}

func (store *SessionStore) Get(sessionId string) *Session {
	session := store.data[sessionId]
	if session == nil {
		return &Session{Id: sessionId}
	}
	return session
}
func (store *SessionStore) Set(session *Session) {
	store.data[session.Id] = session
}
func NewSessionStore() *SessionStore {
	s := new(SessionStore)
	s.data = make(map[string]*Session)
	return s
}
func ensureCookie(r *http.Request, w http.ResponseWriter) string {
	cookie, _ := r.Cookie(COOKIE_NAME)
	if cookie != nil {
		return cookie.Value
	}
	sessionId := utils.GenerateId()

	cookie = &http.Cookie{
		Name:    COOKIE_NAME,
		Value:   sessionId,
		Expires: time.Now().Add(5 * time.Minute),
	}
	http.SetCookie(w, cookie)
	return sessionId
}

var sessionStore = NewSessionStore()

func Middleware(ctx martini.Context, r *http.Request, w http.ResponseWriter) {
	sessionId := ensureCookie(r, w)
	session := sessionStore.Get(sessionId)
	ctx.Map(session)
	ctx.Next()
	sessionStore.Set(session)
}
