package routers

import (
	handler "mongo-server/Handler"
	"net/http"
)

func UserRouter() {
	http.HandleFunc("/user/", Cors(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetUserById(w, r)
		case http.MethodPut:
			handler.UpdateUserById(w, r)
		case http.MethodDelete:
			handler.DeleteUserById(w, r)
		default:
			http.Error(w, "Method not allowed!!!", http.StatusMethodNotAllowed)
			return
		}
	}))

	usersHandler := Cors(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetAllUsers(w, r)
		case http.MethodPost:
			handler.CreateNewUser(w, r)
		default:
			http.Error(w, "Method not allowed!!!!", http.StatusMethodNotAllowed)
			return
		}
	})

	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/users/", usersHandler)
}
