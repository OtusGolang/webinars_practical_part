package user

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

type User struct {
	Id       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

func getComDomains(filename string) map[string]uint32 {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("%v", err)
	}

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("%v", err)
	}

	lines := strings.Split(string(fileContents), "\n")

	users := make([]User, 0)

	for _, line := range lines {
		user := &User{}
		json.Unmarshal([]byte(line), user)

		users = append(users, *user)
	}

	comDomains := make(map[string]uint32)

	for _, user := range users {
		matched, err := regexp.Match("\\.com", []byte(user.Email))
		if err != nil {
			log.Fatalf("%v", err)
		}

		if matched {
			num := comDomains[strings.SplitN(user.Email, "@", 2)[1]]
			num++
			comDomains[strings.SplitN(user.Email, "@", 2)[1]] = num
		}
	}

	return comDomains
}
