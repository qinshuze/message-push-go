package language

import "strings"

type Entry struct {
	value string
}

func NewEntry(value string) *Entry {
	return &Entry{value: value}
}

func (e *Entry) Value() string {
	return e.value
}

func (e *Entry) Default(v string) string {
	if e.value == "" {
		return v
	}

	return e.value
}

func (e *Entry) Replace(replace map[string]string) *Entry {
	for s, s2 := range replace {
		e.value = strings.ReplaceAll(e.value, ":"+s, s2)
	}

	return e
}
