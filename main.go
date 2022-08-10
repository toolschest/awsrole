package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	environ := flag.String("env", "default", "Environment name in config")
	mfa := flag.String("mfa", "000000", "MFA token for login")
	flag.Parse()
	var configPath string
	homeDir, err := os.UserHomeDir()
	if _, err := os.Stat(os.Getenv("AWS_ASSUMEROLE_CONFIG")); err == nil {
		configPath = os.Getenv("AWS_ASSUMEROLE_CONFIG")
	} else if _, err := os.Stat(homeDir + "/.aws/assumerole.yml"); err == nil {
		configPath = homeDir + "/.aws/assumerole.yml"
	} else {
		fmt.Println("Configuration file doesn't exist")
		os.Exit(1)
	}
	config, err := getConfig(configPath)
	if err != nil {
		fmt.Println("Config Error")
	}
	var arn *string
	var username *string
	var region *string
	if val, ok := config.Roles[*environ]; ok {
		arn = val.RoleARN
		username = val.Username
		region = val.Region
	} else {
		fmt.Println("Configuration doesn't exist with the environment name provided")
		os.Exit(1)
	}
	currentTime := (time.Now()).Unix()
	sessionName := *username + strconv.FormatInt(currentTime, 16)
	serialNumber := "arn:aws:iam::" + *config.MasterAccountID + ":mfa/" + *username
	assumedRole, err := assumeRole(arn, &sessionName, &serialNumber, mfa)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("export AWS_ACCESS_KEY_ID=" + *assumedRole.Credentials.AccessKeyId)
	fmt.Println("export AWS_SECRET_ACCESS_KEY=" + *assumedRole.Credentials.SecretAccessKey)
	fmt.Println("export AWS_SESSION_TOKEN=" + *assumedRole.Credentials.SessionToken)
	fmt.Println("export AWS_REGION=" + *region)
}
