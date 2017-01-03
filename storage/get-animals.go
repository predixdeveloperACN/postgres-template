package storage

import (
	pg "github.com/predixdeveloperACN/postgres-template/postgres"
)

func GetAnimals() (animals []string, err error) {

	// get animals from db
	animals, err = pg.GetAnimals()
	if err != nil {
		return
	}

	return
}

