package service

import (
	"github.com/solkn/restaurant_api/entity"
	"github.com/solkn/restaurant_api/user"
)

// UserService implements menu.UserService interface
type UserService struct {
	userRepo user.UserRepository
}

// NewUserService  returns a new UserService object
func NewUserService(userRepository user.UserRepository)*UserService {
	return &UserService{userRepo: userRepository}
}

// Users returns all stored application users
func (us *UserService) Users() ([]entity.User, []error) {
	usrs, errs := us.userRepo.Users()
	if len(errs) > 0 {
		return nil, errs
	}
	return usrs, errs
}

// User retrieves an application user by its id
func (us *UserService) User(id uint) (*entity.User, []error) {
	usr, errs := us.userRepo.User(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}
func (us *UserService)UserByEmail(email string)(*entity.User,[]error){
	usr,errs:= us.userRepo.UserByEmail(email)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// UpdateUser updates  a given application user
func (us *UserService) UpdateUser(user *entity.User) (*entity.User, []error) {
	usr, errs := us.userRepo.UpdateUser(user)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// DeleteUser deletes a given application user
func (us *UserService) DeleteUser(id uint) (*entity.User, []error) {
	usr, errs := us.userRepo.DeleteUser(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// StoreUser stores a given application user
func (us *UserService) StoreUser(user *entity.User) (*entity.User, []error) {
	usr, errs := us.userRepo.StoreUser(user)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}
func (us *UserService)EmailExists(email string)bool{
	exists:= us.userRepo.EmailExists(email)

	return exists
}

func (us *UserService)PhoneExists(phone string)bool{
	exists:= us.userRepo.PhoneExists(phone)
	return exists
}

func (us *UserService)UserRoles(user *entity.User)([]entity.Role,[]error){
	userRoles,errs:= us.userRepo.UserRoles(user)
	if len(userRoles)>0{
		return nil,errs
	}
	return userRoles,errs
}