package maps

// ErrNotFound should be used if a key in the dict couldn't be found.
const ErrNotFound = DictErr("could not find the word you were looking for")

// ErrWordExists should be used if a key that is to be added already exists in the dict.
const ErrWordExists = DictErr("cannot add word because it already exists")

// ErrWordDoesNotExist should be used if a a key that should exist in the dict does not.
const ErrWordDoesNotExist = DictErr("word does not exist")

// DictErr is a type relating to dictionaries
type DictErr string

func (e DictErr) Error() string {
	return string(e)
}

// Dictionary represents a dictionary of words, accessible by index.
type Dictionary map[string]string

// Search returns the string stored under the index 'word' within the dictionary.
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

//Add adds the provided string under the defined word
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

// Update replaces an existing definition existing under 'word' with a new one.
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

// Delete deletes an existing word.
func (d Dictionary) Delete(word string) {
	delete(d, word)
}

// Search returns the map[word] from the given map
func Search(dict map[string]string, word string) string {
	return dict[word]
}
