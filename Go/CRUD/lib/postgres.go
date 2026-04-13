package lib

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type User struct {
	Id     int
	Name   string
	Emails []string
}

func (u *User) String() string {
	return fmt.Sprintf("User<%d %s %v>", u.Id, u.Name, u.Emails)
}

func ExampleDB_Model() {
	db := pg.Connect(&pg.Options{
		User: "postgres",
		Addr: "127.0.0.1:5432",
	})
	defer db.Close()

}

func CreateSchema(db *pg.DB) error {
	models := []interface{}{
		(*User)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
