package main

import (
	"fmt"
	"net/http"
)

type Banner interface {
	GetText() string
}

type GuestBanner struct{}

type UserBanner struct{}

func (g GuestBanner) GetText() string {
	return "Please log in to see your balance"
}

func (u UserBanner) GetText() string {
	return "Welcome back, User!"
}

type PageHandler struct {
	Content Banner
}

func (p PageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, p.Content.GetText())
}

func main() {
	guestPage := PageHandler{Content: GuestBanner{}}
	userPage := PageHandler{Content: UserBanner{}}

	http.Handle("/guest", &guestPage)
	http.Handle("/user", &userPage)

	http.ListenAndServe(":8080", nil)
}
