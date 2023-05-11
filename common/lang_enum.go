package common

import (
	"golang.org/x/text/language"
)

type LanguageValue struct {
	Default  language.Tag
	selected language.Tag
}

func (e *LanguageValue) Set(value string) error {
	lang, err := language.Parse(value)
	if err != nil {
		return err
	}
	e.selected = lang
	return nil
}

func (e LanguageValue) String() string {
	return e.Value().String()
}

func (e LanguageValue) Value() language.Tag {
	if e.selected == language.Und {
		return e.Default
	}
	return e.selected
}
