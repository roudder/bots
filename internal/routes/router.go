package routes

import (
	. "bots/internal/handlers"
	"bots/internal/timer"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func StartServer() {
	//perhaps need to add cache to server structure
	server := &http.Server{
		Addr:    ":8080",
		Handler: InitRoute(),
	}
	go timer.StartTimer()

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
}

func InitRoute() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/count", GetBotsCount).Methods("GET")
	router.HandleFunc("/", UpdateBotsCount).Methods("POST")
	return router
}
