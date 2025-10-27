package main

import (
	"context"
	"fmt"
	"net/http"
)

type userKey int // Custom type for context key

const authenticatedUserKey userKey = 0

type User struct {
	ID    string
	Email string
	Roles []string
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate authentication
		authenticatedUser := User{
			ID:    "user123",
			Email: "test@example.com",
			Roles: []string{"admin", "user"},
		}

		// Store user in context
		ctx := context.WithValue(r.Context(), authenticatedUserKey, authenticatedUser)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve user from context
	user, ok := r.Context().Value(authenticatedUserKey).(User)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	fmt.Fprintf(w, "Welcome, %s! Your roles: %v", user.Email, user.Roles)
}

func main2() {
	mux := http.NewServeMux()
	mux.Handle("/protected", AuthMiddleware(http.HandlerFunc(ProtectedHandler)))
	http.ListenAndServe(":8080", mux)
}
