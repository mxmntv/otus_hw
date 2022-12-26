package hw10programoptimization

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	iter "github.com/json-iterator/go"
)

var json = iter.ConfigCompatibleWithStandardLibrary

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	u, err := getUsers(r)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return countDomains(u, domain)
}

type users []User

func getUsers(r io.Reader) (users, error) {
	user := User{}
	result := make(users, 0, 100_000)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if err := json.Unmarshal(scanner.Bytes(), &user); err != nil {
			return nil, err
		}
		result = append(result, user)
	}
	return result, nil
}

func countDomains(u users, domain string) (DomainStat, error) {
	result := make(DomainStat)

	for _, user := range u {
		if strings.Contains(user.Email, domain) {
			email := strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])
			result[email]++
		}
	}
	return result, nil
}
