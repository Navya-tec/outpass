package models

import (
	"errors"
	"fmt"
	"goprojects/outpass/db"
	"goprojects/outpass/utils"
)

type User struct {
	Id       int64
	Name     string `json:"name"`
	Hostel   string `json:"hostel"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) SaveUSer() error {

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		panic("Error in hashing")
	}

	query := `INSERT INTO users(name,hostel,username,password) VALUES (?,?,?,?)`
	res, err := db.DB.Exec(query, u.Name, u.Hostel, u.Username, hashedPassword)
	if err != nil {
		panic("Error in executing insert query for user")
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic("Error in getting last insert id")
	}

	u.Id = id
	return err
}

func (u User) ValidateCredentials() error {
	query := `SELECT password FROM users WHERE username=?`
	row := db.DB.QueryRow(query, u.Username)

	var retreivedPassword string
	err := row.Scan(&retreivedPassword)

	if err != nil {
		panic("Error in Validating!")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retreivedPassword)

	if !passwordIsValid {
		return errors.New("Credentials Invalid")
	}

	return err
}

func GetUsers() ([]User,error){

	query:=`SELECT id,name,hostel FROM users`
	rows,err:=db.DB.Query(query)
	if err!=nil{
		panic("Error in fetching all users from DB")
	}
 
	var users []User
	for rows.Next(){
		var user User
		err:=rows.Scan(&user.Id,&user.Name,&user.Hostel)
		if err!=nil{
			panic("Error Scanning rows")
		}
        users=append(users, user)
	}

	if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error iterating over user rows: %v", err)
    }

	return users,err
}
