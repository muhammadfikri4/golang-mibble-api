package user

import "golang-prisma-api/prisma/db"

func UsersDTOMapper(users []db.UserModel) []UsersDTO {
	var usersDTO []UsersDTO
	for _, user := range users {
		usersDTO = append(usersDTO, UsersDTO{
			Name:  user.Name,
			Email: user.Email,
		})
	}
	return usersDTO
}
