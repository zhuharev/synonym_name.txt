package synonym

import (
	"bufio"
	"os"
	"strings"
)

type dashComaReader struct {
	f *os.File
}

func newDashComaReader(filepath string, loadInMem bool) (*dashComaReader, error) {
	var (
		e   error
		dcr = new(dashComaReader)
	)
	dcr.f, e = os.Open(filepath)
	return dcr, e
}

func (drc *dashComaReader) Parent(name string) (parent string) {
	br := bufio.NewReader(drc.f)
	var fullLine string
	for line, isPrefix, e := br.ReadLine(); e == nil; line, isPrefix, e = br.ReadLine() {
		fullLine += string(line)
		if isPrefix {
			continue
		}

		arr := strings.SplitN(fullLine, "-", 2)
		if len(arr) != 2 {
			fullLine = ""
			continue
		}

		currentParrent := strings.TrimSpace(arr[0])

		if strings.EqualFold(currentParrent, name) {
			return name
		}

		for _, v := range strings.Split(arr[1], ",") {
			if strings.EqualFold(strings.TrimSpace(v), name) {
				return currentParrent
			}
		}

		fullLine = ""
	}

	return ""
}
