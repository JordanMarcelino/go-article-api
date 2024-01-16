package entity

type Tag struct {
	ID       int64
	Name     string
	Articles []Article `gorm:"many2many:articles_tags;foreignKey:id;joinForeignKey:tag_id;references:id;joinReferences:article_id"`
}
