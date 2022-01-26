package repositories

import (
	"log"

	"github.com/aofdev/fiber-versioning-boilerplate/api/adapters"
	"github.com/aofdev/fiber-versioning-boilerplate/api/versions/v1/entities"
)

//IUserRepository interface allows us to access the CRUD Operations in mysql here.
type IUserRepository interface {
	GetAllUser(id string) (*entities.User, error)
}

type UserRepository struct {
	database *adapters.MysqlAdapter
}

func NewUserRepository(
	db *adapters.MysqlAdapter,
) *UserRepository {
	return &UserRepository{
		database: db,
	}
}

func (m *UserRepository) GetAllUser(id string) (*entities.User, error) {
	user := &entities.User{}
	db := m.database.Connect
	if err := db.First(user, "id = ?", id).Error; err != nil {
		log.Println("--- Error: Repo GetAllUser ---", err)
		return nil, err
	}
	return user, nil

}
