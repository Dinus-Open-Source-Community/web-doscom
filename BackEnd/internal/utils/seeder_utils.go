package utils

import (
	"bytes"
	"fmt"
	"mime"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
)

func OpenSeedImage(path string) (
	*multipart.FileHeader,
	multipart.File,
	func(),
	error,
) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, nil, err
	}

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	contentType := http.DetectContentType(content)
	header := make(textproto.MIMEHeader)
	header.Set(
		"Content-Disposition",
		mime.FormatMediaType("form-data", map[string]string{
			"name":     "file",
			"filename": filepath.Base(path),
		}),
	)
	header.Set("Content-Type", contentType)
	part, err := writer.CreatePart(header)
	if err != nil {
		return nil, nil, nil, err
	}

	if _, err := part.Write(content); err != nil {
		return nil, nil, nil, err
	}

	if err := writer.Close(); err != nil {
		return nil, nil, nil, err
	}

	reader := multipart.NewReader(&body, writer.Boundary())

	form, err := reader.ReadForm(10 << 20)
	if err != nil {
		return nil, nil, nil, err
	}

	headers := form.File["file"]
	if len(headers) == 0 {
		form.RemoveAll()
		return nil, nil, nil, fmt.Errorf("file header not found")
	}

	fileHeader := headers[0]
	file, err := fileHeader.Open()
	if err != nil {
		form.RemoveAll()
		return nil, nil, nil, err
	}

	cleanup := func() {
		_ = file.Close()
		_ = form.RemoveAll()
	}

	return fileHeader, file, cleanup, nil

}
