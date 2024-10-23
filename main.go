package main

import (
	"PeerEdu-BackEnd/database"
	"PeerEdu-BackEnd/util/config"
)

func main() {
	config.Init()
	database.Init()
}
