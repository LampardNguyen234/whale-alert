package common

import "fmt"

// FormatMDLink creates a ref-text message following the MarkDown standard.
func FormatMDLink(msg string, link string) string {
	return fmt.Sprintf("[%v](%v)", msg, link)
}
