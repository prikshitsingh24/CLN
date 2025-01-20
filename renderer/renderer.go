package renderer

import (
	"cln/logger"

	"github.com/charmbracelet/glamour"
)

func Render(text string) string {

	out, err := glamour.Render(text, "dark")
	if err != nil {
		logger.Error("renderer", "Not able to render the text")
		return ""
	}

	return out

}
