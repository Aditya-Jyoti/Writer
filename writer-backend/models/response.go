package models

// @Description Represents a successful response when a new blog is created
type SuccessfulCreationResponse struct {
	Id   string `json:"id" example:"b1d5fca3-4b5b-46f3-9d13-cf6e2d2d4e12"`
	Code int    `json:"code"`
}
