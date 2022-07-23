package main

import (
	"github.com/chrwhy/open-pinyin/parser"
	"github.com/chrwhy/open-pinyin/util"
	"log"
	"os"
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
	Test("yueguang")
	Test("anzn")
}
