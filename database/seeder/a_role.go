package seeder

import (
	"time"

	"github.com/google/uuid"
)

func (s Seed) A_Role() {
	var arr = []string{
		"Superadmin",
		"Verified",
		"Inactive",
	}
	for i := 0; i < len(arr); i++ {
		_, err := s.db.Exec(`INSERT INTO roles(uuid, name, created_at) VALUES ($1,$2,$3)`,
			uuid.New(),
			arr[i],
			time.Now(),
		)
		if err != nil {
			panic(err)
		}
	}
}
