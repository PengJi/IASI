package example

import (
	"testing"
)

func TestPrintStmtNode(t *testing.T) {
	sqls := []string {
		`select * from tbl where id = 1`,
		`select tb1.name, tb2.name from tb1 
		left join tb2 on tb1.id=tb2.ref_id where tb1.a='test' `,
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

