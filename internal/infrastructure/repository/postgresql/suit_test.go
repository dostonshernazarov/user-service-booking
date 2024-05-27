package postgresql

import (
	"Booking/user-service-booking/internal/entity"
	"Booking/user-service-booking/internal/pkg/config"
	"Booking/user-service-booking/internal/pkg/postgres"
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserPostgres(t *testing.T) {
	// Connect to database
	cfg := config.New()
	db, err := postgres.New(cfg)
	if err != nil {
		return
	}

	// Test  Method Create
	repo := NewUserRepo(db)
	user := &entity.User{
		Id:           uuid.New().String(),
		FullName:     "Testov Tester",
		Email:        "testovoy@gamil.com",
		Password:     "Qandaydir?",
		DateOfBirth:  "2013-01-02",
		ProfileImg:   "avatar.png",
		Card:         "123456789",
        Gender:       "Male",
        PhoneNumber:  "05555555555",
        Role:         "Test Role",
		RefreshToken: "test.refresh.token",
		CreatedAt:    time.Now(),
	}

	userMap := make(map[string]string)
	userMap["id"] = user.Id
	userMap["full_name"] = user.FullName
	userMap["email"] = user.Email
	userMap["password"] = user.Password
	userMap["date_of_birth"] = user.DateOfBirth
	userMap["profile_img"] = user.ProfileImg
	userMap["card"] = user.Card
	userMap["gender"] = user.Gender
	userMap["phone_number"] = user.PhoneNumber
	userMap["role"] = user.Role
	userMap["refresh_token"] = user.RefreshToken
	userMap["created_at"] = user.CreatedAt.Format("2006-01-02T15:04:05")

	ctx := context.Background()

	createdUser, err := repo.Create(ctx, user)
	assert.NoError(t, err)
	assert.Equal(t, user.Id, createdUser.Id)
	assert.Equal(t, user.FullName, createdUser.FullName)
	assert.Equal(t, user.Email, createdUser.Email)
	assert.Equal(t, user.Password, createdUser.Password)
	assert.Equal(t, user.DateOfBirth, createdUser.DateOfBirth)
	assert.Equal(t, user.ProfileImg, createdUser.ProfileImg)
	assert.Equal(t, user.Card, createdUser.Card)
	assert.Equal(t, user.Gender, createdUser.Gender)
	assert.Equal(t, user.PhoneNumber, createdUser.PhoneNumber)
	assert.Equal(t, user.Role, createdUser.Role)
	assert.Equal(t, user.RefreshToken, createdUser.RefreshToken)
	assert.Equal(t, user.CreatedAt, createdUser.CreatedAt)

	// Test Method Update
	user.FullName = "Test FullName"
	user.Email = "Test Email"
	user.Password = "Test Password"
	user.DateOfBirth = "2013-01-02"
	user.ProfileImg = "Test ProfileImg"
	user.Card = "Test Card"
	user.Gender = "Test Gender"
	user.PhoneNumber = "Test PhoneNumber"
	user.Role = "Test Role"
	user.RefreshToken = "test.refresh.token"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	updUser, err := repo.Update(ctx, user)
	assert.NoError(t, err)
	assert.Equal(t, user.Id, updUser.Id)
	assert.Equal(t, user.FullName, updUser.FullName)
	assert.Equal(t, user.Email, updUser.Email)
	assert.Equal(t, user.Password, updUser.Password)
	assert.Equal(t, user.DateOfBirth, updUser.DateOfBirth)
	assert.Equal(t, user.ProfileImg, updUser.ProfileImg)
	assert.Equal(t, user.Card, updUser.Card)
	assert.Equal(t, user.Gender, updUser.Gender)
	assert.Equal(t, user.PhoneNumber, updUser.PhoneNumber)
	assert.Equal(t, user.Role, updUser.Role)
	assert.Equal(t, user.RefreshToken, updUser.RefreshToken)
	assert.Equal(t, user.CreatedAt, updUser.CreatedAt)
	assert.Equal(t, user.UpdatedAt, updUser.UpdatedAt)

	//Test Method Get
	getUser, err := repo.Get(ctx, userMap)
	assert.NoError(t, err)
	assert.Equal(t, user.Id, getUser.Id)
	assert.Equal(t, user.FullName, getUser.FullName)
	assert.Equal(t, user.Email, getUser.Email)
	assert.Equal(t, user.DateOfBirth, getUser.DateOfBirth)
	assert.Equal(t, user.ProfileImg, getUser.ProfileImg)
	assert.Equal(t, user.Card, getUser.Card)
	assert.Equal(t, user.Gender, getUser.Gender)
	assert.Equal(t, user.PhoneNumber, getUser.PhoneNumber)
	assert.Equal(t, user.Role, getUser.Role)
	assert.Equal(t, user.RefreshToken, getUser.RefreshToken)

	// Test Method ListUsers
	listedUsers, count, err := repo.ListUsers(ctx, 0, 0, "id", "")
	assert.NoError(t, err)
	assert.NotEmpty(t, listedUsers)
	assert.NotNil(t, count)

	// Test Method SoftDelete User BY ID
	err = repo.SoftDelete(ctx, user.Id)
	assert.NoError(t, err)

	// Test Method Get All Users
	Users, count, err := repo.ListDeletedUsers(ctx, 0, 0, "id", "")
	assert.NoError(t, err)
	assert.NotEmpty(t, Users)
	assert.NotNil(t, count)
}