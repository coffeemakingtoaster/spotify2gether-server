package main

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	InvalidRouteError "spotify2gether.server/pkg/helpers"
	Logger "spotify2gether.server/pkg/helpers"
)

func handleWebsocketInput(w http.ResponseWriter, r *http.Request) {}

func main() {

	logger := Logger.Logger{Name: "main"}
	logger.Verbose("Starting")
	router := mux.NewRouter()

	logger.Verbose("initializing server...")
	router.HandleFunc("/", InvalidRouteError.InvalidRouteError)
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleWebsocketInput(w, r)
	})

	corsObj := handlers.AllowedOrigins([]string{"*"})
	logger.Verbose("Server online")
	err := http.ListenAndServe(":8080", handlers.CORS(corsObj)(router))
	if err != nil {
		logger.Error("Server failed shortly after startup" + err.Error())
	}
}
