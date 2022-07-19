package main

import (
	"log"
	"pinyin/parser"
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

	for _, input := range inputs {
		result := parser.GreedyParse(input)
		log.Println(result)
	}

}

func main() {
	//TestCase()
	parser.Parse("zhangqiange")
	//parser.Parse("lu")
	//parser.Elect("xian", "")
}
