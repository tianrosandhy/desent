package entity

type Book struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title" gorm:"not null"`
	Author string `json:"author" gorm:"not null"`
	Year   int    `json:"year" gorm:"not null"`
}
