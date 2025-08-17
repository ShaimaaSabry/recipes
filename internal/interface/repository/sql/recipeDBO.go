package sql

type recipeDBO struct {
	Id   int `gorm:"primaryKey"`
	Name string
}

func (recipeDBO) TableName() string { return "recipes" }
