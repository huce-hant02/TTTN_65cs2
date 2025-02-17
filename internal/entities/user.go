package entities

type User struct {
	ID        int     `json:"id" db:"id"`
	UserName  string  `json:"username" db:"username"`
	PassWord  string  `json:"password" db:"password"`
	Status    string  `json:"status" db:"status"`
	CreatedAt string  `json:"created_at" db:"created_at"`
	UpdatedAt *string `json:"updated_at" db:"updated_at"`
	DeletedAt *string `json:"deleted_at" db:"deleted_at"`
}
