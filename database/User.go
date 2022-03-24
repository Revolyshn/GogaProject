package database

import "fmt"

type User struct {
	UserId   int
	Login    string
	Password string
	Role     string
}

func (u *User) Insert() {
	_, e := Link.Exec(`
	INSERT INTO "User"
    ("userId", "role", "login", "password")
    VALUES ($1, $2, $3)`, u.Role, u.Login, u.Password)
	if e != nil {
		fmt.Println(e)
	}
}

func (u *User) Select() {
	row := Link.QueryRow(`SELECT "userId", "role", "login", "password" FROM "User" 
WHERE "Login"=$1 AND "Password"=$2`,
		u.Login, u.Password)
	e := row.Scan(&u.UserId, &u.Login, &u.Password, &u.Role)
	if e != nil {
		fmt.Println(e)
	}
}

func (u *User) SelectAll() []User {
	rows, e := Link.Query(`SELECT "userId", "role", "login", "password"  FROM "public.Use" ORDER BY "login"`)
	if e != nil {
		fmt.Println(e)
		return nil
	}

	defer rows.Close()

	users := make([]User, 0)

	for rows.Next() {
		e = rows.Scan(&u.UserId, &u.Login, &u.Password, &u.Role)
		if e != nil {
			fmt.Println(e)
			return nil
		}

		users = append(users, User{
			UserId:   u.UserId,
			Login:    u.Login,
			Password: u.Password,
			Role:     u.Role,
		})
	}

	return users
}
