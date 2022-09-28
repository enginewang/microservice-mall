package repository

// 连接业务Service和底层数据库的UserRepository接口
import (
	"github.com/enginewang/microservice-mall/user/domain/model"
	"github.com/jinzhu/gorm"
)

type IUserRepository interface {
	InitTable() error
	FindUserByID(int64) (*model.User, error)
	FindUserByName(string) (*model.User, error)
	CreateUser(user *model.User) (int64, error)
	DeleteUserByID(int64) error
	UpdateUser(user *model.User) error
	FindAll() ([]model.User, error)
}

type UserRepository struct {
	mysqlDB *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{mysqlDB: db}
}

// 实现接口
func (u *UserRepository) InitTable() error {
	return u.mysqlDB.CreateTable(&model.User{}).Error
}

func (u *UserRepository) FindUserByName(name string) (user *model.User, err error) {
	err = u.mysqlDB.Where("user_name = ?", name).Find(user).Error
	return
}

func (u *UserRepository) FindUserByID(id int64) (user *model.User, err error) {
	err = u.mysqlDB.First(user, id).Error
	return
}

func (u *UserRepository) CreateUser(user *model.User) (int64, error) {
	return user.ID, u.mysqlDB.Create(user).Error
}

func (u *UserRepository) DeleteUserByID(id int64) error {
	return u.mysqlDB.Where("id = ?", id).Delete(&model.User{}).Error
}

func (u *UserRepository) UpdateUser(user *model.User) error {
	return u.mysqlDB.Model(user).Update(&user).Error
}

func (u *UserRepository) FindAll() (allUser []model.User, err error) {
	return allUser, u.mysqlDB.Find(&allUser).Error
}
