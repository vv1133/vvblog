package common

import (
	"fmt"
	"github.com/vv1133/vvblog/models"
	"gopkg.in/mgo.v2/bson"
	"regexp"
	"strings"
	"time"
)

func filterHtml(str string) string {
	// 将Html标签全部转换为小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	str = re.ReplaceAllStringFunc(str, strings.ToLower)

	// 去除Style
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	str = re.ReplaceAllString(str, "")

	// 去Script
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	str = re.ReplaceAllString(str, "")

	// 去除所有尖括号内的Html代码, 并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	str = re.ReplaceAllString(str, "\n")

	// 去除连续的换行符
	re, _ = regexp.Compile("\\S\\s{2,}")
	str = re.ReplaceAllString(str, "\n")

	return str
}

func Preview(str string, length int) string {
	str = filterHtml(str)
	rs := []rune(str)
	rl := len(rs)

	if length > rl {
		str = string(rs[0:rl])
	} else {
		str = string(rs[0:length]) + "..."
	}

	return strings.Replace(str, "\n", "", -1)
}

func GetId(id bson.ObjectId) string {
	return id.Hex()
}

func GetTagSlug(caption string) string {
	var tag models.BlogTag

	models.GetOneByQuery(models.DbTag, bson.M{"caption": caption}, &tag)

	return tag.Slug
}

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
