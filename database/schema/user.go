package schema

type User struct {
	Id       int    `gorm:"primarykey"`
	Name     string `gorm:"column:name;not null"`
	Email    string `gorm:"column:email;not null; unique"`
	Password string `gorm:"column:password; not null"`
}

func (User) TableName() string {
	return "users"
}

func (User) Pk() string {
	return "id"
}

func (u User) Ref() string {
	return u.TableName() + "(" + u.Pk() + ")"
}

func (u User) AddForeignKeys() {
}

func (u User) InsertDefaults() {
}
