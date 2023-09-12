package jsoniter

import (
	"fmt"
	"strings"
)

// ReadObject read one field from object.
// If object ended, returns empty string.
// Otherwise, returns the field name.
func (iter *Iterator) ReadVariable() (ret string) {
	buf := strings.Builder{}
	c := iter.nextToken()
	switch c {
	case '{':
		c = iter.nextToken()
		if c != '{' {
			iter.ReportError("ReadVariable", `expect { after {, but found `+string([]byte{c}))
			return
		}

		for {
			c = iter.nextToken()
			if c == '}' {
				c = iter.nextToken()
				if c != '}' {
					iter.ReportError("ReadVariable", `expect } after }, but found `+string([]byte{c}))
					return
				}
				return buf.String()
			}
			buf.WriteByte(c)
		}
		return
	case '}':
		return "" // end of object
	default:
		iter.ReportError("ReadObject", fmt.Sprintf(`expect { or , or } or n, but found %s`, string([]byte{c})))
		return
	}
}
