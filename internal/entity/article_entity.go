package entity

type Article struct {
	ID          string
	Thumbnail   string
	Title       string
	Description string
	Body        string
	CreatedAt   int64     `gorm:"autoCreateTime:milli"`
	UpdatedAt   int64     `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	User        User      `gorm:"foreignKey:user_id;references:id"`
	Tags        []Tag     `gorm:"many2many:articles_tags;foreignKey:id;joinForeignKey:article_id;references:id;joinReferences:tag_id"`
	Comments    []Comment `gorm:"foreignKey:article_id;references:id"`
}
