package models

type User struct {
	ID uint `gorm:"primaryKey" json:"id"`
	// CreatedAt time.Time  `json:"created_at"`
	// UpdatedAt time.Time  `json:"updated_at"`
	// DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u *User) TableName() string {
	return "users"
}
