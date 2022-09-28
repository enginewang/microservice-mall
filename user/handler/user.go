package handler

import (
	"context"
	"github.com/enginewang/microservice-mall/user/domain/model"
	"github.com/enginewang/microservice-mall/user/domain/service"
	user "github.com/enginewang/microservice-mall/user/proto/user"
)

type User struct {
	UserDataService service.IUserDataService
}

// Return a new handler
func New() *User {
	return &User{}
}

func (u *User) Register(ctx context.Context, userRegisterRequest *user.UserRegisterRequest, userRegisterResponse *user.UserRegisterResponse) error {
	userRegister := &model.User{
		UserName:     userRegisterRequest.UserName,
		FirstName:    userRegisterRequest.FirstName,
		HashPassword: userRegisterRequest.Pwd,
	}
	_, err := u.UserDataService.AddUser(userRegister)
	if err != nil {
		return err
	}
	userRegisterResponse.Message = "注册成功"
	return nil
}

func (u *User) Login(ctx context.Context, userLoginRequest *user.UserLoginRequest, userLoginResponse *user.UserLoginResponse) error {
	isOk, err := u.UserDataService.CheckPwd(userLoginRequest.UserName, userLoginRequest.Pwd)
	if err != nil {
		return err
	}
	userLoginResponse.IsSuccess = isOk
	return nil
}

func (u *User) GetUserInfo(ctx context.Context, userInfoRequest *user.UserInfoRequest, userInfoResponse *user.UserInfoResponse) error {
	userInfo, err := u.UserDataService.FindUserByName(userInfoRequest.UserName)
	if err != nil {
		return err
	}
	userInfoResponse = UserForResponse(userInfo)
	return nil
}

func UserForResponse(userModel *model.User) *user.UserInfoResponse {
	return &user.UserInfoResponse{
		UserName:  userModel.UserName,
		FirstName: userModel.FirstName,
		UserId:    userModel.ID,
	}
}
