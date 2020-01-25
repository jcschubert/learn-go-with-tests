package maps

// Dictionary represents a dictionary of words, accessible by index.
type Dictionary map[string]string

// Search returns the string stored under the index 'word' within the dictionary.
func (d Dictionary) Search(word string) string {
	return d[word]
}

// Search returns the map[word] from the given map
func Search(dict map[string]string, word string) string {
	return dict[word]
}
