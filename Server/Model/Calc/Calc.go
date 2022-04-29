package Calc

import (
	"database/sql"
	"fmt"
	"github.com/mhthrh/Voiceflex/Utils/DbUtil/PgSql"
)

type Calc struct {
	Number1 float64 `json:"Number1" validate:"required"`
	Number2 float64 `json:"Number2" validate:"required"`
	Ticket  string  `json:"ticket" validate:"required"`
}

func New(a, b float64, t string) *Calc {
	return &Calc{
		Number1: a,
		Number2: b,
		Ticket:  t,
	}
}
func (c *Calc) Sum(d *sql.DB) (float64, error) {
	rows, err := PgSql.RunQuery(d, fmt.Sprintf("SELECT count(*) FROM voiceflex_schm.sessions  where session_id='%s' and status=true and valid_till >=  now()::timestamp", c.Ticket))
	var counter int
	if err != nil {
		return 0, err
	}
	for rows.Next() {
		rows.Scan(&counter)
	}
	if counter != 1 {
		return 0, fmt.Errorf("invalid ticket")
	}
	return c.Number1 + c.Number2, nil
}
