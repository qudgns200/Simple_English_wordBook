package word

type Word struct {
	sWord    string
	meanings string
	example  string
}

func (word *Word) GetSWord() string {
	return word.sWord
}

func (word *Word) GetMeanings() string {
	return word.meanings
}

func (word *Word) GetExample() string {
	return word.example
}

func (word *Word) SetSWord(str string) {
	word.sWord = str
}

func (word *Word) SetMeanings(str string) {
	word.meanings = str
}

func (word *Word) SetSExample(str string) {
	word.example = str
}

// method for type SearchWord
func (word *Word) String() string {
	return word.GetSWord() + " " + word.GetMeanings() + " " + word.GetExample()
}
