package controller

import (
	"database/sql"
	"fmt"
	"medium/dto"
	"medium/util"
	"net/http"
	"strconv"
)

type UserController interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUserById(w http.ResponseWriter, r *http.Request)
}

type UserControllerImpl struct {
	db *sql.DB
}

func NewController(_database *sql.DB) UserController {
	return &UserControllerImpl{
		db: _database,
	}
}

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (uc *UserControllerImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userData dto.SignUpUser

	if err := util.ReadJson(r, &userData); err != nil {
		fmt.Println("error reading data ", err)
	}

	//todo -- hash the password  before storing

	query := "INSERT INTO Users (username , email , password) VALUES( ? , ? , ?) "

	res, err := uc.db.Exec(query, userData.Username, userData.Email, userData.Password)

	if err != nil {
		fmt.Println("error executing query", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		fmt.Println("some error", err)
	}

	if rowsAffected > 0 {
		util.JsonResponse(w, http.StatusCreated, map[string]any{
			"message": "user created successfully",
			"data":    userData,
		})
	} else {
		http.Error(w, "failed to create user", http.StatusInternalServerError)
	}

}

func (uc *UserControllerImpl) GetUserById(w http.ResponseWriter, r *http.Request) {
	params := r.PathValue("id")

	id, err := strconv.Atoi(params)

	if err != nil {
		fmt.Println("eror converting params", err)
	}

	query := "SELECT id username email FROM Users WHERE id = ? "

	var user User

	err = uc.db.QueryRow(query, id).Scan(&user.Id, &user.Email, &user.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no rows found with this query", err)
		}

		fmt.Println("error query row ", err)
	}

	util.JsonResponse(w, http.StatusOK, map[string]any{
		"message": "got user details",
		"data":    user,
	})

}
