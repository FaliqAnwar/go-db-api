package main

import (
	"go-db-api/internal/utils/confloader"
)

func main() {
	_, err := confloader.MustLoad("go-db-api")
	if err != nil {
		panic(err)
	}

}
