package entity

type Comment struct {
	ID        string
	UserId    string
	ArticleId string
	Body      string
	CreatedAt int64   `gorm:"autoCreateTime:milli"`
	User      User    `gorm:"foreignKey:user_id;references:id"`
	Article   Article `gorm:"foreignKey:article_id;references:id"`
}
