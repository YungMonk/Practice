package train

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"regexp"
	"runtime"
	"strings"

	xmlpath "gopkg.in/xmlpath.v2"
)

// XpathParserEngine is a parser engine
func XpathParserEngine() {

	_, filename, _, _ := runtime.Caller(0)
	currentDir := path.Dir(filename)

	tmplate := path.Join(currentDir, "./xpathParser/carjob.html")
	file, err := os.OpenFile(tmplate, os.O_RDWR, os.ModePerm)
	if err != nil {
		panic("openFile failed!!!")
	}

	defer file.Close()

	htmlNode, err := xmlpath.ParseHTML(file)
	if err != nil {
		panic("xmlpath parse file failed!!! ")
	}

	configPath := path.Join(currentDir, "./xpathParser/config.json")
	parCfg := SetConifg(configPath)

	result := ParserFileds(parCfg, htmlNode)

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonBytes))
}

// CallBack is call back function.
type CallBack struct {
	Method string   `json:"method"`
	Params []string `json:"params"`
}

// ParserHead is the parser config
type ParserHead struct {
	Fileds []*ParserConfig
}

// Childs is the parser result
type Childs map[string]interface{}

// ParserConfig is the tempate config.
type ParserConfig struct {
	Filed   string          `json:"filed"`
	Rules   string          `json:"rules"`
	Lists   bool            `json:"lists"`
	Default interface{}     `json:"default"`
	Child   []*ParserConfig `json:"child"`
	Cback   []*CallBack     `json:"cback"`
}

// ParserFileds is parse index
func ParserFileds(p *ParserHead, node *xmlpath.Node) []Childs {
	var resulst []Childs
	for _, parcfg := range p.Fileds {
		key, val := Parser(parcfg, node)

		resulst = append(resulst, Childs{key: val})
	}

	return resulst
}

// Parser is detail
func Parser(p *ParserConfig, node *xmlpath.Node) (string, interface{}) {
	if p.Default != nil {
		return p.Filed, p.Default
	}

	// bookstore为根节点编译过后得到一个*Path类型的值 //*[@id="resultList"]/div[7]
	path := xmlpath.MustCompile(p.Rules)
	var val interface{}
	if len(p.Child) != 0 {
		nodes := []*xmlpath.Node{}

		// 可能会有多本书所以使用path.Iter(node)获取该节点下面的node集合也就是iterator
		it := path.Iter(node)

		// 判断是否有下一个
		for it.Next() {
			nodes = append(nodes, it.Node())
		}

		if p.Lists {
			childs := []Childs{}
			for _, chilNode := range nodes {
				content := Childs{}
				for _, filed := range p.Child {
					key, val := Parser(filed, chilNode)
					content[key] = val
				}
				childs = append(childs, content)
			}
			val = childs
		} else {
			content := Childs{}
			for _, filed := range p.Child {
				key, val := Parser(filed, nodes[0])
				content[key] = val
			}
			val = content
		}
	} else {
		nodeString, _ := path.String(node)
		val = strings.TrimSpace(nodeString)

		if len(p.Cback) != 0 {
			for _, callback := range p.Cback {
				val = HelpereFunc(callback.Method)(val, callback.Params)
			}
		}
	}

	return p.Filed, val
}

// HelpereFunc is the function process.
func HelpereFunc(callback string) func(args ...interface{}) interface{} {
	switch callback {
	case "callback1":
		return func(args ...interface{}) interface{} {
			reg, _ := regexp.Compile("\\d{4}-\\d*-\\d*")
			return reg.FindString(args[0].(string))
		}
	default:
		return func(args ...interface{}) interface{} {
			return args[0]
		}
	}
}

// SetConifg is set the template config.
func SetConifg(filename string) *ParserHead {
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

	var parHed ParserHead
	err = json.Unmarshal(fileData, &parHed.Fileds)
	if err != nil {
		fmt.Printf("json export struct faild: %+v \n", err)
	}

	return &parHed
}
