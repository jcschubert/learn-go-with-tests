# Maps

Arrays: How to store value in order, one after the other. Maps: Like dictionaries
in python. Store values by key, so we can look them up quickly.

Maps are declared with two types, the key type and the value type. The key type
must be a comparable type.

Example is CRUD, wrapped in error handling. The dictionary type, and its most common,
built-in methods.

## Note on Maps as reference types

Maps are reference types in go, so are never copied when used as function arguments.

## Notes on the error handling in this Example

I don't agree with how errors are used here, like in the Add Example (see `maps/dictionary.go`).
When adding a word, not finding it in the dictionary is not an error case, but one that is normal.
These shouldn't be handled via errors, but by additional methods that query the state of the dictionary
before adding, and returning some form of state instead of errors.

I'll have to check what the common way to use errors in Go is.
