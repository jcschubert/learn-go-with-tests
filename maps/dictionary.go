package maps

import "errors"

// ErrNotFound should be used if a key in the dict couldn't be found
var ErrNotFound = errors.New("could not find the word you were looking for")

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

// Search returns the map[word] from the given map
func Search(dict map[string]string, word string) string {
	return dict[word]
}
