package models

type Item struct {
	ID    string
	Name  string
	Price float32
}

func (i Item) Valid() bool {
	return len(i.Name) > 0 && i.Price >= 0
}
