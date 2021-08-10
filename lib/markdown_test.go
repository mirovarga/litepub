package lib

import "testing"

func TestMarkdownToPostWithLF(t *testing.T) {
	lf := "# A title\n\n*Aug 10, 2021*\n\n*Test, Markdown*\n\nTesting LF\n"

	_, err := markdownToPost(lf)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMarkdownToPostWithCRLF(t *testing.T) {
	crlf := "# A title\r\n\r\n*Aug 10, 2021*\r\n\r\n*Test, Markdown*\r\n\r\nTesting CRLF\r\n"

	_, err := markdownToPost(crlf)
	if err != nil {
		t.Fatal(err)
	}
}
