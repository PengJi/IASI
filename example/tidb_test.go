package example

import (
	"testing"
)

func TestPrintStmtNode(t *testing.T) {
	sqls := []string {
		`select * from tbl where id = 1`,
		`select * f`,
	}

	for _, sql := range sqls {
		PrintStmtNode(sql, "", "")
	}
}
/*
func TestStmtNode2JSON(t *testing.T) {
	sqls := []string {
		`select * from tbl where id = 1`,
		`select * f`,
	}
	for _, sql := range sqls {
		fmt.Println(StmtNode2JSON(sql, "", ""))
	}
}
*/

