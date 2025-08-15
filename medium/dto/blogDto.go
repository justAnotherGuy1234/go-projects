package dto

// todo add validation here
type CreateBlogDto struct {
	UserId      int    `json:"userId"`
	BlogTitle   string `json:"blogTitle"`
	BlogImage   string `json:"blogImage"`
	BlogContent string `json:"blogContent"`
}
