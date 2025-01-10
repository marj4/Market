package backend

type Products struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Picture_URL string `json:"picture"`
	Price       string `json:"price"`
}

type User struct {
	ID       int    `json:"id"`
	Login    string `json:"Login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
