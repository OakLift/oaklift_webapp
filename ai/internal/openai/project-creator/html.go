package project

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
)

const (
	HTMLCSSLevelsPath = "%s/internal/openai/project-creator/levels/html-css.json"
)

type HTMLCSSProject struct {
	PromptIntro string `json:"prompt-intro"`
	ProjectStructure string `json:"project-structure"`
	PromptConfirmation string `json:"prompt-confirmation"`

	Beginner HTMLCSSTopics `json:"beginner"`
	Intermediate HTMLCSSTopics `json:"intermediate"`
	Advanced HTMLCSSTopics `json:"advanced"`
	Expert HTMLCSSTopics `json:"expert"`
}

type HTMLCSSTopics struct {
	HTMLTopics []string `json:"html-topics"`
	CSSTopics []string `json:"css-topics"`
}

func CreateHTMLProject() (*HTMLCSSProject, error) {
	workingDir, err := os.Getwd()
	if err != nil {
		return nil, errors.Wrap(err, "os.Getwd() failed")
	}

	levelJSONBytes, err := os.ReadFile(fmt.Sprintf(HTMLCSSLevelsPath, workingDir))
	if err != nil {
		return nil, errors.Wrapf(err, "os.ReadFile(%s) failed", HTMLCSSLevelsPath)
	}

	projCurriculum:= &HTMLCSSProject{}
	err = json.Unmarshal(levelJSONBytes, projCurriculum)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal HTML/CSS project curriculum")
	}

	return projCurriculum, nil
}

func (h *HTMLCSSProject) BuildPrompt(level ProjectLevel) (string, error) {
	htmlTopics := ""
	cssTopics := ""

	switch level {
	case Beginner:
		htmlTopics = strings.Join(h.Beginner.HTMLTopics, ", ")
		cssTopics = strings.Join(h.Beginner.CSSTopics, ", ")
	case Intermediate:
		htmlTopics = strings.Join(h.Intermediate.HTMLTopics, ", ")
		cssTopics = strings.Join(h.Intermediate.CSSTopics, ", ")
	case Advanced:
		htmlTopics = strings.Join(h.Advanced.HTMLTopics, ", ")
		cssTopics = strings.Join(h.Advanced.CSSTopics, ", ")
	case Expert:
		htmlTopics = strings.Join(h.Expert.HTMLTopics, ", ")
		cssTopics = strings.Join(h.Expert.CSSTopics, ", ")
	default:
		return "", errors.Errorf("unknown HTML/CSS level: %d", level)
	}

	prompt := fmt.Sprintf("%s\n\nHTML Topics: %s\n\nCSS Topics: %s\n\n%s", h.PromptIntro, htmlTopics, cssTopics, h.PromptConfirmation)
	return prompt, nil
}

