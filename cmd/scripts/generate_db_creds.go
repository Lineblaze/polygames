package main

import (
	"fmt"
	"polygames/internal/pkg/config"
	"polygames/internal/pkg/wcrypto"
)

func main() {
	config.Read("./config.yml")

	username, password, err := wcrypto.EncodeUserPass("polygames", "polygames", config.Block)
	if err != nil {
		panic(err)
	}

	fmt.Println("Username:", username)
	fmt.Println("Password:", password)
}
