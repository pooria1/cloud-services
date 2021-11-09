package handler

import (
	"fmt"
	"strings"

	"github.com/pooria1/cloud-services/backend/database"
)

func addCommentToHTML(commentText string, commentNumber int) (string, error) {
	// audio, err := tts.ConvertTextToSpeech(commentText)
	// if err != nil {
	// 	return "", err
	// }
	commentLine := fmt.Sprintf(commentTemplate, commentText, commentText, commentNumber)
	htmlTemplateVariable = strings.ReplaceAll(htmlTemplateVariable, fmt.Sprintf("<!--cp%d-->", commentNumber), commentLine)
	if err := database.AddToDB(fmt.Sprintf("%d", commentNumber), commentText); err != nil {
		return "", err
	}
	return htmlTemplateVariable, nil
}
