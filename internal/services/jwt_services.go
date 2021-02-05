package services

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	jwt2 "github.com/smh2274/Felstorm/internal/api"
	"github.com/smh2274/Felstorm/internal/logger"
	"github.com/spf13/viper"
	"time"
)

type GRPCTokenServer struct {
	jwt2.TokenServer
	V *viper.Viper
}

func (s *GRPCTokenServer) GetToken(ctx context.Context, in *jwt2.GetTokenRequest) (*jwt2.GetTokenResponse, error) {
	claims := jwt.StandardClaims{
		Audience:  in.Audience,
		ExpiresAt: time.Now().Add(time.Duration(in.Exp)).Unix(),
		Issuer:    s.V.GetString("jwt.issuer"),
		IssuedAt:  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// SecretKey 用于对用户数据进行签名
	tokenStr, err := token.SignedString([]byte(s.V.GetString("jwt.key")))
	if err != nil {
		logger.RecordErr(in, err, "generate token")
		return nil, err
	}
	return &jwt2.GetTokenResponse{Token: tokenStr}, nil
}
