package golevel7

const (
	eof = rune(0)
	LF  = '\x0A'
	CR  = '\x0D'
)

// Delimeters holds the list of hl7 message delimeters
type Delimeters struct {
	DelimeterField string
	Field          rune
	Component      rune
	Repetition     rune
	Escape         rune
	SubComponent   rune
	Truncate       rune
}

// NewDelimeters returns the default set of HL7 delimeters
func NewDelimeters() *Delimeters {
	return &Delimeters{
		DelimeterField: "^~\\&",
		Field:          '|',
		Component:      '^',
		Repetition:     '~',
		Escape:         '\\',
		SubComponent:   '&',
		Truncate:       '#',
	}
}
