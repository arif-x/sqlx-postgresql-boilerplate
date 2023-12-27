# SQLX Go-Fiber Boilerplate
Todo List
- [x]  Code Structure
- [x]  Fix CRUD Sample
- [x]  Dynamic List (Get paginate, search, sort by, and sort)
- [x]  One-to-one & one-to-many/many-to-many relational with native Postgre Query
    - [x]  (One-to-one) Example is in app/repository/public/post_r/PostByUser
    - [x]  (One-to-many/Many-to-many) Example is in app/repository/public/post_r/PostSingle
- [x]  Auth
- [ ]  Role & Permission

# Why no ORM?
- Slow
- Memory leaks vurnerable
- Can be make server hiccup

# How to Run DB Migration
- To create migration file, just run `migrate create -ext sql -dir database/migration -seq your_migration_file`
- Run if doing up `migrate -path database/migration/ -database "postgresql://postgres:password@localhost:5432/db_name?sslmode=disable" -verbose up`
- Run if doing down/rollback `migrate -path database/migration/ -database "postgresql://postgres:password@localhost:5432/db_name?sslmode=disable" -verbose down`
- If you want to migrate with specified file, just run  `migrate -path database/migration/your_migration_file.sql -database "postgresql://postgres:password@localhost:5432/db_name?sslmode=disable" -verbose up`

# How to Run DB Seeder
- Run `go run database/seeder/main/main.go`
