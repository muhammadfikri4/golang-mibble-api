package user

import (
	"golang-prisma-api/prisma/db"
	"golang-prisma-api/src/config"
)

func CreateUser(user *UserBodyRequest) (*db.UserModel, error) {
	prisma, ctx := config.Client()
	created, err := prisma.User.CreateOne(
		db.User.Name.Set(user.Name),
		db.User.Email.Set(user.Email),
		db.User.Password.Set(user.Password),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return created, nil
}

func GetUserByEmail(email string) *db.UserModel {
	prisma, ctx := config.Client()
	user, _ := prisma.User.FindFirst(
		db.User.Email.Contains(email),
	).Exec(ctx)

	return user
}

func GetUsers(page int, perPage int) []db.UserModel {
	prisma, ctx := config.Client()

	users, _ := prisma.User.
		FindMany().
		OrderBy(
			db.User.CreatedAt.Order(db.DESC),
		).
		Take(perPage).
		Skip((page - 1) * perPage).
		Exec(ctx)

	return users
}
func GetUsersCount() int {
	prisma, ctx := config.Client()

	users, _ := prisma.User.FindMany().Exec(ctx)

	return len(users)
}
