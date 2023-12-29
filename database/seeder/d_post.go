package seeder

import (
	"context"
	"log"
	"math/rand"
	"strconv"

	model "github.com/arif-x/sqlx-gofiber-boilerplate/app/model/dashboard"
	"github.com/google/uuid"

	"time"
)

func (s Seed) D_PostSeeder() {
	user, err := s.db.QueryContext(context.Background(), `SELECT uuid, name, email, username, created_at, updated_at FROM users`)
	if err != nil {
		log.Fatal(err)
	}

	defer user.Close()
	var users []model.User
	for user.Next() {
		var i model.User
		if err := user.Scan(
			&i.UUID,
			&i.Name,
			&i.Email,
			&i.Username,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			log.Fatal(err)
		}
		users = append(users, i)
	}
	if err := user.Close(); err != nil {
		log.Fatal(err)
	}
	if err := user.Err(); err != nil {
		log.Fatal(err)
	}

	category, err := s.db.QueryContext(context.Background(), `SELECT uuid, name, created_at, updated_at FROM post_categories`)
	if err != nil {
		log.Fatal(err)
	}

	defer category.Close()
	var categories []model.PostCategory
	for category.Next() {
		var i model.PostCategory
		if err := category.Scan(
			&i.UUID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			log.Fatal(err)
		}
		categories = append(categories, i)
	}
	if err := category.Close(); err != nil {
		log.Fatal(err)
	}
	if err := category.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(users); i++ {
		for j := 0; j < 3; j++ {
			ShufflePostCategory(categories)
			_, err := s.db.Exec(`INSERT INTO posts(uuid, post_category_uuid, user_uuid, title, "content", created_at) VALUES ($1,$2,$3,$4,$5,$6)`,
				uuid.New(),
				categories[0].UUID,
				users[i].UUID,
				"Title "+strconv.Itoa(i+1)+"-"+strconv.Itoa(j+1),
				"Post Title "+strconv.Itoa(i+1)+"-"+strconv.Itoa(j+1)+" By "+users[i].Username,
				time.Now(),
			)
			if err != nil {
				panic(err)
			}
		}
	}
}

func ShufflePostCategory(r []model.PostCategory) {
	for i := len(r) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		r[i], r[j] = r[j], r[i]
	}
}
