package impl

import (
	"context"
	"strings"
	"time"

	"github.com/maxzycon/SIB-Golang-Final-Project-1/pkg/authutil"
	"github.com/maxzycon/SIB-Golang-Final-Project-1/pkg/constant/role"

	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/maxzycon/SIB-Golang-Final-Project-1/internal/domain/user/dto"
	"github.com/maxzycon/SIB-Golang-Final-Project-1/pkg/errors"
	"github.com/maxzycon/SIB-Golang-Final-Project-1/pkg/helper"
	"github.com/maxzycon/SIB-Golang-Final-Project-1/pkg/model"
	"github.com/maxzycon/SIB-Golang-Final-Project-1/pkg/util/pagination"
)

func (s *UserService) Login(ctx context.Context, payload dto.PayloadLogin) (resp *dto.LoginRes, err error) {
	user, err := s.UserRepository.FindUserByUsernameLogin(ctx, payload.Username)
	if err != nil {
		log.Errorf("[Login] findUserByusername :%+v", err)
		return
	}

	password := helper.CheckPasswordHash(payload.Password, user.Password)
	if !password {
		log.Errorf("[Login] err hash password doens't match")
		err = errors.ErrInvalidPassword
		return
	}

	// --- set 30 day exp
	exp := time.Now().Add((time.Hour * 24) * 30).Unix()
	claims := jwt.MapClaims{
		"id":       user.ID,
		"email":    user.Email,
		"username": user.Username,
		"role":     user.Role,
		"exp":      exp,
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	AccessToken, err := tokenClaims.SignedString([]byte(s.conf.JWT_SECRET_KEY))

	if err != nil {
		log.Errorf("[Login] err create access token %+v", err)
		return
	}

	resp = &dto.LoginRes{
		ID:          user.ID,
		Username:    user.Username,
		AccessToken: AccessToken,
		Role:        user.Role,
		Exp:         exp,
	}
	return
}

func (s *UserService) CreateUser(ctx context.Context, payload dto.PayloadCreateUser, claims *authutil.UserClaims) (resp *model.User, err error) {
	password, _ := helper.HashPassword(strings.Trim(payload.Password, " "))
	userPayload := model.User{
		Username: payload.Username,
		Email:    payload.Email,
		Password: password,
		Role:     payload.Role,
	}

	resp, err = s.UserRepository.Create(ctx, &userPayload)
	if err != nil {
		log.Errorf("[user.go][CreateUser] err create user :%+v", err)
		return
	}
	return
}

func (s *UserService) GetUserPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload, claims *authutil.UserClaims) (resp pagination.DefaultPagination, err error) {
	resp, err = s.UserRepository.FindAllUserPaginated(ctx, payload, claims)
	if err != nil {
		log.Errorf("[user.go][GetUserPaginated] err repository at service :%+v", err)
		return
	}

	respToDto := make([]*dto.UserRow, 0)
	list, ok := resp.Items.([]*model.User)
	if ok {
		for _, v := range list {
			respToDto = append(respToDto, &dto.UserRow{
				ID:       v.ID,
				Username: v.Username,
				Role:     v.Role,
				Email:    v.Email,
			})
		}
	}
	resp.Items = respToDto
	return
}

func (s *UserService) GetById(ctx context.Context, id int) (resp *dto.UserRowDetail, err error) {
	row, err := s.UserRepository.FindById(ctx, id)
	if err != nil {
		log.Errorf("[user.go][GetById] err repository at service :%+v", err)
		return
	}
	resp = &dto.UserRowDetail{
		ID:       row.ID,
		Username: row.Username,
		Email:    row.Email,
		Role:     row.Role,
	}
	return
}

func (s *UserService) UpdateUser(ctx context.Context, payload dto.PayloadUpdateUser, id int, claims *authutil.UserClaims) (resp *model.User, err error) {
	_, err = s.UserRepository.FindById(ctx, id)
	if err != nil {
		log.Errorf("[user.go][CreateUser] err create user :%+v", err)
		return
	}
	userPayload := model.User{
		Username: payload.Username,
		Email:    payload.Email,
		Role:     payload.Role,
	}

	if payload.Password != nil && *payload.Password != "" {
		password, _ := helper.HashPassword(strings.Trim(*payload.Password, " "))
		userPayload.Password = password
	}

	resp, err = s.UserRepository.Update(ctx, &userPayload, id)
	if err != nil {
		log.Errorf("[user.go][UpdateUser] err create user :%+v", err)
		return
	}
	return
}

func (s *UserService) GetUserByUsername(ctx context.Context, username string) (resp *model.User, err error) {
	resp, err = s.UserRepository.FindUserByUsername(ctx, username)
	if err != nil {
		log.Errorf("[user.go][FindByUsername] err repository at service :%+v", err)
		return
	}
	return
}

func (s *UserService) DeleteUserById(ctx context.Context, id int, claims *authutil.UserClaims) (resp *int64, err error) {
	if claims.Role == role.ROLE_MANAGER {
		_, err = s.UserRepository.FindByIdAndDepartmentId(ctx, id, *claims.DepartmentId)
		if err != nil {
			log.Errorf("[user.go][FindByUsername] err repository at service :%+v", err)
			return
		}
	}

	resp, err = s.UserRepository.DeleteUserById(ctx, id)
	if err != nil {
		log.Errorf("[user.go][FindByUsername] err repository at service :%+v", err)
		return
	}
	return
}

func (s *UserService) UpdateUserProfile(ctx context.Context, id int, password string) (resp *int64, err error) {
	newPassword, _ := helper.HashPassword(strings.Trim(password, " "))
	resp, err = s.UserRepository.UpdatePasswordByUserId(ctx, id, &newPassword)
	if err != nil {
		log.Errorf("[user.go][FindByUsername] err repository at service :%+v", err)
		return
	}
	return
}
func (s *UserService) GetUserByIdToken(ctx context.Context, userId uint) (resp *model.User, err error) {
	resp, err = s.UserRepository.GetUserByIdToken(ctx, userId)
	if err != nil {
		log.Errorf("[user.go][GetById] err repository at service :%+v", err)
		return
	}
	return
}
