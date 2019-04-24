package bnet

import "fmt"

type IndexOutOfRangeError struct {
	N      int64
	Offset int64
	Struct string
	Field  string
}

func (e *IndexOutOfRangeError) Error() string {
	return fmt.Sprintf("bnet: index out of range when attempting to unmarshal %d bytes starting at offset %d into Go struct field %s.%s", e.N, e.Offset, e.Struct, e.Field)
}

type InvalidSavedValueError struct {
	Expected string
	Value    string
	Struct   string
	Field    string
}

func (e *InvalidSavedValueError) Error() string {
	return fmt.Sprintf("bnet: invalid saved value for Go struct field %s.%s, expected %s, got %s", e.Struct, e.Field, e.Expected, e.Value)
}

type InvalidTagValueError struct {
	Expected string
	Value    string
	Struct   string
	Field    string
}

func (e *InvalidTagValueError) Error() string {
	return fmt.Sprintf("bnet: invalid tag value for Go struct field %s.%s, expected %s, got %s", e.Struct, e.Field, e.Expected, e.Value)
}

type InvalidValueError struct {
	Struct string
	Field  string
	Value  []byte
}

func (e *InvalidValueError) Error() string {
	return fmt.Sprintf("bnet: invalid value for Go struct field %s.%s is %#v", e.Struct, e.Field, e.Value)
}

type NilPointerError struct{}

func (e *NilPointerError) Error() string {
	return "bnet: nil pointer error"
}

type TagDefinitionRequiredError struct {
	Tag    string
	Struct string
	Field  string
}

func (e *TagDefinitionRequiredError) Error() string {
	return fmt.Sprintf("bnet: tag definition %s required for Go struct field %s.%s", e.Tag, e.Struct, e.Field)
}

type UndefinedSavedValueError struct {
	Name   string
	Struct string
	Field  string
}

func (e *UndefinedSavedValueError) Error() string {
	return fmt.Sprintf("bnet: attempting to access unsaved declaration %s for Go struct field %s.%s", e.Name, e.Struct, e.Field)
}
