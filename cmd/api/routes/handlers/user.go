package handlers

import (
	"encoding/base64"
	"shadow/internal/config"
	"shadow/internal/registrations"

	"github.com/asaskevich/govalidator"
	"github.com/dgrijalva/jwt-go.git"
	"github.com/gin-gonic/gin"

	"shadow/internal/helpers/messages"
)

type UserHandler struct{}
type E = messages.Error
type S = messages.Success

var userModel = new(registrations.User)
var userLogin = new(registrations.UserLogin)

func validateUserInfo(u *registrations.User) (bool, error) {
	valid, err := govalidator.ValidateStruct(u)
	return valid, err
}

func (u UserHandler) GetUserInfoById(c *gin.Context) {
	if c.Param("id") != "" {
		user, err := userModel.GetByID(c.Param("id"))
		if err != nil {
			c.JSON(E{500, "ERR006", err}.Gen())
			c.Abort()
			return
		}
		c.JSON(S{200, "SSS003", gin.H{"user": user}}.Gen())
		c.Abort()
		return
	}
	c.JSON(E{400, "ERR005", nil}.Gen())
	c.Abort()
	return
}

func (u UserHandler) Register(c *gin.Context) {

	c.BindJSON(userModel)
	if userModel.CheckUserExist() {
		c.JSON(E{500, "ERR001", nil}.Gen())
		c.Abort()
		return
	}

	validation, err := validateUserInfo(userModel)
	if validation == false {
		c.JSON(E{500, "ERR002", err.Error()}.Gen())
		c.Abort()
		return
	}

	token, err := userModel.Register()
	if err != nil {
		c.JSON(E{500, "ERR003", err}.Gen())
		c.Abort()
		return
	}

	c.JSON(S{200, "SSS001", gin.H{"token": token}}.Gen())
	c.Abort()
	return
}

func (u UserHandler) Login(c *gin.Context) {

	c.BindJSON(userLogin)

	token, err := userLogin.Login()
	if err != nil {
		c.JSON(E{500, "ERR004", err}.Gen())
		c.Abort()
		return
	}
	c.JSON(S{200, "SSS002", gin.H{"token": token}}.Gen())
	c.Abort()
	return
}

func (u UserHandler) Resend(c *gin.Context) {
	username, _ := c.Get("Username")
	email, _ := c.Get("Email")

	err := userModel.ResendVerificationEmail(email.(string), username.(string))
	if err != nil {
		c.JSON(E{500, "ERR008", err}.Gen())
		c.Abort()
		return
	}

	c.JSON(S{200, "SSS004", nil}.Gen())
	c.Abort()
	return
}

func (u UserHandler) Verify(c *gin.Context) {

	if c.Param("token") != "" {

		config := config.GetConfig()
		secret := config.GetString("auth.secret")

		decodedToken, err := base64.StdEncoding.DecodeString(c.Param("token"))

		token, err := jwt.Parse(string(decodedToken), func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err == nil && token.Valid {
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				err := userModel.SetUserActive(claims["Email"].(string))
				if err != nil {
					c.JSON(E{500, "ERR009", err}.Gen())
					c.Abort()
					return
				}
				c.JSON(S{200, "SSS005", nil}.Gen())
				c.Abort()
				return
			}
		} else {
			c.JSON(E{500, "ERR010", err}.Gen())
			c.Abort()
			return
		}

	}
	c.JSON(E{400, "ERR005", nil}.Gen())
	c.Abort()
	return
}

func (u UserHandler) Logout(c *gin.Context) {

	username, _ := c.Get("Username")

	err := userModel.Logout(username.(string))
	if err != nil {
		c.JSON(E{500, "ERR011", err}.Gen())
		c.Abort()
		return
	}
	c.JSON(S{200, "SSS006", nil}.Gen())
	c.Abort()
	return
}

func (u UserHandler) GetUserInfo(c *gin.Context) {
	username, _ := c.Get("Username")
	user, err := userModel.GetUserInfo(username.(string))
	if err != nil {
		c.JSON(E{500, "ERR002", err}.Gen())
		c.Abort()
		return
	}
	user.Password = ""
	c.JSON(S{200, "SSS003", gin.H{"user": user}}.Gen())
	c.Abort()
	return
}
