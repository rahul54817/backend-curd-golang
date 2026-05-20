package routers

import (
	"go-backend/handler"
	"net/http"
)

func RegisterUserRoutes(handler *handler.UserHandler) {

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.CreateUser(w, r)
		case http.MethodDelete:
			handler.DeleteUser(w, r)
		case http.MethodGet:
			id := r.URL.Query().Get("id")
			if id != "" {
				handler.GetUserById(w, r)
			} else {
				handler.GetUserList(w, r)
			}
		}
	})
}
