package main

import (
	"bufio"
	"fmt"
	"github.com/chrwhy/open-pinyin/parser"
	"github.com/chrwhy/open-pinyin/util"
	"log"
	"os"
	"strings"
)

func TestCase() {
	inputs := make([]string, 0)
	inputs = append(inputs, "chenhairong")
	inputs = append(inputs, "xianrenmin")
	inputs = append(inputs, "xianr")
	inputs = append(inputs, "aqiang")
	inputs = append(inputs, "ana")
	inputs = append(inputs, "abc")
	inputs = append(inputs, "aaijifeji")
	inputs = append(inputs, "yiqungaoguiqizhidechairenzaichufaweizhangdongwu")
	inputs = append(inputs, "lianggehuanglimingcuiliu")
	inputs = append(inputs, "renshengdeyixunjinhuanmoshijinzunkongduiyue")
	inputs = append(inputs, "ziranyuyanchulishiyigelishinantideyuyanzhedetianxia")
	inputs = append(inputs, "bairiyishanj")
	inputs = append(inputs, "luxian")
	inputs = append(inputs, "pdfzhuanhuachengword")
	inputs = append(inputs, "pdfzenmezhuanchengexcel")
	inputs = append(inputs, "chuangqianmingyueguang")
	inputs = append(inputs, "angui")
	inputs = append(inputs, "ning")
	inputs = append(inputs, "xiongge")
	inputs = append(inputs, "liaojie")
	inputs = append(inputs, "liangting")
	inputs = append(inputs, "jianing")
	inputs = append(inputs, "libai")
	inputs = append(inputs, "luxian")
	inputs = append(inputs, "lixiaoguang")
	inputs = append(inputs, "chuangqianqianmingyueguang")
	inputs = append(inputs, "xiangbiyudaduoshurenshuxideshujukudesuoyinxiaolvshangshiwanbaochuantongshujukudexingneng")

	path := "./output.txt"
	f, _ := os.Create(path)
	defer f.Close()

	for _, input := range inputs {
		pinyinGroups := parser.Parse(input)
		f.WriteString(input + "\n")

		for _, pinyinGroup := range pinyinGroups {
			f.WriteString(util.Concat(pinyinGroup, " ") + "\n")
		}

		f.WriteString("=========================\n\n")
	}
}

func Test(input string) {
	result := parser.Parse(input)
	log.Println(input, ":")
	for i, _ := range result {
		log.Println(result[i])
	}
}

func main() {
	//TestCase()
	//Test("yueguang")
	//Test("anzn")
	main1()
}

type ParseCase struct {
	Input  string
	Output [][]string
}

func main1() {
	// 打开文件
	file, err := os.Open("output.txt")
	if err != nil {
		fmt.Println("打开文件出错:", err)
		return
	}
	defer file.Close()

	var cases []ParseCase
	var currentGroup [][]string
	var currentInput string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// 跳过分隔符和空行
		if line == "" || strings.Contains(line, "=========================") {
			if len(currentGroup) > 0 {
				cases = append(cases, ParseCase{Input: currentInput, Output: currentGroup})
				currentGroup = nil
			}
			currentInput = ""
			continue
		}

		// 如果是组的第一行，作为输入
		if currentInput == "" {
			currentInput = line
		} else {
			// 将行分割成单词并添加到当前组
			words := strings.Fields(line)
			currentGroup = append(currentGroup, words)
		}
	}

	// 添加最后一组（如果不为空）
	if len(currentGroup) > 0 {
		cases = append(cases, ParseCase{Input: currentInput, Output: currentGroup})
	}

	// 按指定格式打印结果
	for _, c := range cases {
		fmt.Printf("cases = append(cases, ParseCase{\"%s\", [][]string{", c.Input)
		for i, group := range c.Output {
			fmt.Printf("{")
			for j, word := range group {
				fmt.Printf("\"%s\"", word)
				if j < len(group)-1 {
					fmt.Printf(",")
				}
			}
			fmt.Printf("}")
			if i < len(c.Output)-1 {
				fmt.Printf(",")
			}
		}
		fmt.Printf("})\n")
	}
}
