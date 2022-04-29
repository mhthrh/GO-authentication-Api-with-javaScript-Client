package User

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/mhthrh/Voiceflex/Utils/CryptoUtil"
	"github.com/mhthrh/Voiceflex/Utils/DbUtil/PgSql"
	"github.com/pborman/uuid"
	"time"
)

type User struct {
	ID         uuid.UUID `json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	UserName   string    `json:"user_name"`
	Password   string    `json:"password"`
	CellNo     string    `json:"cell_no"`
	Email      string    `json:"email"`
	CreateDate string    `json:"create_date"`
	Ticket     string    `json:"ticket"`
}

func New() *User {
	return &User{
		ID:         uuid.NewUUID(),
		FirstName:  "",
		LastName:   "",
		UserName:   "",
		Password:   "",
		CellNo:     "",
		Email:      "",
		CreateDate: time.Now().Format("02-01-2006"),
	}
}

func (c *User) SignUp(db *sql.DB) error {
	var result int
	rows, err := PgSql.RunQuery(db, fmt.Sprintf("select count(*) from voiceflex_schm.users where user_name='%s'", c.UserName))
	if err != nil {
		return err
	}
	for rows.Next() {
		rows.Scan(&result)
	}
	if result != 0 {
		return errors.New("user already exist")
	}

	_, err = PgSql.ExecuteCommand(fmt.Sprintf("INSERT INTO voiceflex_schm.users(id, first_name, last_name, user_name, pass, cell_number, email, create_date) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')", c.ID, c.FirstName, c.LastName, c.UserName, crypto(c.Password), c.CellNo, c.Email, c.CreateDate), db)
	return err
}
func (c *User) SignOut(db *sql.DB) error {
	r, err := PgSql.ExecuteCommand(fmt.Sprintf("update voiceflex_schm.sessions set status=%t where session_id='%s' and status=true", false, c.Ticket), db)
	if err != nil {
		return err
	}
	count, err := (*r).RowsAffected()
	if err != nil {
		return err
	}
	if count != 1 {
		return errors.New("session not founded")
	}
	return nil
}
func (c *User) SignIn(db *sql.DB) (*User, error) {
	counter := 0

	rows, err := PgSql.RunQuery(db, fmt.Sprintf("SELECT id, first_name, last_name, user_name, cell_number, email FROM voiceflex_schm.users  where user_name='%s' and pass='%s'", c.UserName, crypto(c.Password)))

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		counter++
		rows.Scan(&c.ID, &c.FirstName, &c.LastName, &c.UserName, &c.CellNo, &c.Email)
	}
	if counter != 1 {
		return nil, fmt.Errorf("user not found")
	}
	ticket := uuid.NewUUID()
	_, err = PgSql.ExecuteCommand(fmt.Sprintf("INSERT INTO voiceflex_schm.sessions(id, user_id, session_id, valid_till,status) VALUES ('%s', '%s','%s',current_timestamp + (30 * interval '1 minute'),%t)", uuid.NewUUID(), c.ID, ticket, true), db)
	if err != nil {
		return nil, err
	}
	c.Ticket = ticket.String()
	c.Password = "********"
	return c, nil
}

func crypto(p string) string {
	k := CryptoUtil.NewKey()
	k.Text = p
	k.Sha256()
	return k.Result
}
