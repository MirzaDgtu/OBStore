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

	err := r.store.db.Model(&u).Updates(map[string]interface{}{"firstname": u.FirstName,
		"lastname": u.LastName,
		"email":    u.Email,
		"inn":      u.Inn}).Error
	if err != nil {
		return u, err
	}

	u.Pass = ""
	return u, nil
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

	err = result.Update("loggedin", 1).Error
	if err != nil {
		return user, err
	}

	err = result.Find(&user).Error
	if err != nil {
		return model.User{}, err
	}

	user.Pass = ""
	return user, nil
}

func (r *UserRepository) SignOutUserById(id int) error {
	user := model.User{
		Model: gorm.Model{
			ID: uint(id),
		},
	}

	return r.store.db.Model(&user).Where("id = ?", id).Update("loggedin", 0).Error
}

func (r *UserRepository) ChangePassword(id int, pass string) error {
	var user model.User
	result := r.store.db.Table("users").Where("id = ?", id)
	err := result.First(&user).Error
	if err != nil {
		return err
	}

	err = hashPassword(&pass)
	if err != nil {
		return err
	}

	err = result.Update("pass", pass).Error
	if err != nil {
		return err
	}
	return nil
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

func (r *UserRepository) GetAll() (users []model.User, err error) {
	return users, r.store.db.Select("id, firstname, lastname, email, inn").Find(&users).Error
}
