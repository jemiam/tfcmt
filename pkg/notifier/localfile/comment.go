package localfile

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
)

type OutputService service

// Post posts comment
func (f *OutputService) Post(ctx context.Context, body string, outputFile string) error {
	logE := logrus.WithFields(logrus.Fields{
		"program": "tfcmt",
	})

	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err2 := file.WriteString(body)

	if err2 != nil {
		return err2
	}
	logE.Debug("Output to file success")

	return err
}
