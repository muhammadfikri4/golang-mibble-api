package config

import (
	"context"
	"golang-prisma-api/prisma/db"

	"github.com/rs/zerolog/log"
)

func Client() (*db.PrismaClient, context.Context) {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		log.Fatal().Err(err).Msg("Error connecting to Prisma Client")
	}
	ctx := context.Background()
	return client, ctx
}
