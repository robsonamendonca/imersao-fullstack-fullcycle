package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type User struct {
	Base     `valid:"required"`
	Email     string     `json:"code" gorm:"type:varchar(20)" valid:"notnull"`
	Name     string     `json:"name" gorm:"type:varchar(255)" valid:"notnull"`
	Accounts []*Account `gorm:"ForeignKey:UserID" valid:"-"`
}

func (user *User) isValid() error {
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return err
	}
	return nil
}

func NewUser(email string, name string) (*User, error) {
	user := User{
		Email: email,
		Name: name,
	}
	user.ID = uuid.NewV4().String()
	user.CreatedAt = time.Now()
	err := user.isValid()
	if err != nil {
		return nil, err
	}
	return &user, nil
}