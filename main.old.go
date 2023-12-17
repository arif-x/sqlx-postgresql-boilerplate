package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/jmoiron/sqlx"
// 	_ "github.com/lib/pq"
// )

// var db *sqlx.DB

// type UserWithPostsJSON struct {
// 	ID    int    `json:"id"`
// 	Name  string `json:"name"`
// 	Email string `json:"email"`
// 	Posts string `json:"posts"`
// }

// type Post struct {
// 	ID      int    `json:"id"`
// 	UserID  int    `json:"user_id"`
// 	Title   string `json:"title"`
// 	Content string `json:"content"`
// }

// func initDB() {
// 	var err error
// 	db, err = sqlx.Open("postgres", "user=postgres password=password dbname=go_test sslmode=disable")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	db.SetMaxOpenConns(25)
// 	db.SetMaxIdleConns(25)
// 	db.SetConnMaxLifetime(5 * time.Minute)
// }

// func getUsersWithPostsJSON() ([]UserWithPostsJSON, error) {
// 	users := []UserWithPostsJSON{}
// 	err := db.Select(&users, `
// 		SELECT users.id, name, email,
// 			   COALESCE(json_agg(row_to_json(posts.*)), '[]') as posts
// 		FROM users
// 		LEFT JOIN posts ON users.id = posts.user_id
// 		GROUP BY users.id, users.name, users.email
// 	`)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return users, nil
// }

// func main() {
// 	initDB()

// 	// Example: Get a list of users with their posts using row_to_json
// 	users, err := getUsersWithPostsJSON()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// fmt.Print(users)

// 	b, err := json.Marshal(users)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(string(b))

// 	// // Print users with posts
// 	// for _, user := range users {
// 	// 	fmt.Printf("User ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
// 	// 	fmt.Println("Posts:")

// 	// 	var posts []Post
// 	// 	err := json.Unmarshal([]byte(user.Posts), &posts)
// 	// 	if err != nil {
// 	// 		log.Fatal(err)
// 	// 	}

// 	// 	for _, post := range posts {
// 	// 		fmt.Printf("  Post ID: %d, UserID: %d, Title: %s, Content: %s\n", post.ID, post.UserID, post.Title, post.Content)
// 	// 	}
// 	// 	fmt.Println("------------------------------")
// 	// }
// }
