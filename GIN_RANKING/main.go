package main

import (
	"gin_ranking/router"
)

func main() {
	r := router.Router()
	r.Run(":9999")
}
