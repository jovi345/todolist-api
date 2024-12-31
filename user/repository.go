package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	UpdateUser(email string, updatedData UserUpdatedData) (User, error)
	FindByRefreshToken(token string) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (r *repository) UpdateUser(email string, updatedData UserUpdatedData) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return User{}, err
	}

	err = r.db.Model(&user).Updates(updatedData).Error
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (r *repository) FindByRefreshToken(token string) (User, error) {
	var user User
	err := r.db.Where("refresh_token = ?", token).First(&user).Error
	if err != nil {
		return User{}, err
	}

	return user, nil
}
