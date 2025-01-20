package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Session struct {
	ID            string
	Authenticated bool
	Creation      time.Time
	Duration      time.Time
}

func (s *Session) Init() {
	b := make([]byte, 64)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%x", b)
	s.ID = base64.URLEncoding.EncodeToString(b)
}

func (s *Session) IsAuthenticated() bool {
	return s.Authenticated
}

func (s *Session) GetID() string {
	return s.ID
}

func (s *Session) GenerateCookie(name string, expiration int64) http.Cookie {
	cookie := http.Cookie{
		Name:     name,
		Value:    s.ID,
		Expires:  s.Creation.Add(time.Second * 12), // 1 day expiration
		HttpOnly: true,                             // Prevents JavaScript access to the cookie
	}

	return cookie
}

// Session attribute set
// Session attribute get

func main() {
	fmt.Println("Hello World")
	session := &Session{"1", false, time.Now(), time.Month} // 30 days
	session.IsAuthenticated()
}
