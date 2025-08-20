// Copyright (c) ZStack.io, Inc.

package multipart

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"net/textproto"
	"sort"
	"strings"
)

type SReader struct {
	body       io.Reader
	header     string
	tail       string
	boundary   string
	headOffset int
	bodyEof    bool
	tailOffset int
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

func randomBoundary() string {
	var buf [30]byte
	_, err := io.ReadFull(rand.Reader, buf[:])
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", buf[:])
}

func buildHeader(header textproto.MIMEHeader, boundary string) string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "--%s\r\n", boundary)
	keys := make([]string, 0, len(header))
	for k := range header {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		for _, v := range header[k] {
			fmt.Fprintf(&b, "%s: %s\r\n", k, v)
		}
	}
	fmt.Fprintf(&b, "\r\n")
	return b.String()
}

func buildTail(boundary string) string {
	return fmt.Sprintf("\r\n--%s--\r\n", boundary)
}

func NewReader(r io.Reader, fieldname, filename string) *SReader {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			escapeQuotes(fieldname), escapeQuotes(filename)))
	h.Set("Content-Type", "application/octet-stream")
	boundary := randomBoundary()
	return &SReader{
		body:       r,
		header:     buildHeader(h, boundary),
		tail:       buildTail(boundary),
		boundary:   boundary,
		headOffset: 0,
		bodyEof:    false,
		tailOffset: 0,
	}
}

func (r *SReader) FormDataContentType() string {
	return "multipart/form-data; boundary=" + r.boundary
}

func (r *SReader) Read(p []byte) (n int, err error) {
	read := 0
	for read < len(p) && r.headOffset < len(r.header) {
		p[read] = r.header[r.headOffset]
		r.headOffset += 1
		read += 1
	}
	if read < len(p) && !r.bodyEof {
		n, err := r.body.Read(p[read:])
		read += n
		if err == io.EOF || n == 0 {
			r.bodyEof = true
		} else {
			return read, err
		}
	}
	for read < len(p) && r.tailOffset < len(r.tail) {
		p[read] = r.tail[r.tailOffset]
		r.tailOffset += 1
		read += 1
	}
	if read == 0 && r.tailOffset >= len(r.tail) {
		return 0, io.EOF
	} else {
		return read, nil
	}
}
