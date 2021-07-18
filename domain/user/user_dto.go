package user

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDetail struct {
	Status  string     `json:"status"`
	Meta    MetaDetail `json:"meta"`
	Data    User       `json:"data"`
	Message string     `json:"message"`
}

type MetaDetail struct {
	TotalPage   int64 `json:"total_page"`
	CurrentPage int64 `json:"current_page"`
	PageSize    int64 `json:"page_size"`
	Total       int64 `json:"total"`
}
