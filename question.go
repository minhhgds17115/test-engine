package handlefunction

import (
	"database/sql"
	"errors"
	"\src\server\databaseGo\database.go"
)

func getQuestion(description string) (*Question, error) {
	return question(description, nil)
 }