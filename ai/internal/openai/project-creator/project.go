package project

import (
	"encoding/json"
	"log"

	"github.com/pkg/errors"
)

type Language int
const (
	HTMLCSS Language = iota
	Python
	Javascript
	Golang
	UnknownProject
)

type ProjectLevel int
const (
	Beginner ProjectLevel = iota
	Intermediate
	Advanced
	Expert
	Unknown
)

func CreatePrompt(language Language, level ProjectLevel) (string, error) {
	switch language {
	case HTMLCSS:
		htmlProject, err := CreateHTMLProject()
		if err != nil {
			return "", errors.Wrap(err, "failed to create HTML/CSS project")
		}

		prompt, err := htmlProject.BuildPrompt(level)
		if err != nil {
			return "", errors.Wrap(err, "failed to build prompt for HTML/CSS project")
		}
		return prompt, nil
	case Python:
		return "", errors.New("Python project not implemented")
	case Javascript:
		return "", errors.New("Javascript project not implemented")
	case Golang:
		return "", errors.New("Golang project not implemented")
	default:
		return "", errors.Errorf("unknown project language %d", language)
	}
}

func PrintCurriculum() error {
	htmlCSSProject, err := CreateHTMLProject()
	if err != nil {
		return errors.Errorf("failed to parse HTML/CSS project curriculum: %v", err)
	}

	prettyJSON, err := json.MarshalIndent(htmlCSSProject, "", " ")
	if err != nil {
		return errors.Wrap(err, "json.MarshalIndent failure for HTML/CSS project curriculum")
	}

	
	log.Printf("HTML/CSS project curriculum: %s", string(prettyJSON))
	return nil
}


func PrintHTMLCSSPrompt(level ProjectLevel) error {
	htmlCSSProject, err := CreateHTMLProject()
	if err != nil {
		return errors.Errorf("failed to parse HTML/CSS project curriculum: %v", err)
	}

	prompt, err := htmlCSSProject.BuildPrompt(level)
	if err != nil {
		return errors.Wrap(err, "failed to build prompt for HTML/CSS project")
	}

	log.Printf("Prompt:\n%s",  prompt)
	return nil
}
