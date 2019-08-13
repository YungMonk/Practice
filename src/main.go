package main

import (
	"encoding/json"
	"fmt"
	"os"

	xmlpath "gopkg.in/xmlpath.v2"
)

var file *os.File
var htmlNode *xmlpath.Node

func init() {
	var err error

	file, err = os.OpenFile("carjob.html", os.O_RDWR, os.ModePerm)
	if err != nil {
		panic("openFile failed!!!")
	}

	defer file.Close()

	htmlNode, err = xmlpath.ParseHTML(file)
	if err != nil {
		panic("xmlpath parse file failed!!! ")
	}
}

func main() {
	parCfg := setConifg("/mnt/d/Development/workspace/src/config.json")
	ParserFileds(parCfg, htmlNode)
	// nodename()
}

// ParserConfig init
// config := parserConfig{
// 	filed: "position",
// 	rules: "//*[@id=\"resultList\"]/div[@class=\"el\"]",
// 	child: []*parserConfig{
// 		&parserConfig{filed: "name", rules: "./span[2]", child: nil},
// 	},
// }
type ParserConfig struct {
	Filed string          `json:"filed"`
	Rules string          `json:"rules"`
	Lists bool            `json:"lists"`
	Child []*ParserConfig `json:"child"`
}

// ParserHead is the parser config
type ParserHead struct {
	Fileds []*ParserConfig
}

// ParserFileds is parse index
func ParserFileds(p *ParserHead, node *xmlpath.Node) {
	// result := make(map[string]interface{})
	for _, parcfg := range p.Fileds {
		Parser(parcfg, node)
	}
}

// Parser is detail
func Parser(p *ParserConfig, node *xmlpath.Node) {
	// bookstore为根节点编译过后得到一个*Path类型的值 //*[@id="resultList"]/div[7]
	path := xmlpath.MustCompile(p.Rules)

	if len(p.Child) != 0 {
		nodes := []*xmlpath.Node{}

		it := path.Iter(node)
		// 判断是否有下一个
		for it.Next() {
			nodes = append(nodes, it.Node())
		}

		if p.Lists {
			for _, chilNode := range nodes {
				for _, filed := range p.Child {
					Parser(filed, chilNode)
				}
			}
		} else {

			for _, filed := range p.Child {
				Parser(filed, nodes[0])
			}
		}

	} else {
		val, _ := path.String(node)
		fmt.Println(val)
	}
}

// filed, rules, child
func nodename() {

	// bookstore为根节点编译过后得到一个*Path类型的值 //*[@id="resultList"]/div[7]
	path := xmlpath.MustCompile("//div[@class='myResume_box']//span[contains(text(), '工作经历')]/ancestor::div[@class='myResume_show']/div[@class='showInfo clearfix']")

	// 可能会有多本书所以使用path.Iter(node)获取该节点下面的node集合也就是iterator
	it := path.Iter(htmlNode)

	nodes := []*xmlpath.Node{}
	// 判断是否有下一个
	for it.Next() {
		nodes = append(nodes, it.Node())
	}

	childPath := xmlpath.MustCompile(".//div[@class='myResume_title']")
	childStr, err := childPath.String(nodes[0])
	if !err {
		fmt.Printf("Parser rule is: %s\nResult: %s", "./span[1]", childStr)
	}

	fmt.Printf("%+v \n", childStr)
}

func setConifg(filename string) *ParserHead {
	// 获取文件指针
	fp, err := os.Open(filename)
	if err != nil {
		fmt.Println("config file is not exist.")
	}
	defer fp.Close()

	fileInfo, err := fp.Stat()
	if err != nil {
		fmt.Println("config file is read error.")
	}

	// 文件内容读取到buffer中
	fileData := make([]byte, fileInfo.Size())
	_, err = fp.Read(fileData)
	if err != nil {
		fmt.Println("config file is read faild.")
	}

	// fmt.Printf("%s \n", string(fileData))

	var parHed ParserHead

	err = json.Unmarshal(fileData, &parHed.Fileds)
	if err != nil {
		fmt.Printf("json export struct faild: %+v \n", err)
	}

	// fmt.Printf("%+v \n", parHed)

	return &parHed
}
