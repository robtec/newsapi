package api

// Response data returned from newsapi.org
type Response struct {
	News
}

// News news data returned from newsapi.org
type News struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

// Source source of news data returned from newsapi.org
type Source struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Article data retuned from newsapi.org
type Article struct {
	Source      Source
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	URLToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
}
