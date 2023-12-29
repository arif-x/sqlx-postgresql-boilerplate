package response

import (
	dashboard "github.com/arif-x/sqlx-gofiber-boilerplate/app/model/dashboard"
	public "github.com/arif-x/sqlx-gofiber-boilerplate/app/model/public"
)

type AuthResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type ErrorResponse struct {
	Status  bool   `json:"status" example:"false"`
	Message string `json:"message"`
}

type UserResponse struct {
	Status  bool           `json:"status"`
	Message string         `json:"message"`
	Data    dashboard.User `json:"data"`
}

type UsersResponse struct {
	Status  bool             `json:"status"`
	Message string           `json:"message"`
	Data    []dashboard.User `json:"data"`
	Limit   int              `json:"limit"`
	Page    int              `json:"page"`
	Total   int              `json:"total"`
}

type PostCategoryResponse struct {
	Status  bool                   `json:"status"`
	Message string                 `json:"message"`
	Data    dashboard.PostCategory `json:"data"`
}

type PostCategoriesResponse struct {
	Status  bool                     `json:"status"`
	Message string                   `json:"message"`
	Data    []dashboard.PostCategory `json:"data"`
	Limit   int                      `json:"limit"`
	Page    int                      `json:"page"`
	Total   int                      `json:"total"`
}

type PostResponse struct {
	Status  bool           `json:"status"`
	Message string         `json:"message"`
	Data    dashboard.Post `json:"data"`
}

type PostsResponse struct {
	Status  bool             `json:"status"`
	Message string           `json:"message"`
	Data    []dashboard.Post `json:"data"`
	Limit   int              `json:"limit"`
	Page    int              `json:"page"`
	Total   int              `json:"total"`
}

type PublicPostResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    public.Post `json:"data"`
}

type PublicPostsResponse struct {
	Status  bool          `json:"status"`
	Message string        `json:"message"`
	Data    []public.Post `json:"data"`
	Limit   int           `json:"limit"`
	Page    int           `json:"page"`
	Total   int           `json:"total"`
}

type PublicPostsByCategoryResponse struct {
	Status  bool                        `json:"status"`
	Message string                      `json:"message"`
	Data    public.PostCategoryWithPost `json:"data"`
	Limit   int                         `json:"limit"`
	Page    int                         `json:"page"`
	Total   int                         `json:"total"`
}

type PublicPostsByUserResponse struct {
	Status  bool                `json:"status"`
	Message string              `json:"message"`
	Data    public.UserWithPost `json:"data"`
	Limit   int                 `json:"limit"`
	Page    int                 `json:"page"`
	Total   int                 `json:"total"`
}