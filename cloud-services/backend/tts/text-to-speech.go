package tts

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

// ConvertTextToSpeech converts given text using
// google translate and downloads speech then
// stores the mp3 file into given path
func ConvertTextToSpeech(text string) (string, error) {

	var err error
	generatedHashName := generateHashName(text)
	fileName := "/var/www/html/audio/" + generatedHashName + ".mp3"

	if err = createFolderIfNotExists("/var/www/html/audio"); err != nil {
		return "", err
	}
	if err = downloadIfNotExists(fileName, text); err != nil {
		return "", err
	}
	return generatedHashName + ".mp3", nil
}

func createFolderIfNotExists(folder string) error {
	dir, err := os.Open(folder)
	if os.IsNotExist(err) {
		return os.MkdirAll(folder, 0744)
	}

	dir.Close()
	return nil
}

// downloadIfNotExists does not allow the programm
// to redownload old mp3 files
func downloadIfNotExists(fileName string, text string) error {
	f, err := os.Open(fileName)
	if err != nil {
		url := fmt.Sprintf("http://translate.google.com/translate_tts?ie=UTF-8&total=1&idx=0&textlen=32&client=tw-ob&q=%s&tl=%s", url.QueryEscape(text), "en")
		response, err := http.Get(url)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		output, err := os.Create(fileName)
		if err != nil {
			return err
		}

		_, err = io.Copy(output, response.Body)
		return err
	}

	f.Close()
	return nil
}

func generateHashName(name string) string {
	hash := md5.Sum([]byte(name))
	return fmt.Sprintf("%s_%s", "en", hex.EncodeToString(hash[:]))
}
