package config

import (
	"golang-prisma-api/prisma/db"

	"github.com/rs/zerolog/log"
)

func DBConfig() (*db.PrismaClient, error) {
	// Inisialisasi Prisma client
	client := db.NewClient()

	// Menghubungkan Prisma client ke database
	if err := client.Prisma.Connect(); err != nil {
		return nil, err
	}

	// Mengembalikan client ke fungsi pemanggil (main)
	return client, nil
}

func CloseDB(client *db.PrismaClient) {
	if err := client.Prisma.Disconnect(); err != nil {
		log.Fatal().Err(err).Msg("Error disconnecting from Prisma Client")
	}
}
