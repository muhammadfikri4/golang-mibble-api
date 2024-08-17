package user

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"golang-prisma-api/src/interfaces"
	"golang-prisma-api/src/utils/meta"

	"github.com/gin-gonic/gin"
)

func PostUserHandler(app *gin.Context) {
	var user UserBodyRequest

	err := app.ShouldBindJSON(&user)

	if err != nil {
		log.Fatal(err)
	}

	if user.Email == "" || user.Name == "" {
		app.JSON(http.StatusBadRequest, interfaces.StatusBadRequest("Email and Name cannot be empty"))
		return
	}

	if user.Password == "" {
		app.JSON(http.StatusBadRequest, interfaces.StatusBadRequest("Password cannot be empty"))
		return
	}

	userExist := GetUserByEmail(user.Email)

	if userExist != nil {
		app.JSON(http.StatusBadRequest, interfaces.StatusBadRequest("User already exist"))
		return
	}

	newUser := UserBodyRequest{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	result, err := CreateUser(&newUser)

	if err != nil {
		app.JSON(http.StatusBadRequest, interfaces.StatusBadRequest(err.Error()))
		return
	}

	fmt.Println(result)

	fmt.Println(newUser, "us")
	fmt.Println(newUser.Email)
	res := interfaces.CREATED("User created successfully", result)

	app.JSON(http.StatusCreated, res)
}

func GetUsersHandler(app *gin.Context) {

	var (
		page    = app.Query("page")
		perPage = app.Query("perPage")
	)

	if perPage == "" {
		perPage = "10"
	}
	if page == "" {
		page = "1"
	}

	resPage, err := strconv.Atoi(page)
	if err != nil {
		fmt.Println("Error converting page:", err)
		return
	}

	resPerPage, err := strconv.Atoi(perPage)
	if err != nil {
		fmt.Println("Error converting perPage:", err)
		return
	}
	users := GetUsers(
		resPage,
		resPerPage,
	)
	totalData := GetUsersCount()

	result := UsersDTOMapper(users)
	meta := meta.Meta(
		resPage,
		resPerPage,
		totalData,
	)

	app.JSON(http.StatusOK, interfaces.OK("Get users successfully", result, &meta))
}
