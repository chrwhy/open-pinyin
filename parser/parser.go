package parser

import (
	"github.com/chrwhy/open-pinyin/dict"
	"log"
	"strings"
	"time"
)

type PinyinNode struct {
	Pinyin        string
	Leftover      string
	DirectedNodes []*PinyinNode
	Illegal       bool
}

func ParseInitial(input string) []string {
	result := make([]string, 0)
	for _, character := range input {
		if dict.IsIuv(string(character)) {
			return nil
		}
		result = append(result, string(character))
	}

	return result
}

func Parse(text string) [][]string {
	root := &PinyinNode{}
	root.DirectedNodes = make([]*PinyinNode, 0)
	root.Leftover = strings.ToLower(text)
	t1 := time.Now()
	parsePinyinDAG(root, make(map[string][]*PinyinNode))
	pinyinGroups := Traverse(root)
	log.Println("DAG way cost:", time.Since(t1))
	//pinyinGroups = append(pinyinGroups, ParseInitial(text))
	return pinyinGroups
}

func Traverse(root *PinyinNode) [][]string {
	rawPinyinGroups := TraverseDAG([]string{}, root)
	pinyinGroups := make([][]string, 0)
	for _, rawPinyinGroup := range rawPinyinGroups {
		nodePath := rawPinyinGroup
		legal := true
		for i := 0; i < len(nodePath); i++ {
			step := nodePath[i]
			if i == len(nodePath)-1 {
				if !dict.IsPinyinPrefix(step) {
					legal = false
					break
				}
			} else {
				if !dict.IsLegalPinyin(step) {
					legal = false
					break
				}
			}
		}

		if len(nodePath) < 1 {
			continue
		}

		if legal {
			//log.Println("legal: ", nodePath)
			pinyinGroups = append(pinyinGroups, nodePath)
		} else {
			//log.Println("illegal: ", nodePath)
			pinyinGroups = append(pinyinGroups, nodePath)
		}
	}

	return pinyinGroups
}

func TraverseDAG(prefix []string, root *PinyinNode) [][]string {
	result := make([][]string, 0)
	if root.Illegal {
		return nil
	}
	if len(root.DirectedNodes) < 1 {
		/*
			prefix = append(prefix, root.Pinyin) FUCK this...
			for _, s := range prefix {
				temp = append(temp, s)
			}
		*/
		temp := make([]string, len(prefix))
		copy(temp, prefix)
		if len(root.Pinyin) > 0 {
			temp = append(temp, root.Pinyin)
		}
		return [][]string{temp}
	} else {
		temp := make([]string, 0)
		if len(root.Pinyin) > 0 {
			temp = append(prefix, root.Pinyin)
		}
		for _, child := range root.DirectedNodes {
			result = append(result, TraverseDAG(temp, child)...)
		}

		return result
	}
}

func parsePinyinDAG(node *PinyinNode, parseCache map[string][]*PinyinNode) {
	if node.Leftover == "" {
		return
	}
	raw := node.Leftover
	if len(parseCache[raw]) > 0 {
		node.DirectedNodes = parseCache[raw]
		return
	}

	head := GreedyFirst(node.Leftover)
	if head == "" && len(node.Leftover) > 0 {
		node.Illegal = true
		return
	}
	//log.Println("Greedy head:", head)
	node.DirectedNodes = make([]*PinyinNode, 0)
	//candidates := ElectCandidates(head)
	candidates := make([]string, 0)
	if len(node.Leftover[len(head):]) > 0 {
		nextChar := node.Leftover[len(head):][0:1]
		if Splittable(head) {
			candidates = ElectCandidatesV2(head, nextChar)
		}
		if strings.HasSuffix(head, dict.N) || strings.HasSuffix(head, dict.G) {
			//jianing case
			if dict.IsPinyinPrefix(nextChar) {
				candidates = append(candidates, head)
			}
		} else {
			candidates = append(candidates, head)
		}
	} else {
		candidates = ElectCandidatesV2(head, "")
		candidates = append(candidates, head)
	}

	for _, candidate := range candidates {
		leftover := raw[len(candidate):]
		child := &PinyinNode{Pinyin: candidate, Leftover: leftover}
		parseCache[node.Leftover] = append(parseCache[node.Leftover], child)
		node.DirectedNodes = append(node.DirectedNodes, child)
		parsePinyinDAG(child, parseCache)
	}
}

func Splittable(text string) bool {
	if _, ok := dict.NOT_SPLIT[text]; ok {
		return false
	}

	return true
}

func ElectCandidates(text string) []string {
	result := make([]string, 0)
	candidate := ""
	for i, t := range text {
		if i == len(text)-1 {
			continue
		}
		//leftover := text[i:]
		candidate += string(t)
		if dict.IsLegalPinyin(candidate) {
			result = append(result, candidate)
		}
	}

	//result = append(result, text)
	//log.Println(result)
	return result
}

func ElectCandidatesV2(text, nextChar string) []string {
	result := make([]string, 0)
	candidate := ""

	for i, t := range text {
		if i == len(text)-1 {
			continue
		}
		isNgSuffixAndPrefix := true
		lastChar := text[len(text)-1:]
		if (i == len(text)-2) && lastChar == dict.G {
			if !dict.IsPinyinPrefix(dict.G + nextChar) {
				isNgSuffixAndPrefix = false
			}
		}

		if (i == len(text)-2) && lastChar == dict.N {
			if !dict.IsPinyinPrefix(dict.N + nextChar) {
				isNgSuffixAndPrefix = false
			}
		}

		candidate += string(t)
		leftover := text[i+1:]
		if leftover == dict.NG {
			continue
		}

		if len(leftover) == 1 && !dict.IsPinyinPrefix(leftover) {
			continue
		}

		if dict.IsLegalPinyin(candidate) && isNgSuffixAndPrefix {
			result = append(result, candidate)
		}
	}

	return result
}

func GreedyFirst(text string) string {
	candidate := ""
	leftover := ""

	if len(text) < 1 {
		return ""
	}

	if len(text) < 6 {
		candidate = text
		leftover = ""
	} else {
		candidate = text[0:6]
		leftover = text[6:]
	}

	if dict.IsPinyin(candidate) || (leftover == "" && dict.IsPinyinPrefix(candidate)) {
		return candidate
	} else {
		pinyin, cutLeftover := maxCut(candidate)
		if pinyin == "" || cutLeftover == "" {
			log.Println("!!!!!!!!!!!!!!!!!!!!!!")
			return pinyin
		}
		return pinyin
	}
}

func GreedyParse(text string) []string {
	finalResult := make([]string, 0)
	candidate := ""
	leftover := ""

	if len(text) < 1 {
		return finalResult
	}

	if len(text) < 6 {
		candidate = text
		leftover = ""
	} else {
		candidate = text[0:6]
		leftover = text[6:]
	}

	if dict.IsPinyin(candidate) || (leftover == "" && dict.IsPinyinPrefix(candidate)) {
		finalResult = append(finalResult, candidate)
		finalResult = append(finalResult, GreedyParse(leftover)...)
	} else {
		pinyin, cutLeftover := maxCut(candidate)
		if pinyin == "" || cutLeftover == "" {
			log.Println("!!!!!!!!!!!!!!!!!!!!!!")
		}

		finalResult = append(finalResult, pinyin)
		finalResult = append(finalResult, GreedyParse(cutLeftover+leftover)...)
	}

	return finalResult
}

func minCut(text string) (pinyin, leftover string) {
	for i := 0; i < len(text); i++ {
		candidate := text[0:i]
		if dict.IsLegalPinyin(candidate) {
			leftover = text[i:]
			return candidate, leftover
		}
	}

	if dict.IsLegalPinyin(text) {
		return text, ""
	} else {
		return text[0:1], text[1:]
	}
}

func maxCut(text string) (pinyin, leftover string) {
	for i := len(text) - 1; i > 0; i-- {
		candidate := text[0:i]
		if dict.IsPinyin(candidate) && !dict.IsIuv(candidate) && !dict.Is2LetterConsonant(candidate) {
			leftover = text[i:]
			return candidate, leftover
		}
	}

	return "", text
}
