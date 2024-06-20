package models

type Role struct {
	ID   int
	Role string `gorm:"unique;not null"`
}

type User struct {
	ID        uint   `gorm:"primary_key"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Username  string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	RoleID    int
	Role      Role `gorm:"foreignKey:RoleID"`
}
