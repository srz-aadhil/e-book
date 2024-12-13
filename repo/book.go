package repo

type Book struct {
	ID int `gorm:"primarykey"`
	Title string `gorm:`
}