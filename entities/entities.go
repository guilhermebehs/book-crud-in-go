package entities

type Book struct {
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Year   string `json:"year"`
	Author string `json:"author"`
}

func (b *Book) AdjustAuthor(author string) {
	b.Author = author
}

func (b *Book) AdjustYear(year string) {
	b.Year = year
}

func (b *Book) AdjustTitle(title string) {
	b.Title = title
}

type HttpResponse struct {
	StatusCode int
	Data       any
}

type UpdateBookDto struct {
	Title  string
	Year   string
	Author string
}
