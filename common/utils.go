package common

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

func GetSlug(str string, isslug bool) string {
	retstr := ""

	for _, i := range str {
		inside_code := i
		if inside_code == 12288 {
			inside_code = 32
		} else {
			inside_code -= 65248
		}
		if inside_code < 32 || inside_code > 126 {
			retstr += string(i)
		} else {
			retstr += string(inside_code)
		}
	}

	reg := regexp.MustCompile(`[\pP]+`)
	str = reg.ReplaceAllString(retstr, "")
	str = strings.TrimSpace(str)
	reg = regexp.MustCompile(`[\sS]+`)
	str = reg.ReplaceAllString(retstr, "")
	str = strings.Replace(str, " ", "-", -1)

	return str
}

func LoadTimes(startTime time.Time) string {
	return fmt.Sprintf("%dms", time.Now().Sub(startTime)/1000000)
}
