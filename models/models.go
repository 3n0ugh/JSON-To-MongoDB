package models

type Books struct {
	Books []Book `json:"books"`
}

type Book struct {
	Title            string            `json:"title"`
	ISBN             string            `json:"isbn" gorm:"uniqueIndex" bson:"_id"`
	PageCount        int               `json:"pageCount"`
	PublishDate      map[string]string `json:"publishDate"`
	ThumbnailURL     string            `json:"thumbnailUrl"`
	ShortDescription string            `json:"shortDescription"`
	LongDescription  string            `json:"longDescription"`
	Status           string            `json:"status"`
	Authors          []string          `json:"authors"`
	Category         []string          `json:"categories"`
}
