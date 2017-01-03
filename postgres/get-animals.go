package postgres

import (
	"database/sql"
	"strings"
)

func GetAnimals() (animals []string, err error) {
	var query []string

	query = nil
	query = append(query, "SELECT")
	query = append(query, "data")
	query = append(query, "FROM")
	query = append(query, "postgres_template.data_rows")

	rows, err := Database.Query(strings.Join(query, " "))
	if err == sql.ErrNoRows {
		return
	} else if err != nil {
		return
	} else {

		for rows.Next() {
			var animal string

			err = rows.Scan(&animal)

			if err != nil {
				panic(err)
			}

			// add to return variable
			animals = append(animals, animal)
		}
	}

	return
}

