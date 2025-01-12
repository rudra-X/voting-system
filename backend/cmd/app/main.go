package main

import (
	"os"
	"voiting-system/internal/routers"
	"voiting-system/internal/utils"
)

var GlobalConfigs utils.Config

func main() {
	loadConfig(&GlobalConfigs)
	router := routers.SetupRouter()
	router.Run("localhost:8080")
}

func loadConfig(globalConfigs *utils.Config) {
	globalConfigs.JwtSecretKey = os.Getenv("JWT_SECRET_KEY")
}
