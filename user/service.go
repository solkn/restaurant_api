package user

import "github.com/solkn/restaurant_api/entity"

// UserService specifies application user related services
type UserService interface {
	Users() ([]entity.User, []error)
	User(id uint) (*entity.User, []error)
	UserByEmail(email string)(*entity.User,[]error)
	UpdateUser(user *entity.User) (*entity.User, []error)
	DeleteUser(id uint) (*entity.User, []error)
	StoreUser(user *entity.User) (*entity.User, []error)
	EmailExists(email string)bool
	PhoneExists(phone string)bool
	UserRoles(*entity.User)([]entity.Role,[]error)
}

// RoleService speifies application user role related services
type RoleService interface {
	Roles() ([]entity.Role, []error)
	Role(id uint) (*entity.Role, []error)
	RoleByName(name string)(*entity.Role,[]error)
	UpdateRole(role *entity.Role) (*entity.Role, []error)
	DeleteRole(id uint) (*entity.Role, []error)
	StoreRole(role *entity.Role) (*entity.Role, []error)
}
