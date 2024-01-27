package seeder

import (
	"time"

	"github.com/google/uuid"
)

func (s Seed) E_permission() {
	var arr = []string{
		"role-index", "role-show", "role-store", "role-update", "role-destroy",
		"permission-index", "permission-show", "permission-store", "permission-update", "permission-destroy",
		"user-index", "user-show", "user-store", "user-update", "user-destroy",
		"tags-index", "tags-show", "tags-store", "tags-update", "tags-destroy",
		"post-index", "post-show", "post-store", "post-update", "post-destroy",
		"sync-permission-index", "sync-permission-update",
	}
	for i := 0; i < len(arr); i++ {
		_, err := s.db.Exec(`INSERT INTO permissions(uuid, name, created_at) VALUES ($1,$2,$3)`,
			uuid.New(),
			arr[i],
			time.Now(),
		)
		if err != nil {
			panic(err)
		}
	}
}
