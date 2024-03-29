package helpers

import (
	"fmt"
	"net/http"
)

func InvalidRouteError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Invalid Route. If you are trying to reach the game API please interact with api.speedrun.io")
}

func ConnectionNotWebsocketError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusForbidden)
	fmt.Fprintf(w, "Error: Connection to the /ws part of the gameserver should only be via websockets")
}

//Raise Error and send to client
func InvalidRequestError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "Error: The request is invalid in this context")
}
