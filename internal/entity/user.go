package entity

import "time"

type User struct {
	Id string
	FullName string
	Email string
	Password string
	DateOfBirth string
	ProfileImg string
	Card string
	Gender string
	PhoneNumber string
	Role string
	RefreshToken string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type UserRes struct {
	Id string
	FullName string
	Email string
	DateOfBirth string
	ProfileImg string
	Card string
	Gender string
	PhoneNumber string
	Role string
	RefreshToken string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}