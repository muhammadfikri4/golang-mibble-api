package main

import (
	"golang-prisma-api/src/config"
	"golang-prisma-api/src/router"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

func main() {
	client, err := config.DBConfig()
	if err != nil {
		log.Info().Msg(err.Error())
		panic(err)
	}
	log.Info().Msg("Connected to Prisma on localhost:4466")
	defer config.CloseDB(client)

	app := gin.Default()
	router.UserRoute(app)
	app.Run(":3000")
}
