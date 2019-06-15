package registrations

import (
	"os"
	"strconv"
	"time"

	"shadow/internal/platform/mongo"

	"github.com/dgrijalva/jwt-go.git"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"crypto/sha256"
	"encoding/base64"
	"shadow/cmd/cli/email-template/templates"
	"shadow/internal/config"
	"shadow/internal/platform/smtp"
)

type UserLogin struct {
	Username string `json:"username" binding:"required" valid:"required"`
	Password string `json:"password,omitempty" binding:"required" valid:"stringlength(5|9),required"`
}

type User struct {
	Name      string            `json:"fullname" binding:"required" valid:"required"`
	Username  string            `json:"username" binding:"required" valid:"required"`
	Password  string            `json:"password,omitempty" binding:"required" valid:"stringlength(5|9),required"`
	Email     string            `json:"email" binding:"required" valid:"email,required"`
	BirthDay  string            `json:"dob,omitempty"`
	Gender    string            `json:"gender,omitempty"`
	Avatar    string            `json:"avatar,omitempty"`
	Payload   map[string]string `json:"payload,omitempty"`
	Active    bool              `json:"-"`
	UpdatedAt time.Time         `json:"_"`
	Tokens    []string          `json:"-"`
}

func (h User) Logout(username string) error {
	col := mongo.Session.C("Users")

	col.Update(bson.M{"username": username}, bson.M{"$set": bson.M{"tokens": ""}})

	return nil
}

func (h UserLogin) Login() (string, error) {

	config := config.GetConfig()

	var user User
	col := mongo.Session.C("Users")

	sha := sha256.Sum256([]byte(h.Password))
	h.Password = base64.StdEncoding.EncodeToString(sha[:])

	// Ghi dữ liệu
	err := col.Find(
		bson.M{"$and": []bson.M{
			bson.M{"username": h.Username},
			bson.M{"password": h.Password},
		}},
	).One(&user)

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Username": h.Username,
		"Name":     user.Name,
		"Email":    user.Email,
		"exp":      time.Now().UTC().Add(time.Duration(config.GetInt("auth.exp")) * time.Second).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString([]byte(config.GetString("auth.secret")))

	user.Tokens = append(user.Tokens, tokenString)

	col.Update(
		bson.M{"username": h.Username},
		bson.M{"$set": bson.M{"tokens": user.Tokens, "updatedat": time.Now()}},
	)

	return tokenString, nil
}

func (h User) ResendVerificationEmail(email, username string) error {

	config := config.GetConfig()
	host, _ := os.Hostname()

	go smtp.SendEmail(email, "Confirmation Email!", templates.Confirmation{
		username,
		[]string{
			"Welcome to Golang! We're very excited to have you on board.",
		},
		"To get started with Golang, please click here:",
		templates.Button{
			"#22BC66",
			"Confirm your account",
			"http://" + host + ":" + strconv.Itoa(config.GetInt("server.port")) + "/v1/user/verify/" + GenerateVerificationToken(email, 3600),
		},
		[]string{
			"Need help, or have questions? Just reply to this email, we'd love to help.",
		},
	}.Init())
	return nil
}

func GenerateVerificationToken(email string, exp int) string {
	config := config.GetConfig()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Email": email,
		"exp":   time.Now().UTC().Add(time.Duration(exp) * time.Second).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString([]byte(config.GetString("auth.secret")))

	return base64.StdEncoding.EncodeToString([]byte(tokenString))
}

func (h User) Register() (string, error) {

	config := config.GetConfig()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Username": h.Username,
		"Name":     h.Name,
		"Email":    h.Email,
		"exp":      time.Now().UTC().Add(time.Duration(config.GetInt("auth.exp")) * time.Second).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(config.GetString("auth.secret")))
	if err != nil {
		return "", err
	}

	// Tạo collection
	col := mongo.Session.C("Users")

	sha := sha256.Sum256([]byte(h.Password))
	h.Password = base64.StdEncoding.EncodeToString(sha[:])
	h.Tokens = append(h.Tokens, tokenString)
	h.UpdatedAt = time.Now()

	// Ghi dữ liệu
	err = col.Insert(h)
	if err != nil {
		return "", err
	}

	err = h.ResendVerificationEmail(h.Email, h.Name)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (h User) GetByID(id string) (*User, error) {
	col := mongo.Session.C("Users")

	var user User
	err := col.Find(bson.M{"id": id}).One(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func CheckUserTokenMatch(username, email, token string) bool {
	col := mongo.Session.C("Users")
	var user User
	err := col.Find(
		bson.M{"$and": []bson.M{
			bson.M{"username": username},
			bson.M{"email": email},
			bson.M{"tokens": token},
		}},
	).One(&user)

	if err != mgo.ErrNotFound || err == nil {
		return true
	}

	return false
}

func (h User) CheckUserExist() bool {
	col := mongo.Session.C("Users")
	var user User
	err := col.Find(bson.M{"$or": []bson.M{bson.M{"username": h.Username}, bson.M{"email": h.Email}}}).One(&user)

	if err != mgo.ErrNotFound {
		return true
	}

	return false
}

func (h User) SetUserActive(email string) error {
	col := mongo.Session.C("Users")
	return col.Update(bson.M{"email": email}, bson.M{"$set": bson.M{"active": true}})
}

func (h User) GetUserInfo(username string) (*User, error) {
	col := mongo.Session.C("Users")

	var user User
	err := col.Find(bson.M{"username": username}).One(&user)
	if err == mgo.ErrNotFound {
		return nil, err
	}

	return &user, nil
}
