package article

import (
	"goblog/app/models"
)

// Article 文章模型
type Article struct {
	models.BaseModel
	Title string
	Body  string
}

// Link 方法用来生成文章链接
func (a Article) Link() string {
	return ""
	//return route.Name2URL("articles.show", "id", strconv.FormatInt(a.ID, 10))
}
