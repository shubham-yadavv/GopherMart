package models

import "time"

type User struct {
	UserID       uint      `json:"user_id" gorm:"primaryKey"`
	Username     string    `json:"username" validate:"required,min=2,max=30"`
	Password     string    `json:"password" validate:"required,min=6"`
	Email        string    `json:"email" validate:"email,required"`
	UserRole     string    `json:"user_role" validate:"required"`
	RegisteredAt time.Time `json:"registered_at"`

	// Address_Details []Address     `json:"address" `
	// Order_Status    []Order       `json:"orders" `
	// UserCart        []ProductUser `json:"usercart"`
}
