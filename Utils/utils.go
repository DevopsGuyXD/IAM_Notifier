package util

import (
	"log"
	"os"
	"os/exec"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/joho/godotenv"
)

func CheckForNil(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func InitEnvFile(){
	err := godotenv.Load(".env"); CheckForNil(err)
}

func InitAws(){
	InitEnvFile()

	exec.Command("aws", "configure", "set", "aws_access_key_id", os.Getenv("ACCESS_KEY_ID")).Output()
	exec.Command("aws", "configure", "set", "aws_secret_access_key", os.Getenv("ACCESS_KEY_SECRET")).Output()
	exec.Command("aws", "configure", "set", "region", os.Getenv("AWS_REGION")).Output()

	_, err := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region: aws.String(os.Getenv("AWS_REGION")),
		},
	})
	CheckForNil(err)
}