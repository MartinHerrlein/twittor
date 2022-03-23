package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/MartinHerrlein/twittor/middlew"
	"github.com/MartinHerrlein/twittor/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores setea el puerto, el Handler y pone a escuchar el server*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
