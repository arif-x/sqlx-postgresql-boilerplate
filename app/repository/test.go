package repository

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/arif-x/sqlx-gofiber-boilerplate/pkg/database"
	"github.com/google/uuid"
)

// JSONRaw ...
type JSONRaw json.RawMessage

// Value ...
func (j JSONRaw) Value() (driver.Value, error) {
	byteArr := []byte(j)

	return driver.Value(byteArr), nil
}

// Scan ...
func (j *JSONRaw) Scan(src interface{}) error {
	asBytes, ok := src.([]byte)
	if !ok {
		return error(errors.New("Scan source was not []bytes"))
	}
	err := json.Unmarshal(asBytes, &j)
	if err != nil {
		return error(errors.New("Scan could not unmarshal to []string"))
	}

	return nil
}

// MarshalJSON ...
func (j *JSONRaw) MarshalJSON() ([]byte, error) {
	return *j, nil
}

// UnmarshalJSON ...
func (j *JSONRaw) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
	}
	*j = append((*j)[0:0], data...)
	return nil
}

type UserWithPostsJSON struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Posts JSONRaw   `json:"posts"`
}

type Post struct {
	ID      uuid.UUID `json:"id"`
	UserID  int       `json:"user_id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
}

type TestRepository interface {
	Index() ([]UserWithPostsJSON, error)
}

type TestRepo struct {
	db *database.DB
}

func (repo *TestRepo) Index() ([]UserWithPostsJSON, error) {
	users := []UserWithPostsJSON{}
	err := repo.db.Select(&users, `
	SELECT
    u.id,
    u.name,
    u.email,
    json_agg(json_build_object(
            'post_id', p.id,
            'title', p.title,
            'content', p.content,
            'category', json_build_object(
                'category_id', c.id,
                'category_name', c.name
            )
        )) as posts
FROM
    users u
left join posts p on p.user_id = u.id
left join post_categories c on c.id = p.post_category_id 
group by u.id 



	`)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func NewTestRepo(db *database.DB) TestRepository {
	return &TestRepo{db}
}
