package routers

import (
	Handler "mongo-server/Handler"
	"net/http"
)

func OrderRouter() {
	http.HandleFunc("/orders", Cors(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			Handler.GetAllOrders(w, r)
		case http.MethodPost:
			Handler.CreateNewOrder(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))
	http.HandleFunc("/orders/", Cors(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			Handler.GetOrderById(w, r)
		case http.MethodPut:
			Handler.UpdateOrderById(w, r)
		case http.MethodDelete:
			Handler.DeleteOrderById(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))
}
