package seeder

import (
	"strconv"

	"github.com/google/uuid"

	"time"

	hash "github.com/arif-x/sqlx-gofiber-boilerplate/pkg/hash"
)

func (s Seed) A_UserSeeder() {
	password, _ := hash.Hash([]byte("password"))
	for i := 0; i < 100; i++ {
		_, err := s.db.Exec(`INSERT INTO users(uuid, name, username, email, password, created_at) VALUES ($1,$2,$3,$4,$5,$6)`,
			uuid.New(),
			"Name "+strconv.Itoa(i+1),
			"username"+strconv.Itoa(i+1),
			"email"+strconv.Itoa(i+1)+"@gmail.com",
			password,
			time.Now(),
		)
		if err != nil {
			panic(err)
		}
	}
}
