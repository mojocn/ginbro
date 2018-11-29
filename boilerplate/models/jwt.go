package models

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

func jwtGenerateToken(m *AuthorizationModel) (*jwtObj, error) {
	m.Password = ""
	expireAfterTime := time.Hour * time.Duration(viper.GetInt("app.jwt_expire_hour"))
	iss := viper.GetString("app.name")
	appSecret := viper.GetString("app.secret")
	expireTime := time.Now().Add(expireAfterTime)
	stdClaims := jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		IssuedAt:  time.Now().Unix(),
		Id:        fmt.Sprintf("%d", m.Id),
		Issuer:    iss,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, stdClaims)
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(appSecret))
	if err != nil {
		logrus.WithError(err).Fatal("config is wrong, can not generate jwt")
	}
	data := &jwtObj{AuthorizationModel: *m, Token: tokenString, Expire: expireTime, ExpireTs: expireTime.Unix()}
	return data, err
}

type jwtObj struct {
	AuthorizationModel
	Token    string    `json:"token"`
	Expire   time.Time `json:"expire"`
	ExpireTs int64     `json:"expire_ts"`
}

//JwtParseUser parse a jwt token and return an authorized identity
func JwtParseUser(tokenString string) (*AuthorizationModel, error) {
	if tokenString == "" {
		return nil, errors.New("token is not found in Authorization Bearer")
	}
	claims := jwt.StandardClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secret := viper.GetString("app.secret")
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims.VerifyExpiresAt(time.Now().Unix(), true) == false {
		return nil, errors.New("token is expired")
	}
	appName := viper.GetString("app.name")
	if !claims.VerifyIssuer(appName, true) {
		return nil, errors.New("token's issuer is wrong,greetings Hacker")
	}
	key := fmt.Sprintf("login:%s", claims.Id)
	jwtObj, err := mem.GetJwtObj(key)
	if err != nil {
		return nil, err
	}
	return &jwtObj.AuthorizationModel, err
}

//get an authorized user form memory store
func (s *memoryStore) GetJwtObj(id string) (value *jwtObj, err error) {
	vv, err := s.Get(id, false)
	if err != nil {
		return nil, err
	}
	value, ok := vv.(*jwtObj)
	if ok {
		return value, nil
	}
	return nil, errors.New("mem:has value of this id, but is not type of *jwtObj")
}
