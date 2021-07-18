package schema

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	Tittle      string `gorm:"column:tittle;not null"`
	UserId      int    `gorm:"column:user_id; not nuill"`
	Description string `gorm:"column:description;not null"`
}

func (Todo) TableName() string {
	return "todos"
}

func (Todo) Pk() string {
	return "id"
}

func (t Todo) Ref() string {
	return t.TableName() + "(" + t.Pk() + ")"
}

func (t Todo) AddForeignKeys() {
	Database.Model(&t).AddForeignKey("user_id", User{}.Ref(), "CASCADE", "RESTRICT")
}

func (t Todo) InsertDefaults() {

}
