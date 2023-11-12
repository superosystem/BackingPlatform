package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/superosystem/BackingPlatform/backend/src/entity"
	"github.com/superosystem/BackingPlatform/backend/src/model"
	mocks "github.com/superosystem/BackingPlatform/backend/src/repository/mocks"
	service "github.com/superosystem/BackingPlatform/backend/src/service"
)

func TestRegisterUser(t *testing.T) {
	mockUser := entity.User{
		ID:         1,
		Name:       "johndoe",
		Email:      "johndoe@test.com",
		Occupation: "Tester",
		Password:   "pass12345678",
	}

	// di
	mockUserRepository := new(mocks.UserRepository)
	userService := service.NewUserService(mockUserRepository)

	t.Run("should success register user", func(t *testing.T) {
		tempMockUser := model.RegisterRequest{
			Name:       "johndoe",
			Email:      "johndoe@test.com",
			Occupation: "Tester",
			Password:   "pass12345678",
		}

		mockUserRepository.On("Save", mock.AnythingOfType("entity.User")).Return(mockUser, nil).Once()

		user, err := userService.RegisterUser(tempMockUser)
		assert.NoError(t, err)

		assert.Equal(t, mockUser.Name, user.Name)
	})

	t.Run("should fail register user cause empty name", func(t *testing.T) {
		tempMockUser := model.RegisterRequest{
			Name:       "johndoe",
			Email:      "johndoe@test.com",
			Occupation: "Tester",
			Password:   "pass12345678",
		}

		mockUserRepository.On("Save", mock.AnythingOfType("entity.User")).Return(mockUser, nil).Once()

		user, err := userService.RegisterUser(tempMockUser)
		assert.NoError(t, err)

		assert.Equal(t, mockUser.Name, user.Name)
	})

}
