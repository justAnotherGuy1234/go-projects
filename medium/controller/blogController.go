package controller

import (
	"database/sql"
	"fmt"
	"medium/dto"
	"medium/util"
	"net/http"
)

type BlogController interface {
	CreateBlog(w http.ResponseWriter, r *http.Request)
	UpdateBlog(w http.ResponseWriter, r *http.Request)
	DeleteBlog(w http.ResponseWriter, r *http.Request)
	GetBlogById(w http.ResponseWriter, r *http.Request)
	GetAllBlog(w http.ResponseWriter, r *http.Request)
}

type BlogControllerImpl struct {
	db *sql.DB
}

func NewBlogController(_database *sql.DB) BlogController {
	return &BlogControllerImpl{
		db: _database,
	}
}

func (bl *BlogControllerImpl) CreateBlog(w http.ResponseWriter, r *http.Request) {
	var blogData dto.CreateBlogDto

	if err := util.ReadJson(r, &blogData); err != nil {
		fmt.Println("error reading json body", err)
	}

	query := "INSERT INTO Blogs (userId , blogTitle , blogImage , blogContent) VALUES( ? , ? , ? , ?) "

	res, err := bl.db.Exec(query, blogData.UserId, blogData.BlogTitle, blogData.BlogImage, blogData.BlogContent)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		fmt.Println("error in query exece ", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		fmt.Println("some err", err)
	}

	if rowsAffected > 0 {
		util.JsonResponse(w, http.StatusCreated, map[string]any{
			"message": "created blog",
			"data":    blogData,
		})
	} else {
		http.Error(w, "failed to create blog", http.StatusInternalServerError)
	}

}

func (bl *BlogControllerImpl) UpdateBlog(w http.ResponseWriter, r *http.Request) {

}

func (bl *BlogControllerImpl) DeleteBlog(w http.ResponseWriter, r *http.Request) {

}

func (bl *BlogControllerImpl) GetBlogById(w http.ResponseWriter, r *http.Request) {

}

func (bl *BlogControllerImpl) GetAllBlog(w http.ResponseWriter, r *http.Request) {

}
