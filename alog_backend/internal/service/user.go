package service

import (
	"alog/configs"
	"alog/internal/model"
	"alog/internal/pkg/db"
	"alog/internal/pkg/log"
	"alog/internal/pkg/utils"
	"alog/internal/repo"
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(ctx context.Context, username string, password string) (*UserResponse, error)
	ChangeUserInfo(ctx context.Context, user model.User) (*UserResponse, error)
}

var userServiceImpl *userService

func InitUserService() {
	userServiceImpl = &userService{
		userRepo: repo.NewUserRepo(),
	}
}
func GetUserService() UserService {
	return userServiceImpl
}

type userService struct {
	userRepo *repo.UserRepo
}

func (u *userService) Login(ctx context.Context, username string, password string) (resp *UserResponse, err error) {
	dbUser := &model.User{
		Username: username,
	}
	if err = u.userRepo.Get(db.GetDB(), dbUser); err != nil {
		return nil, err
	}
	log.InfoContextf(ctx, "user info: %+v", dbUser.Username)
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("password error: %v", err)
	}
	token, err := utils.GenerateJwtToken(configs.GetGlobalConfig().JwtConfig.Secret,
		configs.GetGlobalConfig().JwtConfig.Issuer,
		configs.GetGlobalConfig().JwtConfig.Expire,
		dbUser.ID,
		dbUser.Username)
	if err != nil {
		return nil, err
	}
	response := &UserResponse{
		UserID:   dbUser.ID,
		Token:    token,
		Username: dbUser.Username,
	}
	return response, nil
}

func (u *userService) ChangeUserInfo(ctx context.Context, user model.User) (resp *UserResponse, err error) {
	if user.Password != "" {
		bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(bcryptPassword)
	}

	if err = u.userRepo.Update(db.GetDB(), &user); err != nil {
		return nil, err
	}
	response := &UserResponse{
		UserID:   user.ID,
		Token:    "token",
		Username: user.Username,
	}
	return response, nil
}
