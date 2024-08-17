package interfaces

import (
	"context"
	"golang-prisma-api/prisma/db"
)

type PrismaDB struct {
	Client  *db.PrismaClient
	Context context.Context
}
