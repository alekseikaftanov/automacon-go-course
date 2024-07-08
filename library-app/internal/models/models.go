package models

type Book struct {
	ID       int    `db:"id"` // тег db - это тег для sqlx, он указывает на то, что это поле в базе данных
	Title    string `db:"title"`
	AuthorID int    `db:"author_id"`
	Year     int    `db:"year"`
	ISBN     string `db:"isbn"`
}

type Author struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Surname  string `db:"surname"`
	Bio      string `db:"bio"`
	Birthday string `db:"birthday"`
}
