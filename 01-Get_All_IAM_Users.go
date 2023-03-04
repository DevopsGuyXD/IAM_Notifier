package main

import (
	"encoding/json"
	"os/exec"
	"strings"

	util "github.com/DevopsGuyXD/IAM-Access-Key-Rotation/Utils"
)

//======================= STRUCTS =======================
type AllUsers struct{
	Users []Users `json:"Users"`
}

type Users struct{
	Path string `json:"Path"`
	UserName string `json:"UserName"`
	UserId string	`json:"UserId"`	
	Arn string	`json:"Arn"`
	CreateDate string	`json:"CreateDate"`
	PasswordLastUsed string	`json:"PasswordLastUsed"`
}

//======================= METHODS =======================
func (u AllUsers) GetUserCount() int{
	return len(u.Users)
}

func (u AllUsers) GetValedUsers(i int) string{

	var users string

	if strings.Contains(u.Users[i].UserName, "@"){
		users = u.Users[i].UserName
	}

	return users
}

//======================= FUNCTION =======================
func GetAllIAMUsers() {

	res, err := exec.Command("aws", "iam", "list-users").Output()
	util.CheckForNil(err)

	var allUsers AllUsers
	err = json.Unmarshal(res, &allUsers)
	util.CheckForNil(err)

	FilterUsers(allUsers)
}