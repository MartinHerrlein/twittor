package main

import (
	"log"

	"github.com/MartinHerrlein/twittor/bd"
	"github.com/MartinHerrlein/twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
