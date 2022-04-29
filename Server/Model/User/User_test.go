package User

import (
	"database/sql"
	"github.com/pborman/uuid"
	"reflect"
	"testing"
)

func TestUser_SignIn(t *testing.T) {
	type fields struct {
		ID         uuid.UUID
		FirstName  string
		LastName   string
		UserName   string
		Password   string
		CellNo     string
		Email      string
		CreateDate string
		Ticket     string
	}
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *User
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &User{
				ID:         tt.fields.ID,
				FirstName:  tt.fields.FirstName,
				LastName:   tt.fields.LastName,
				UserName:   tt.fields.UserName,
				Password:   tt.fields.Password,
				CellNo:     tt.fields.CellNo,
				Email:      tt.fields.Email,
				CreateDate: tt.fields.CreateDate,
				Ticket:     tt.fields.Ticket,
			}
			got, err := c.SignIn(tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SignIn() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_SignOut(t *testing.T) {
	type fields struct {
		ID         uuid.UUID
		FirstName  string
		LastName   string
		UserName   string
		Password   string
		CellNo     string
		Email      string
		CreateDate string
		Ticket     string
	}
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &User{
				ID:         tt.fields.ID,
				FirstName:  tt.fields.FirstName,
				LastName:   tt.fields.LastName,
				UserName:   tt.fields.UserName,
				Password:   tt.fields.Password,
				CellNo:     tt.fields.CellNo,
				Email:      tt.fields.Email,
				CreateDate: tt.fields.CreateDate,
				Ticket:     tt.fields.Ticket,
			}
			if err := c.SignOut(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("SignOut() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_SignUp(t *testing.T) {
	type fields struct {
		ID         uuid.UUID
		FirstName  string
		LastName   string
		UserName   string
		Password   string
		CellNo     string
		Email      string
		CreateDate string
		Ticket     string
	}
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &User{
				ID:         tt.fields.ID,
				FirstName:  tt.fields.FirstName,
				LastName:   tt.fields.LastName,
				UserName:   tt.fields.UserName,
				Password:   tt.fields.Password,
				CellNo:     tt.fields.CellNo,
				Email:      tt.fields.Email,
				CreateDate: tt.fields.CreateDate,
				Ticket:     tt.fields.Ticket,
			}
			if err := c.SignUp(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("SignUp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
