package controller

import (
	"database/sql"
	"fmt"
	"medium/dto"
	"medium/util"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type UserController interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUserById(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)
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

type LoginUser struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func (uc *UserControllerImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userData dto.SignUpUser

	if err := util.ReadJson(r, &userData); err != nil {
		fmt.Println("error reading data ", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), 10)

	query := "INSERT INTO Users (username , email , password) VALUES( ? , ? , ?) "

	res, err := uc.db.Exec(query, userData.Username, userData.Email, string(hashedPassword))

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

func (uc *UserControllerImpl) LoginUser(w http.ResponseWriter, r *http.Request) {
	var userData dto.LoginUserDto

	if err := util.ReadJson(r, &userData); err != nil {
		fmt.Println("error reading body", err)
	}

	query := "SELECT id , username , email , password FROM Users WHERE email=?"

	var user LoginUser

	res := uc.db.QueryRow(query, userData.Email).Scan(&user.Id, &user.Email, &user.Username, &user.Password)

	if res != nil {
		if res == sql.ErrNoRows {
			fmt.Println("no row found with this email", res)
		}
		fmt.Println("error querying db", res)
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userData.Password))

	if err != nil {
		fmt.Println("error verifying password", err)
	}

	payload := jwt.MapClaims{
		"userId": user.Id,
		"email":  user.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	err = godotenv.Load()

	if err != nil {
		fmt.Println("error getting data from env", err)
	}

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		fmt.Println("something went wrong", err)
	}

	cookie := http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)

	util.JsonResponse(w, http.StatusOK, map[string]any{
		"msg":      "logged in user",
		"id":       user.Id,
		"email":    user.Email,
		"username": user.Username,
	})
}
