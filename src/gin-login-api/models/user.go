package auth_models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// user represents the user model

type User struct {
	gorm.Model        // ID, CreatedAt, UpdatedAt, DeletedAt
	Username   string `gorm:"uniuqe;not null" json:"username" `
	Email      string `gorm:"unique;not null" json:"email" binding:"required"`
	Password   string `gorm:"not null" json:"password" binding:"required"`
}

// HashPassword hashes the user's password before saving.
// Flow : user password -> hashed using salt -> store in db.
func (u *User) HashPassword() error {

	// on call of GenerateFromPassword bcrypt will generate new salt which it will use for hash.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}
