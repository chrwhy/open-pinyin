package dict

import (
	"bufio"
	"log"
	"os"
	"strings"
	"unicode"
)

var PINYIN = make(map[string]string)
var CN_PINYIN = make(map[string][]string)
var PINYIN_PREFIX = make(map[string]string)
var NOT_SPLIT = make(map[string]string)

const (
	PINYIN_DICT    = "pinyin.dict"
	CN_PINYIN_DICT = "cn_pinyin.dict"
	N              = "n"
	G              = "g"
	NG             = "ng"
	ER             = "er"
)

func init() {
	loadPinyin()
	loadCnPinyin()
}

func loadPinyin() {
	file, err := os.Open(PINYIN_DICT)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		PINYIN[line] = line
		if !IsIuv(line) {
			if len(line) > 1 {
				for i := 1; i <= len(line); i++ {
					PINYIN_PREFIX[line[0:i]] = line[0:i]
				}
			} else {
				PINYIN_PREFIX[line] = line
			}
		}
	}

	NOT_SPLIT[ER] = ER
}

func loadCnPinyin() {
	file, err := os.Open(CN_PINYIN_DICT)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		mapping := strings.Split(line, "=")
		CN_PINYIN[mapping[0]] = strings.Split(mapping[1], ",")
	}
}

func IsPinyinPrefix(pinyin string) bool {
	_, prefixMatched := PINYIN_PREFIX[pinyin]
	return prefixMatched
}

func IsPinyin(pinyin string) bool {
	_, matched := PINYIN[pinyin]
	return matched
}

func IsLegalPinyin(pinyin string) bool {
	if len(pinyin) == 1 {
		if "a" == pinyin || "o" == pinyin || "e" == pinyin {
			return true
		}
		return false
	}

	if len(pinyin) == 2 && (pinyin == "zh" || pinyin == "ch" || pinyin == "sh" || pinyin == "ng") {
		return false
	}

	_, matched := PINYIN[pinyin]
	return matched
}

func IsIuv(pinyin string) bool {
	return pinyin == "i" || pinyin == "u" || pinyin == "v"
}

func HasIuv(pinyin string) bool {
	return strings.Contains(pinyin, "i") || strings.Contains(pinyin, "u") || strings.Contains(pinyin, "v")
}

func GetCnPinyin(cn string) []string {
	if _, ok := CN_PINYIN[cn]; ok {
		return CN_PINYIN[cn]
	} else {
		if unicode.Is(unicode.Han, []rune(cn)[0]) {
			log.Println(cn, "has no pinyin...")
		}
		return nil
	}
}
