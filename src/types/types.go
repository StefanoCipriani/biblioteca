package types

type Book struct {
	Titolo, Autore string
}

func (b *Book) String() string {
	return "Titolo: " + b.Titolo + "\nAutore: " + b.Autore
}
