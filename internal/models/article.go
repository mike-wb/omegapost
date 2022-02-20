// article.go

package models

import "fmt"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	{ID: 1, Title: "Article 1", Content: "Article 1 Body"},
	{ID: 2, Title: "Article 2", Content: "Article 2 Body"},
}

func GetAllArticles() []article {
	return articleList
}

func GetArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, fmt.Errorf("article not found: %d", id)
}
