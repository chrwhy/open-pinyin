package dict

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var PINYIN = make(map[string]string)
var CN_PINYIN = make(map[string][]string)
var PINYIN_PREFIX = make(map[string]string)

func init() {
	loadPinyin()
	loadCnPinyin()
}

func loadPinyin() {
	file, err := os.Open("pinyin.dict")
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
				for i := 2; i < len(line); i++ {
					PINYIN_PREFIX[line[0:i]] = line[0:i]
				}
			} else {
				PINYIN_PREFIX[line] = line
			}
		}
	}
}

func loadCnPinyin() {
	file, err := os.Open("cn_pinyin.dict")
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

	if len(pinyin) == 2 && (pinyin == "zh" || pinyin == "ch" || pinyin == "sh") {
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
