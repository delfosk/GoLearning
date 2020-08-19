package main

//Dictionary type
type Dictionary map[string]string

// Search function
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

//ErrNotFound Constant

//Add a word in the Dicctionary function
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

//DictionaryErr type
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Constants
const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

//Update a word in the Dicctionary function
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist

	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

//Delete function
func (d Dictionary) Delete(word string) {
	delete(d, word)
}

var dictionary = make(map[string]string)

func main() {

}
