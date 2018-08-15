package properties

import (
	"os"
	"strconv"
)

type EmailCredentails struct {
	EmailAddress  string
	EmailPassword string
	SMTPServer    string
	SMTPPort      int
}

var emailCredentails *EmailCredentails

func getRequiredEnv(env string) string {
	result := os.Getenv(env)
	if len(result) == 0 {
		panic("NEED ENV " + env)
	}
	return result
}

func GetEmailCredentails() *EmailCredentails {
	if emailCredentails == nil {
		envSmtpPort := getRequiredEnv("SMTPPORT")
		smtpPort, err := strconv.ParseInt(envSmtpPort, 10, 32)
		if err != nil {
			panic(err)
		}
		emailCredentails = &EmailCredentails{
			EmailAddress:  getRequiredEnv("EMAIL"),
			EmailPassword: getRequiredEnv("EMAILPASSWORD"),
			SMTPServer:    getRequiredEnv("SMTPSERVER"),
			SMTPPort:      int(smtpPort),
		}
	}
	return emailCredentails

}
