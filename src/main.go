package main

import (
	"fmt"
	"os"

	xmlpath "gopkg.in/xmlpath.v2"
)

var file *os.File
var node *xmlpath.Node
var htmlNode *xmlpath.Node

func init() {
	var err error
	file, err = os.OpenFile("t.xml", os.O_RDWR, os.ModePerm)
	if err != nil {
		panic("openFile failed!!!")
	}

	defer file.Close()

	// 解析文件获得经过处理的并且可以被path访问的Node类型的node
	node, err = xmlpath.Parse(file)
	if err != nil {
		panic("xmlpath parse file failed!!! ")
	}

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
	nodename()
	// getFieldValue()
	// getSpecifiedValue()
}

func nodename() {
	// bookstore为根节点编译过后得到一个*Path类型的值 //*[@id="resultList"]/div[7]
	path := xmlpath.MustCompile("//*[@id=\"resultList\"]/div[@class=\"el\"]")

	// 可能会有多本书所以使用path.Iter(node)获取该节点下面的node集合也就是iterator
	it := path.Iter(htmlNode)

	// 判断是否有下一个
	for it.Next() {
		// 如果有把当前的Node节点取出 并打印出值
		fmt.Println(it.Node().String())
	}
}

func getFieldValue() {
	// 选取全文中属性包含@lang的节点 而不管他的位置
	path := xmlpath.MustCompile("//@lang")
	it := path.Iter(node)
	for it.Next() {
		fmt.Println(it.Node().String())
	}

	// 选取全文中属性包含lang并且值为en的节点 而不管他的位置
	path = xmlpath.MustCompile("//*[@lang=\"en\"]")
	it = path.Iter(node)
	for it.Next() {
		fmt.Println(it.Node().String())
	}
}

func getSpecifiedValue() {
	// 选取Bookstore中的第二本书的title节点
	path := xmlpath.MustCompile("/bookstore/book[2]/title")
	fmt.Println(path.String(node))
}
