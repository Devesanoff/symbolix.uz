package usecase

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"io"

	"github.com/Devesanoff/symbolix.uz/pkg/translator"
	"github.com/ledongthuc/pdf"
)


func CyrillicToLatin(text string) string {
	return translator.CyrillicToLatin(text)
}


func LatinToCyrillic(text string) string {
	return translator.LatinToCyrillic(text)
}


func PDFToDocx(pdfStream io.Reader, size int64, translate bool) ([]byte, error) {
	readerAt, ok := pdfStream.(io.ReaderAt)
	if !ok {
		data, err := io.ReadAll(pdfStream)
		if err != nil {
			return nil, err
		}
		readerAt = bytes.NewReader(data)
		size = int64(len(data))
	}

	pdfReader, err := pdf.NewReader(readerAt, size)
	if err != nil {
		return nil, err
	}

	var totalText bytes.Buffer
	for i := 1; i <= pdfReader.NumPage(); i++ {
		page := pdfReader.Page(i)
		if page.V.IsNull() {
			continue
		}
		content, err := page.GetPlainText(nil)
		if err == nil {
			totalText.WriteString(content)
			totalText.WriteString("\n")
		}
	}

	text := totalText.String()
	if translate {
		text = translator.CyrillicToLatin(text)
	}

	return createDocx(text)
}

func createDocx(text string) ([]byte, error) {
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)

	
	contentTypes := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<Types xmlns="http://schemas.openxmlformats.org/package/2006/content-types">
  <Default Extension="rels" ContentType="application/vnd.openxmlformats-package.relationships+xml"/>
  <Default Extension="xml" ContentType="application/xml"/>
  <Override PartName="/word/document.xml" ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml"/>
</Types>`
	f, _ := w.Create("[Content_Types].xml")
	f.Write([]byte(contentTypes))

	
	rels := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<Relationships xmlns="http://schemas.openxmlformats.org/package/2006/relationships">
  <Relationship Id="rId1" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument" Target="word/document.xml"/>
</Relationships>`
	f, _ = w.Create("_rels/.rels")
	f.Write([]byte(rels))

	
	var textBuf bytes.Buffer
	xml.EscapeText(&textBuf, []byte(text))
	
	paragraphs := ""
	lines := bytes.Split(textBuf.Bytes(), []byte("\n"))
	for _, line := range lines {
		paragraphs += `<w:p><w:r><w:t>` + string(line) + `</w:t></w:r></w:p>`
	}

	docXML := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<w:document xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">
  <w:body>` + paragraphs + `</w:body>
</w:document>`
	
	f, _ = w.Create("word/document.xml")
	f.Write([]byte(docXML))

	w.Close()
	return buf.Bytes(), nil
}
