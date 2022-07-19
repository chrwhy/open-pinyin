package parser

import (
	"log"
	"pinyin/dict"
	"strings"
)

type TreeNode struct {
	Pinyin   string
	Leftover string
	Children []*TreeNode
	Parent   *TreeNode
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
	root := &TreeNode{}
	root.Children = make([]*TreeNode, 0)
	root.Leftover = text
	parseNode(root)
	leaveNodes := TraverseLeaveNodes(root)
	pinyinGroups := make([][]string, 0)
	for _, leaveNode := range leaveNodes {
		nodePath := ReversePath(leaveNode)
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
		if legal {
			log.Println("legal: ", nodePath)
			pinyinGroups = append(pinyinGroups, nodePath)
		} else {
			//log.Println("illegal: ", nodePath)
		}
	}

	return pinyinGroups
}

func TraverseLeaveNodes(root *TreeNode) []*TreeNode {
	result := make([]*TreeNode, 0)

	if len(root.Children) < 1 {
		result = append(result, root)
		return result
	} else {
		for _, child := range root.Children {
			result = append(result, TraverseLeaveNodes(child)...)
		}
		return result
	}
}

func parseNode(node *TreeNode) {
	raw := node.Leftover
	head := GreedyFirst(node.Leftover)
	if head == "" || node.Leftover == "" {
		return
	}
	log.Println("Greedy head:", head)
	node.Children = make([]*TreeNode, 0)
	candidates := ElectCandidates(head)
	candidates = append(candidates, head)
	for _, candidate := range candidates {
		leftover := raw[len(candidate):]
		log.Println("Candidate:", candidate, "Leftover:", leftover)
		child := &TreeNode{Pinyin: candidate, Leftover: leftover}
		child.Parent = node
		node.Children = append(node.Children, child)
		parseNode(child)
	}
}

func ReversePath(node *TreeNode) []string {
	pathNodes := make([]string, 0)

	for {
		temp := []string{node.Pinyin}
		if len(node.Pinyin) > 0 {
			temp = append(temp, pathNodes...)
			pathNodes = temp
		}
		//pathNodes = append(pathNodes, node.Pinyin)
		if node.Parent == nil {
			break
		} else {
			node = node.Parent
		}
	}
	//log.Println("here:", pathNodes)
	return pathNodes
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
	} else {
		pinyin, cutLeftover := maxCut(candidate)
		if pinyin == "" || cutLeftover == "" {
			log.Println("!!!!!!!!!!!!!!!!!!!!!!")
		}

		finalResult = append(finalResult, pinyin)
		finalResult = append(finalResult, GreedyParse(cutLeftover+leftover)...)
	}

	concatPinyin := ""
	for _, pinyin := range finalResult {
		concatPinyin += pinyin + " "
	}
	return []string{strings.TrimSpace(concatPinyin)}
}

func minCut(text string) (pinyin, leftover string) {
	for i := 0; i < len(text); i++ {
		candidate := text[0:i]
		if dict.IsLegalPinyin(candidate) {
			leftover = text[i:]
			return candidate, leftover
		}
	}

	if dict.IsLegalPinyin(pinyin) {
		return text, ""
	} else {
		return text[0:1], text[1:]
	}
}

func maxCut(text string) (pinyin, leftover string) {
	for i := len(text) - 1; i > 0; i-- {
		candidate := text[0:i]
		if dict.IsPinyin(candidate) {
			leftover = text[i:]
			return candidate, leftover
		}
	}

	return text, ""
}
