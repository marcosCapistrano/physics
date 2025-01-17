// main.go
package main

import (
	"log"
)

func main() {
	app := &application{}

	if err := app.start(); err != nil {
		log.Fatal(err)
	}

	app.run()
	app.quit()
}
