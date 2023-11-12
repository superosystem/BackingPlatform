package service

import (
	"errors"
	"fmt"

	"github.com/superosystem/BackingPlatform/backend/src/entity"
	"github.com/superosystem/BackingPlatform/backend/src/model"
	"github.com/superosystem/BackingPlatform/backend/src/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(request model.RegisterRequest) (entity.User, error)
	Login(request model.LoginRequest) (entity.User, error)
	IsEmailAvailable(request model.CheckEmailRequest) (bool, error)
	SaveAvatar(ID int, fileLocation string) (entity.User, error)
	GetUserByID(ID int) (entity.User, error)
	GetAllUsers() ([]entity.User, error)
	UpdateUser(request model.FormUpdateRequest) (entity.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository}
}

func (s *userService) RegisterUser(input model.RegisterRequest) (entity.User, error) {
	user := entity.User{}
	user.Name = input.Name
	user.Password = input.Password
	user.Email = input.Email
	user.Occupation = input.Occupation

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)
	user.Role = "user"

	newUser, err := s.userRepository.Save(user)
	if err != nil {
		fmt.Println(err)
		return newUser, err
	}

	return newUser, nil
}

func (s *userService) Login(input model.LoginRequest) (entity.User, error) {
	email := input.Email
	password := input.Password

	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *userService) IsEmailAvailable(input model.CheckEmailRequest) (bool, error) {
	email := input.Email

	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *userService) SaveAvatar(ID int, fileLocation string) (entity.User, error) {
	user, err := s.userRepository.FindByID(ID)
	if err != nil {
		return user, err
	}

	user.AvatarFileName = fileLocation

	updatedUser, err := s.userRepository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *userService) GetUserByID(ID int) (entity.User, error) {
	user, err := s.userRepository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found on with that ID")
	}

	return user, nil
}

func (s *userService) GetAllUsers() ([]entity.User, error) {
	users, err := s.userRepository.FindAll()
	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *userService) UpdateUser(input model.FormUpdateRequest) (entity.User, error) {
	user, err := s.userRepository.FindByID(input.ID)
	if err != nil {
		return user, err
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation

	updatedUser, err := s.userRepository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}
