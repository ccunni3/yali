package yali

import (
	"fmt"
)

type ParseError struct {
	Err      error
	Encoding string
	Type     any
}

func (e ParseError) Error() string {
	return fmt.Sprintf("yali: unable to parse %s into %T", e.Encoding, e.Type)
}

func (e ParseError) Unwrap() error { return e.Err }

func (p Policies) UnmarshalText(text []byte) error {
	panic("//TODO: implement CSV and XML parsing")
	return ParseError{Err: nil, Encoding: "", Type: p}
}

func (p Policies) UnmarshalJSON(b []byte) error {
	panic("//TODO: implement JSON parsing")
	return ParseError{Err: nil, Encoding: "JSON", Type: p}
}

func (p *Policy) UnmarshalText(text []byte) error {
	panic("//TODO: implement CSV and XML parsing")
	return ParseError{Err: nil, Encoding: "", Type: p}
}

func (p *Policy) UnmarshalJSON(b []byte) error {
	panic("//TODO: implement JSON parsing")
	return ParseError{Err: nil, Encoding: "JSON", Type: p}
}

func (p *Predicate) UnmarshalText(text []byte) error {
	panic("//TODO: implement CSV and XML parsing")
	return ParseError{Err: nil, Encoding: "", Type: p}
}

func (p *Predicate) UnmarshalJSON(b []byte) error {
	panic("//TODO: implement JSON parsing")
	return ParseError{Err: nil, Encoding: "JSON", Type: p}
}
