package repository_test

import (
	"database/sql"
	"log"
	"math/rand"
	"os"
	"regexp"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/superosystem/BackingPlatform/backend/src/entity"
	"github.com/superosystem/BackingPlatform/backend/src/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SuiteUser struct {
	suite.Suite
	DB         *gorm.DB
	mock       sqlmock.Sqlmock
	repository repository.UserRepository
}

func (s *SuiteUser) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,          // Disable color
		},
	)

	s.DB, err = gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: newLogger,
	})
	require.NoError(s.T(), err)

	s.repository = repository.NewUserRepository(s.DB)
}

func DataUser() entity.User {
	userMock := entity.User{
		ID:             rand.Int(),
		Name:           "User Test 1",
		Occupation:     "Tester",
		Email:          "user1@test.com",
		Password:       "user12345678",
		AvatarFileName: "filepath",
		Role:           "USER",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	return userMock
}

func (s *SuiteUser) Test_Create_User() {
	user1 := DataUser()
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "users" ("id", "name", "occupation", "email", "password", "avatar_file_name", "role", "created_at", "updated_at")
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETUNRING "users"."id"`)).
		WithArgs(user1.ID, user1.Name, user1.Occupation, user1.Email, user1.Password, user1.AvatarFileName, user1.Role, user1.CreatedAt, user1.UpdatedAt).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(user1.ID))

	_, err := s.repository.Save(user1)
	require.NoError(s.T(), err)
}

func (s *SuiteUser) Test_FindByID_User() {
	user1 := DataUser()
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE ("id" = $1)`)).
		WithArgs(user1.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "occupation", "email", "password_hash", "avatar_file_name", "role", "created_at", "updated_at"}).
			AddRow(user1.ID, user1.Name, user1.Occupation, user1.Email, user1.Password, user1.AvatarFileName, user1.Role, user1.CreatedAt, user1.UpdatedAt))
	res, err := s.repository.FindByID(user1.ID)

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(DataUser(), res))
}

func (s *SuiteUser) FindByUpdate_User() {
	// TODO
}

func (s *SuiteUser) FindALL_User() {
	// TODO
}
