package service

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func InitPrinter(lang string) *message.Printer {
	var langTag language.Tag
	if lang == "" {
		langTag = language.English
	} else {
		langTag = language.MustParse(lang)
	}

	return message.NewPrinter(langTag)
}
