package password

type Password struct {
	Value    []byte
	Filename string
	Passname string
	Copied   bool
	Saved bool
}