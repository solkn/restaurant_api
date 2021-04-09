package entity

import "time"

// Category represents Food Menu Category
type Category struct {
	ID          uint
	Name        string `gorm:"type:varchar(255);not null"`
	Description string
	Image       string `gorm:"type:varchar(255)"`
	Items       []Item `gorm:"many2many:item_categories;auto_preload;constraint:OnUpdate:cascade:OnDelete:cascade"`
}

// Role repesents application user roles
type Role struct {
	ID   uint
	Name string `gorm:"type:varchar(255)"`
}

// Item represents food menu items
type Item struct {
	ID          uint
	Name        string `gorm:"type:varchar(255);not null"`
	Price       float32
	Description string   `gorm:"type:varchar(255);not null"`
	Categories  []Category   `gorm:"many2many:item_categories;auto_preload;constraint:OnUpdate:cascade:OnDelete:cascade"`
	Image       string       `gorm:"type:varchar(255)"`
	IngredientID  uint        `json:"ingredient_id"`         
	Ingredients []Ingredient `gorm:"foreignKey:ingredient_id;auto_preload;constraint:OnUpdate:cascade:OnDelete:cascade"`
}

// Ingredient represents ingredients in a food item
type Ingredient struct {
	ID          uint
	Name        string `gorm:"type:varchar(255);not null"`
	Description string  `gorm:"type:varchar(255);not null"`
}

// Order represents customer order
type Order struct {
	ID       uint
	PlacedAt time.Time
	UserID   uint
	ItemID   uint
	Quantity uint
}

// User represents application user
type User struct {
	ID       uint
	UserName string `gorm:"type:varchar(255);not null"`
	FullName string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);not null; unique"`
	Phone    string `gorm:"type:varchar(100);not null; unique"`
	Password string `gorm:"type:varchar(255)"`
	RoleID    uint   `json:"role_id"`
	Roles    []Role `gorm:"foreignKey:role_id";auto_preload;constraint:OnUpdate:cascade:OnDelete:cascade`
	OrderID   uint    `json:"order_id"`
	Orders   []Order `gorm:"many2many:user_orders;auto_preload;constraint:OnUpdate:cascade:OnDelete:cascade"`
}

// Comment represents comments forwarded by application users
type Comment struct {
	ID       uint      `json:"id"`
	FullName string    `json:"fullname" gorm:"type:varchar(255)"`
	Message  string    `json:"message"`
	Phone    string    `json:"phone" gorm:"type:varchar(100);not null; unique"`
	Email    string    `json:"email" gorm:"type:varchar(255);not null; unique"`
	PlacedAt time.Time `json:"placedat"`
}

// Error represents error message
type Error struct {
	Code    int
	Message string
}
