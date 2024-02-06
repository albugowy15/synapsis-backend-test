package seeder

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var users = []User{
	{
		Fullname:    "Angelo Schmitt",
		Address:     "Greenfelder Glen",
		PhoneNumber: "200-644-5434",
	},
	{
		Fullname:    "Natalie Kuhlman V",
		Address:     "Stark Courts",
		PhoneNumber: "984-993-0894",
	},
	{
		Fullname:    "Wm Williamson",
		Address:     "Hilton Highway",
		PhoneNumber: "383-534-6945",
	},
	{
		Fullname:    "Drew Breitenberg",
		Address:     "Trantow Corners",
		PhoneNumber: "637-850-3218",
	},
	{
		Fullname:    "Judith Donnelly",
		Address:     "Strosin Fords",
		PhoneNumber: "663-530-4726",
	},
	{
		Fullname:    "Traci Hickle",
		Address:     "Marcelle River",
		PhoneNumber: "658-339-9696",
	},
}

func createEmail(username string) string {
	return fmt.Sprintf("%s@gmail.com", username)
}

func createUsername(fullname string) string {
	username := strings.ReplaceAll(fullname, " ", "")
	username = strings.ToLower(username)
	return username
}

func hashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatalf("Error hashing password on seeder: %v", err)
	}
	return string(bytes)
}

type User struct {
	Fullname    string
	Address     string
	PhoneNumber string
}

func SeedUser(db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO users (username, email, fullname, password, address, phone_number) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	userPassword := hashPassword("password1for&All")
	for _, user := range users {
		username := createUsername(user.Fullname)
		email := createEmail(username)
		_, err := stmt.Exec(username, email, user.Fullname, userPassword, user.Address, user.PhoneNumber)
		if err != nil {
			return err
		}
	}

	log.Println("Seed user successfull")
	return nil
}
