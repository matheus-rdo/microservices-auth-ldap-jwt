package app

import (
	"log"

	env "github.com/Netflix/go-env"
	"github.com/subosito/gotenv"
)

// Environment variables
type Environment struct {
}

// ENV - output variable
var ENV Environment

func init() {
	gotenv.Load() // load .env file (if exists)
	if _, err := env.UnmarshalFromEnviron(&ENV); err != nil {
		log.Fatal("Fatal error unmarshalling environment config: ", err)
	}
}
