package store

import (
	"database/sql"
	"time"
)

type Post struct {
	ID         int64     `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Created_at time.Time `json:"created_at"`
}

type PostStore interface { //database
	CreatePost(*Post) (*Post, error)
	GetPosts() ([]Post, error)
	GetPostById(id int64) (*Post, error)
}

type PostgresPostStore struct {
	db *sql.DB
}

func NewPostgresPostStore(db *sql.DB) *PostgresPostStore {
	return &PostgresPostStore{db: db}
}

func (pg *PostgresPostStore) CreatePost(post *Post) (*Post, error) {
	tx, err := pg.db.Begin() //put data to database..
	if err != nil {
		return nil, err
	}

	query := `INSERT INTO posts (title, content) VALUES ($1,$2) RETURNING id`

	err = tx.QueryRow(query, post.Title, post.Content).Scan(&post.ID) //return createdId into database. 想要 RETURNING id，怎么要，用scan导出来。
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return post, nil

}

// ///new
func (pg *PostgresPostStore) GetPosts() ([]Post, error) {
	query := `SELECT id, title, content, created_at FROM posts ORDER BY created_at DESC`

	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// cuz it is coming in rows ,.. very confuse.. so rebuild the stracture..
	//so we create this var posts ..

	var posts []Post //why[] cuz using returning GetPost returning.
	for rows.Next() {
		var post Post

		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.Created_at,
		)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (pg *PostgresPostStore) GetPostById(id int64) (*Post, error) {
	post := &Post{}

	query := `SELECT id, title, content, creeated_at FROM posts WHERE id = $1`

	err := pg.db.QueryRow(query, id).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.Created_at,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return post, nil
}
