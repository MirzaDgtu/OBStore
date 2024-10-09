package sqlstore

import (
	"errors"
	"obstore/internal/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u model.User) (model.User, error) {
	hashPassword(&u.Pass)
	err := r.store.db.Create(&u).Error
	u.Pass = ""
	return u, err
}

func (r *UserRepository) Update(u model.User) (model.User, error) {
	return u, r.store.db.Table("users").Save(&u).Error
}

func (r *UserRepository) SignInUser(email, password string) (user model.User, err error) {
	result := r.store.db.Table("users").Where(&model.User{Email: email})
	err = result.First(&user).Error
	if err != nil {
		return user, err
	}

	if !checkPassword(user.Pass, password) {
		return user, errors.New("Invalid password")
	}

	user.Pass = ""
	err = result.Update("loggedin", 1).Error
	if err != nil {
		return user, err
	}

	return user, result.Find(&user).Error
}

func (r *UserRepository) SignOutUserById(id int) error {
	user := model.User{
		Model: gorm.Model{
			ID: uint(id),
		},
	}
	return r.store.db.Table("Customers").Where(&user).Update("loggedin", 0).Error
}

func hashPassword(s *string) error {
	if s == nil {
		return errors.New("Reference provided for hashing password is nil")
	}
	//converd password string to byte slice
	sBytes := []byte(*s)
	//Obtain hashed password
	hashedBytes, err := bcrypt.GenerateFromPassword(sBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	//update password string with the hashed version
	*s = string(hashedBytes[:])
	return nil
}

func checkPassword(existingHash, incomingPass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(existingHash), []byte(incomingPass)) == nil
}
