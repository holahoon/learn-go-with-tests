package main

type Dictionary map[string]string

// -- Type 1 --
// func Search(dictionary map[string]string, word string) string {
// 	return dictionary[word]
// }

// -- Type 2 --
// func (d Dictionary) Search(word string) string {
// 	return d[word]
// }

// -- Type 3 --
// var ErrNotFound = errors.New("could not find the word")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// -- Type 4 --
// func (d Dictionary) Add(word, definition string) {
// 	d[word] = definition
// }

// -- Type 5 --

const (
	ErrNotFound         = DictionaryErr("could not find the word")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = newDefinition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
