package dashboard

// type PostRepository interface {
// 	Index(limit int, offset uint, search string, sort_by string, sort string) ([]model.Post, int, error)
// 	Show(ID string) (model.Post, error)
// 	Store(model *model.StorePost) (model.Post, error)
// 	Update(ID string, Post *model.UpdatePost) (model.Post, error)
// 	Destroy(ID string) (model.Post, error)
// }

// type PostRepo struct {
// 	db *database.DB
// }

// func (repo *PostRepo) Index(limit int, offset uint, search string, sort_by string, sort string) ([]model.User, int, error) {
// 	_select := "id, title, content, users.id as users_id, username, created_at, updated_at"
// 	_conditions := database.Search([]string{"title", "content", "username"}, search)
// 	_order := database.OrderBy(sort_by, "posts.created_at")
// 	_limit := database.Limit(limit, offset)

// 	count_query := fmt.Sprintf(`SELECT count(*) FROM posts LEFT JOIN users ON users.id = posts.user_id %s`, _conditions)
// 	var count int
// 	_ = repo.db.QueryRow(count_query).Scan(&count)

// 	query := fmt.Sprintf(`SELECT %s FROM posts LEFT JOIN users ON users.id = posts.user_id %s %s %s`, _select, _conditions, _order, _limit)

// 	rows, err := repo.db.QueryContext(context.Background(), query)
// 	if err != nil {
// 		return nil, 0, err
// 	}

// 	defer rows.Close()
// 	var items []model.Post
// 	for rows.Next() {
// 		var i model.Post
// 		if err := rows.Scan(
// 			&i.ID,
// 			&i.Name,
// 			&i.Email,
// 			&i.Username,
// 			&i.CreatedAt,
// 			&i.UpdatedAt,
// 		); err != nil {
// 			return nil, 0, err
// 		}
// 		items = append(items, i)
// 	}
// 	if err := rows.Close(); err != nil {
// 		log.Fatal(err)
// 	}
// 	if err := rows.Err(); err != nil {
// 		log.Fatal(err)
// 	}

// 	return items, count, nil
// }

// func (repo *PostRepo) Show(ID string) (model.Post, error) {
// 	var user model.User
// 	query := "SELECT id, name, email, username, created_at, updated_at FROM users WHERE id = $1 LIMIT 1"
// 	err := repo.db.QueryRowContext(context.Background(), query, ID).Scan(
// 		&user.ID,
// 		&user.Name,
// 		&user.Email,
// 		&user.Username,
// 		&user.CreatedAt,
// 		&user.UpdatedAt,
// 	)
// 	if err != nil {
// 		return model.User{}, err
// 	}
// 	return user, err
// }

// func (repo *PostRepo) Store(request *model.StoreUser) (model.Post, error) {
// 	query := `INSERT INTO "users" (id, name, username, email, password, created_at) VALUES($1, $2, $3, $4, $5, $6)
// 	RETURNING id, name, username, email, created_at`
// 	var user model.User
// 	err := repo.db.QueryRowContext(context.Background(), query, uuid.New(), request.Name, request.Username, request.Email, request.Password, time.Now()).Scan(
// 		&user.ID,
// 		&user.Name,
// 		&user.Username,
// 		&user.Email,
// 		&user.CreatedAt,
// 	)
// 	if err != nil {
// 		return model.User{}, err
// 	}
// 	return user, err
// }

// func (repo *PostRepo) Update(ID string, request *model.UpdateUser) (model.Post, error) {
// 	query := `UPDATE "users" SET name = $2, username = $3, email = $4, password = $5, updated_at = $6 WHERE id = $1
// 	RETURNING id, name, username, email, created_at`
// 	var user model.User
// 	err := repo.db.QueryRowContext(context.Background(), query, ID, request.Name, request.Username, request.Email, request.Password, time.Now()).Scan(
// 		&user.ID,
// 		&user.Name,
// 		&user.Username,
// 		&user.Email,
// 		&user.CreatedAt,
// 	)
// 	if err != nil {
// 		return model.User{}, err
// 	}
// 	return user, err
// }

// func (repo *PostRepo) Destroy(ID string) (model.Post, error) {
// 	query := `UPDATE "users" SET updated_at = $2, deleted_at = $3 WHERE id = $1
// 	RETURNING id, name, username, email, created_at`
// 	var user model.User
// 	err := repo.db.QueryRowContext(context.Background(), query, ID, time.Now(), time.Now()).Scan(
// 		&user.ID,
// 		&user.Name,
// 		&user.Username,
// 		&user.Email,
// 		&user.CreatedAt,
// 	)
// 	if err != nil {
// 		return model.User{}, err
// 	}
// 	return user, err
// }
