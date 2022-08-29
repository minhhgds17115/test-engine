package handlefunction

import (
	"database/sql"
	"errors"
	"\src\server\databaseGo\database.go"
)

func creatUser(w http.ResponseWriter, r *http.Request) error {
	user := r.FormValue("user")
	_, err := json.NewDecoder(r.Body).Decode(user)
	user.id = strconv.Itoa(len(user)+1)
	user = append(user, user)

	json.NewDecoder(r.Body).Decode(user)
 }