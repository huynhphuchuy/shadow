package messages

import "github.com/gin-gonic/gin"

type Error struct {
	Code    int
	Message string
	Detail  interface{}
}

type Success struct {
	Code    int
	Message string
	Data    interface{}
}

var ErrorMsg = map[string]string{
	"ERR001": "Username or Email already existed!",
	"ERR002": "Error validating user info!",
	"ERR003": "An unexpected error has occured while registering, please try again!",
	"ERR004": "Login error, please try again!",
	"ERR005": "Bad request!",
	"ERR006": "Error retrieving user info!",
	"ERR007": "Email mismatch!",
	"ERR008": "Send verification email fail!",
	"ERR009": "Unable to set user active!",
	"ERR010": "Token is invalid or expired!",
	"ERR011": "Logout failed!",
	"ERR012": "Token is required!",
	"ERR013": "Authentication Credential not match!",
	"ERR014": "That's not even a token",
	"ERR015": "Token has been expired!",
	"ERR016": "An unexpected error has occured!",
}

var SuccessMsg = map[string]string{
	"SSS001": "Register successfully!",
	"SSS002": "Login successfully!",
	"SSS003": "User founded!",
	"SSS004": "Send verification email successfully, please check your mailbox!",
	"SSS005": "Verify email successfully!",
	"SSS006": "Logout successfully!",
}

func (err Error) Gen() (int, gin.H) {
	if err.Detail != nil {
		return err.Code, gin.H{"message": ErrorMsg[err.Message], "detail": err.Detail}
	}
	return err.Code, gin.H{"message": ErrorMsg[err.Message]}
}

func (sss Success) Gen() (int, gin.H) {
	if sss.Data != nil {
		return sss.Code, gin.H{"message": SuccessMsg[sss.Message], "data": sss.Data}
	}
	return sss.Code, gin.H{"message": SuccessMsg[sss.Message]}
}
