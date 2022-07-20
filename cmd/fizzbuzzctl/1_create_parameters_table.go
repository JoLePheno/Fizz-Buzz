package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	fmt.Println("init 1_create_parameters_table")
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table parameters...")
		_, err := db.Exec(`CREATE TABLE IF NOT EXISTS parameters (
			id                  SERIAL PRIMARY KEY,

			created_at          TIMESTAMP NOT NULL,
			deleted_at			TIMESTAMP,
			
			str1				VARCHAR,
			str2				VARCHAR,
			int1				integer,
			int2				integer,
			limit_number		integer
		);`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table parameters...")
		_, err := db.Exec(`DROP TABLE parameters`)
		return err
	})
}
