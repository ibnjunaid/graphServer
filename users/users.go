package users

import (
	"database/sql"
	database "graphServer/db/mydatabase"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// User is exported
type User struct {
	ID       string `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

//Save is ..
func (user User) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO users(Username,Password) VALUES($1,$2)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(user.Username, user.Password)
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.RowsAffected()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted")
	return id
}

// Create is
func (user *User) Create() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO Users(Username,Password) VALUES($1,$2)")
	if err != nil {
		log.Fatal(err)
	}
	hashedPassword, err := HashPassword(user.Password)
	_, err = stmt.Exec(user.Username, hashedPassword)
	if err != nil {
		log.Fatal(err)
	}
	return 1
}

// HashPassword is
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash is ..
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//GetUserIDByUsername is ..
func GetUserIDByUsername(username string) (int, error) {
	stmt, err := database.Db.Prepare("SELECT ID FROM Users WHERE Username=$1")
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRow(username)
	var ID int
	err = row.Scan(&ID)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return 0, err
	}
	return ID, nil
}

func (user *User) Authenticate() bool {
	stmt, err := database.Db.Prepare("SELECT Password FROM Users WHERE Username=$1")
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRow(user.Username)

	var hashedPassword string
	err = row.Scan(&hashedPassword)

	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Fatal(err)
		}
	}
	return CheckPasswordHash(user.Password, hashedPassword)
}

type WrongUsernameOrPasswordError struct{}

func (m *WrongUsernameOrPasswordError) Error() string {
	return "WrongUsernameOrPasswordError"
}
