package seeder

import (
	"strconv"
	"time"

	"github.com/google/uuid"
)

func (s Seed) B_PostCategory() {
	for i := 0; i < 3; i++ {
		_, err := s.db.Exec(`INSERT INTO post_categories(uuid, name, created_at) VALUES ($1,$2,$3)`,
			uuid.New(),
			"Category "+strconv.Itoa(i+1),
			time.Now(),
		)
		if err != nil {
			panic(err)
		}
	}
}
