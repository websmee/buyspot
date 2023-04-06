package http

import "strings"

type SimpleAuth struct {
	users       map[string]SimpleAuthUser
	idsByTokens map[string]string
}

type SimpleAuthUser struct {
	Username string
	Password string
	ID       string
	Token    string
}

func NewSimpleAuth(usersStr string) *SimpleAuth {
	users := make(map[string]SimpleAuthUser)
	idsByTokens := make(map[string]string)
	for _, userStr := range strings.Split(usersStr, ",") {
		userArr := strings.Split(userStr, " ")
		users[userArr[0]] = SimpleAuthUser{
			Username: userArr[0],
			Password: userArr[1],
			ID:       userArr[2],
			Token:    userArr[3],
		}
		idsByTokens[userArr[3]] = userArr[2]
	}

	return &SimpleAuth{
		users:       users,
		idsByTokens: idsByTokens,
	}
}

func (s *SimpleAuth) CheckCredentials(username, password string) bool {
	if _, ok := s.users[username]; !ok {
		return false
	}

	return s.users[username].Password == password
}

func (s *SimpleAuth) GetUserID(token string) string {
	return s.idsByTokens[token]
}

func (s *SimpleAuth) GetToken(username string) string {
	if _, ok := s.users[username]; !ok {
		return ""
	}

	return s.users[username].Token
}
