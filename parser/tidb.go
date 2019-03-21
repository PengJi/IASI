package parser

import (
	"encoding/json"
	"fmt"

	"github.com/kr/pretty"
	_ "github.com/pingcap/tidb/types/parser_driver"

	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"
)

// 语法解析
func TiParse(sql, charset, collation string) ([]ast.StmtNode, error){
	p := parser.New()
	stmt, _, err := p.Parse(sql, charset, collation)
	return stmt, err
}

// 打印语法树
func PrintStmtNode(sql, charset, collation string){
	tree, err := TiParse(sql, charset, collation)
	if err != nil {
		fmt.Println(err)
	} else {
		_, err = pretty.Println(tree)
		fmt.Print(err)
	}
}

// 转化为json格式
func StmtNode2JSON(sql, charset, collation string) string {
	var str string
	tree, err := TiParse(sql, charset, collation)
	if err != nil {
		fmt.Println(err)
	} else {
		b, err := json.MarshalIndent(tree, "", "  ")
		if err != nil {
			fmt.Println(err)
		} else {
			str = string(b)
		}
	}

	return str
}





