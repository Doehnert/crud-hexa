package entity

type UserEntityMysql struct {
	ID       int64  `json:"id,omitempty"` // Assuming auto-incremented INT or BIGINT as the primary key
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name,omitempty"`
	Age      int8   `json:"age,omitempty"`
}
