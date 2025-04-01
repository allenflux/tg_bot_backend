package managment

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"tg_bot_backend/internal/dao"
	"tg_bot_backend/internal/middleware"
	"tg_bot_backend/internal/model/entity"
	"time"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/golang-jwt/jwt/v5"
	"tg_bot_backend/api/managment/v1"
)

// JWTClaims represents the custom claims for the JWT token
type JWTClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (c *ControllerV1) LoginUser(ctx context.Context, req *v1.LoginUserReq) (res *v1.LoginUserRes, err error) {
	dbQuery := dao.Users.Ctx(ctx).Where("name = ?", req.UserName).Where("password = ?", req.PassWord)
	var users []entity.Users
	var totalCount int
	if err = dbQuery.ScanAndCount(&users, &totalCount, false); err != nil {
		g.Log().Errorf(ctx, "Failed to query and count users List: %v", err)
		return nil, fmt.Errorf("failed to fetch users list: %w", err)
	}
	if totalCount == 0 {
		g.Log().Errorf(ctx, "Invalid login credentials: %v", err)
		return nil, fmt.Errorf("invalid login credentials: %w", err)
	}

	// Create claims with user information
	claims := &JWTClaims{
		Username: req.UserName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2400 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(middleware.JwtSecretKey))
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, "Failed to generate token")
	}

	return &v1.LoginUserRes{
		Token:       signedToken,
		UserName:    req.UserName,
		Avatar:      "https://avatars.githubusercontent.com/u/44761321",
		Roles:       []string{"admin"},
		Permissions: []string{"*:*:*"},
	}, nil
}
