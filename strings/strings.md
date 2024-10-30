# strings Package

the strings package provides various functions to manipulate and work with string data.

## Creating and Modifying Strings

- `strings.TrimSpace(s string)`: Removes leading and trailing whitespace from the string `s`.
- `strings.ToLower(s string)`: Converts the string `s` to lowercase.
- `strings.ToUpper(s string)`: Converts the string `s` to uppercase.
- `strings.Title(s string)`: Converts the first character of each word in the string `s` to uppercase.

## Searching and Replacing

- `strings.Contains(s, substr string)`: Checks if the string s contains the substring substr.
- `strings.Index(s, substr string)`: Returns the index of the first occurrence of substr in s, or -1 if not found.
- `strings.Replace(s, old, new string, n int)`: Replaces the first n occurrences of old with new in the string s.

## Splitting and Joining

- `strings.Split(s, sep string)`: Splits the string s into a slice of substrings separated by the string sep.
- `strings.Join(a []string, sep string)`: `Joins` the elements of the slice a into a single string, separated by the string sep.
