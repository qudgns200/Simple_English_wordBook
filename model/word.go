package word

type Word struct {
	SWord    string
	Meanings string
	Example  string
}

// method for type SearchWord
func (word *Word) String() string {
	return word.SWord + " " + word.Meanings + " " + word.Example
}
