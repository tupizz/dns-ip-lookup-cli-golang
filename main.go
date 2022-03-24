package main

import (
	"cli-searchable-ip-dns/app"
	"log"
	"os"
)

func main() {
	aplicacao := app.Gerar()

	// without if init
	/*
		error := aplicacao.Run(os.Args)
		if error != nil {
			log.Fatal(error)
		}
	*/

	// if init
	if error := aplicacao.Run(os.Args); error != nil {
		log.Fatal(error)
	}

}
