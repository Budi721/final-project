package domain

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id            int        `json:"id_user,omitempty"`
	Username      string     `json:"username,omitempty"`
	Password      string     `json:"password,omitempty"`
	LoginAs       uint       `json:"login_as,omitempty"`
	TopikDiminati string     `gorm:"-" json:"topik_diminati,omitempty"`
	AuthToken     string     `gorm:"-" json:"token,omitempty"`
	Article       Article    `gorm:"foreignKey:IdUser" json:"-"`
	Project       Project    `gorm:"foreignKey:Admin" json:"-"`
	Enrollment    Enrollment `gorm:"foreignKey:IdUser" json:"-"`
}

func (user *User) SetPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(hashedPassword)
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

type Admin User
