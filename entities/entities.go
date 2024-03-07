package entities

type Book struct {
	Isbn   string
	Title  string
	Year   string
	Author string
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
	Msg        any
}
