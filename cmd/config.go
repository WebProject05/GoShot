package cmd

type Config struct {
	File string
	StartLine int
	EndLine int

	Output string
	Theme string
	Split bool
	NoFrame bool
	Scale float64
	FontSize int
}