package parser

import (
	"fmt"
	"log"
	"testing"
)



/*
func TestPrintStmtNode(t *testing.T) {
	sqls := []string {
		`select tb1.name, tb2.name from tb1 
		left join tb2 on tb1.id=tb2.ref_id where tb1.a='test'`,
	}

	for _, sql := range sqls {
		PrintStmtNode(sql, "", "")
	}
}
*/

func TestStmtNode2JSON(t *testing.T) {
	sqls := []string {
		`select tb1.name, tb2.name from tb1 
		left join tb2 on tb1.id=tb2.ref_id where tb1.a='test'`,
	}
	for _, sql := range sqls {
		fmt.Println(StmtNode2JSON(sql, "", ""))
	}
}

func TestSomething(t *testing.T){
	sqls := []string{
		`select tb1.name, tb2.name from tb1 
		left join tb2 on tb1.id=tb2.ref_id where tb1.a='test'`,
	}

	str := StmtNode2JSON(sqls[0], "","")
	fmt.Println(len(str))
	fmt.Printf("%T\n",str)
}

func TestPredicatePushDown(t *testing.T) {
	sqls := []string{
		"select count(*) from t a, t b where a.a = b.a",
		"select a from (select a from t where d = 0) k where k.a = 5",
		"select a from (select a+1 as a from t) k where k.a = 5",
		"select a from (select 1+2 as a from t where d = 0) k where k.a = 5",
		"select a from (select d as a from t where d = 0) k where k.a = 5",
		"select * from t ta, t tb where (ta.d, ta.a) = (tb.b, tb.c)",
		"select * from t t1, t t2 where t1.a = t2.b and t2.b > 0 and t1.a = t1.c and t1.d like 'abc' and t2.d = t1.d",
		"select * from t ta join t tb on ta.d = tb.d and ta.d > 1 where tb.a = 0",
		"select * from t ta join t tb on ta.d = tb.d where ta.d > 1 and tb.a = 0",
		"select * from t ta left outer join t tb on ta.d = tb.d and ta.d > 1 where tb.a = 0",
		"select * from t ta right outer join t tb on ta.d = tb.d and ta.a > 1 where tb.a = 0",
		"select * from t ta left outer join t tb on ta.d = tb.d and ta.a > 1 where ta.d = 0",
		"select * from t ta left outer join t tb on ta.d = tb.d and ta.a > 1 where tb.d = 0",
		"select * from t ta left outer join t tb on ta.d = tb.d and ta.a > 1 where tb.c is not null and tb.c = 0 and ifnull(tb.d, 1)",
		"select * from t ta left outer join t tb on ta.a = tb.a left outer join t tc on tb.b = tc.b where tc.c > 0",
		"select * from t ta left outer join t tb on ta.a = tb.a left outer join t tc on tc.b = ta.b where tb.c > 0",
		"select * from t as ta left outer join (t as tb left join t as tc on tc.b = tb.b) on tb.a = ta.a where tc.c > 0",
		"select * from ( t as ta left outer join t as tb on ta.a = tb.a) join ( t as tc left join t as td on tc.b = td.b) on ta.c = td.c where tb.c = 2 and td.a = 1",
		"select * from t ta left outer join (t tb left outer join t tc on tc.b = tb.b) on tb.a = ta.a and tc.c = ta.c where tc.d > 0 or ta.d > 0",
		"select * from t ta left outer join t tb on ta.d = tb.d and ta.a > 1 where ifnull(tb.d, 1) or tb.d is null",
		"select a, d from (select * from t union all select * from t union all select * from t) z where a < 10",
		"select (select count(*) from t where t.a = k.a) from t k",
		"select a from t where exists(select 1 from t as x where x.a < t.a)",
		"select a from t where exists(select 1 from t as x where x.a = t.a and t.a < 1 and x.a < 1)",
		"select a from t where exists(select 1 from t as x where x.a = t.a and x.a < 1) and a < 1",
		"select a from t where exists(select 1 from t as x where x.a = t.a) and exists(select 1 from t as x where x.a = t.a)",
		"select * from (select a, b, sum(c) as s from t group by a, b) k where k.a > k.b * 2 + 1",
		"select * from (select a, b, sum(c) as s from t group by a, b) k where k.a > 1 and k.b > 2",
		"select * from (select k.a, sum(k.s) as ss from (select a, sum(b) as s from t group by a) k group by k.a) l where l.a > 2",
		"select * from (select a, sum(b) as s from t group by a) k where a > s",
		"select * from (select a, sum(b) as s from t group by a + 1) k where a > 1",
		"select * from (select a, sum(b) as s from t group by a having 1 = 0) k where a > 1",
		"select a, count(a) cnt from t group by a having cnt < 1",
		"select t1.a, t2.a from t as t1 left join t as t2 on t1.a = t2.a where t1.a < 1.0",
		"select * from t t1 join t t2 on t1.a = t2.a where t2.a = null",
	}

	//_ := parser.New()

	for _, sql := range sqls {
		fmt.Println(sql)
	}
}

func TestJoinPredicatePushDown(t *testing.T){
	sqls := []string{
		// inner join
		"select * from t as t1 join t as t2 on t1.b = t2.b where t1.a > t2.a",
		"select * from t as t1 join t as t2 on t1.b = t2.b where t1.a=1 or t2.a=1",
		"select * from t as t1 join t as t2 on t1.b = t2.b where (t1.a=1 and t2.a=1) or (t1.a=2 and t2.a=2)",
		"select * from t as t1 join t as t2 on t1.b = t2.b where (t1.c=1 and (t1.a=3 or t2.a=3)) or (t1.a=2 and t2.a=2)",
		"select * from t as t1 join t as t2 on t1.b = t2.b where (t1.c=1 and ((t1.a=3 and t2.a=3) or (t1.a=4 and t2.a=4)))",
		"select * from t as t1 join t as t2 on t1.b = t2.b where (t1.a>1 and t1.a < 3 and t2.a=1) or (t1.a=2 and t2.a=2)",
		"select * from t as t1 join t as t2 on t1.b = t2.b and ((t1.a=1 and t2.a=1) or (t1.a=2 and t2.a=2))",
		// left join
		"select * from t as t1 left join t as t2 on t1.b = t2.b and ((t1.a=1 and t2.a=1) or (t1.a=2 and t2.a=2))",
		"select * from t as t1 left join t as t2 on t1.b = t2.b and t1.a > t2.a",
		"select * from t as t1 left join t as t2 on t1.b = t2.b and (t1.a=1 or t2.a=1)",
		"select * from t as t1 left join t as t2 on t1.b = t2.b and ((t1.c=1 and (t1.a=3 or t2.a=3)) or (t1.a=2 and t2.a=2))",
		"select * from t as t1 left join t as t2 on t1.b = t2.b and ((t2.c=1 and (t1.a=3 or t2.a=3)) or (t1.a=2 and t2.a=2))",
		"select * from t as t1 left join t as t2 on t1.b = t2.b and ((t1.c=1 and ((t1.a=3 and t2.a=3) or (t1.a=4 and t2.a=4))) or (t1.a=2 and t2.a=2))",
		"select * from t t1 join t t2 on t1.a > 1 and t1.a > 1",
	}

	for _, sql := range sqls {
		fmt.Println(sql)
	}

}

func TestOuterWherePredicatePushDown(t *testing.T) {
	sqls := []string {
		// left join with where condition
		"select * from t as t1 left join t as t2 on t1.b = t2.b where (t1.a=1 and t2.a is null) or (t1.a=2 and t2.a=2)",
		"select * from t as t1 left join t as t2 on t1.b = t2.b where (t1.c=1 and (t1.a=3 or t2.a=3)) or (t1.a=2 and t2.a=2)",
		"select * from t as t1 left join t as t2 on t1.b = t2.b where (t1.c=1 and ((t1.a=3 and t2.a=3) or (t1.a=4 and t2.a=4))) or (t1.a=2 and t2.a is null)",
	}

	for _, sql := range sqls {
		fmt.Println(sql)
	}
}

func TestSimplifyOuterJoin(t *testing.T) {
	sqls := []string {
		"select * from t t1 left join t t2 on t1.b = t2.b where t1.c > 1 or t2.c > 1;",
		"select * from t t1 left join t t2 on t1.b = t2.b where t1.c > 1 and t2.c > 1;",
		"select * from t t1 left join t t2 on t1.b = t2.b where not (t1.c > 1 or t2.c > 1);",
		"select * from t t1 left join t t2 on t1.b = t2.b where not (t1.c > 1 and t2.c > 1);",
		"select * from t t1 left join t t2 on t1.b > 1 where t1.c = t2.c;",
		"select * from t t1 left join t t2 on true where t1.b <=> t2.b;",
		// semi join
		"select a from t t1 where not exists (select a from t t2 where t1.a = t2.a and t2.b = 1 and t2.b = 2)",
	}

	for _, sql := range sqls {
		fmt.Println(sql)
	}
}

func TestDeriveNotNullConds(t *testing.T) {
	sqls := []string {
		"select * from t t1 inner join t t2 on t1.e = t2.e",
		"select * from t t1 inner join t t2 on t1.e > t2.e",
		"select * from t t1 inner join t t2 on t1.e = t2.e and t1.e is not null",
		"select * from t t1 left join t t2 on t1.e = t2.e",
		"select * from t t1 left join t t2 on t1.e > t2.e",
		"select * from t t1 left join t t2 on t1.e = t2.e and t2.e is not null",
		"select * from t t1 right join t t2 on t1.e = t2.e and t1.e is not null",
		"select * from t t1 inner join t t2 on t1.e <=> t2.e",
		"select * from t t1 left join t t2 on t1.e <=> t2.e",
		// Not deriving if column has NotNull flag already.
		"select * from t t1 inner join t t2 on t1.b = t2.b",
		"select * from t t1 left join t t2 on t1.b = t2.b",
		"select * from t t1 left join t t2 on t1.b > t2.b",
		// Not deriving for AntiSemiJoin
		"select * from t t1 where not exists (select * from t t2 where t2.e = t1.e)",
	}

	for _, sql := range sqls {
		fmt.Println(sql)
	}
}

func TestSubquery(t *testing.T) {
	sqls := []string {
		// This will be resolved as in sub query.
		"select * from t where 10 in (select b from t s where s.a = t.a)",
		"select count(c) ,(select b from t s where s.a = t.a) from t",
		"select count(c) ,(select count(s.b) from t s where s.a = t.a) from t",
		// Semi-join with agg cannot decorrelate.
		"select t.c in (select count(s.b) from t s where s.a = t.a) from t",
		"select (select count(s.b) k from t s where s.a = t.a having k != 0) from t",
		"select (select count(s.b) k from t s where s.a = t1.a) from t t1, t t2",
		"select (select count(1) k from t s where s.a = t.a having k != 0) from t",
		"select a from t where a in (select a from t s group by t.b)",
		// This will be resolved as in sub query.
		"select * from t where 10 in (((select b from t s where s.a = t.a)))",
		// This will be resolved as in function.
		"select * from t where 10 in (((select b from t s where s.a = t.a)), 10)",
		"select * from t where exists (select s.a from t s having sum(s.a) = t.a )",
		// Test MaxOneRow for limit.
		"select (select * from (select b from t limit 1) x where x.b = t1.b) from t t1",
		// Test Nested sub query.
		"select * from t where exists (select s.a from t s where s.c in (select c from t as k where k.d = s.d) having sum(s.a) = t.a )",
		"select t1.b from t t1 where t1.b = (select max(t2.a) from t t2 where t1.b=t2.b)",
		"select t1.b from t t1 where t1.b = (select avg(t2.a) from t t2 where t1.g=t2.g and (t1.b = 4 or t2.b = 2))",
	}

	for _, sql := range sqls {
		log.Println(sql)
	}
}

func TestJoinReOrder(t *testing.T) {
	sqls := []string{
		"select * from t t1, t t2, t t3, t t4, t t5, t t6 where t1.a = t2.b and t2.a = t3.b and t3.c = t4.a and t4.d = t2.c and t5.d = t6.d",
		"select * from t t1, t t2, t t3, t t4, t t5, t t6, t t7, t t8 where t1.a = t8.a",
		"select * from t t1, t t2, t t3, t t4, t t5 where t1.a = t5.a and t5.a = t4.a and t4.a = t3.a and t3.a = t2.a and t2.a = t1.a and t1.a = t3.a and t2.a = t4.a and t5.b < 8",
		"select * from t t1, t t2, t t3, t t4, t t5 where t1.a = t5.a and t5.a = t4.a and t4.a = t3.a and t3.a = t2.a and t2.a = t1.a and t1.a = t3.a and t2.a = t4.a and t3.b = 1 and t4.a = 1",
		"select * from t o where o.b in (select t3.c from t t1, t t2, t t3 where t1.a = t3.a and t2.a = t3.a and t2.a = o.a)",
		"select * from t o where o.b in (select t3.c from t t1, t t2, t t3 where t1.a = t3.a and t2.a = t3.a and t2.a = o.a and t1.a = 1)",
	}

	for _, sql := range sqls {
		log.Println(sql)
	}
}

func TestEagerAggregation(t *testing.T) {
	sqls := []string {
		"select sum(t.a), sum(t.a+1), sum(t.a), count(t.a), sum(t.a) + count(t.a) from t",
		"select sum(t.a + t.b), sum(t.a + t.c), sum(t.a + t.b), count(t.a) from t having sum(t.a + t.b) > 0 order by sum(t.a + t.c)",
		"select sum(a.a) from t a, t b where a.c = b.c",
		"select sum(b.a) from t a, t b where a.c = b.c",
		"select sum(b.a), a.a from t a, t b where a.c = b.c",
		"select sum(a.a), b.a from t a, t b where a.c = b.c",
		"select sum(a.a), sum(b.a) from t a, t b where a.c = b.c",
		"select sum(a.a), max(b.a) from t a, t b where a.c = b.c",
		"select max(a.a), sum(b.a) from t a, t b where a.c = b.c",
		"select sum(a.a) from t a, t b, t c where a.c = b.c and b.c = c.c",
		"select sum(b.a) from t a left join t b on a.c = b.c",
		"select sum(a.a) from t a left join t b on a.c = b.c",
		"select sum(a.a) from t a right join t b on a.c = b.c",
		"select sum(a) from (select * from t) x",
		"select sum(c1) from (select c c1, d c2 from t a union all select a c1, b c2 from t b union all select b c1, e c2 from t c) x group by c2",
		"select max(a.b), max(b.b) from t a join t b on a.c = b.c group by a.a",
		"select max(a.b), max(b.b) from t a join t b on a.a = b.a group by a.c",
		"select max(c.b) from (select * from t a union all select * from t b) c group by c.a",
		"select max(a.c) from t a join t b on a.a=b.a and a.b=b.b group by a.b",
		"select t1.a, count(t2.b) from t t1, t t2 where t1.a = t2.a group by t1.a",
	}

	for _, sql := range sqls {
		log.Println(sql)
	}
}

func TestColumnPruning(t *testing.T) {
	sqls := []string {
		"select count(*) from t group by a",
		"select count(*) from t",
		"select count(*) from t a join t b where a.a < 1",
		"select count(*) from t a join t b on a.a = b.d",
		"select count(*) from t a join t b on a.a = b.d order by sum(a.d)",
		"select count(b.a) from t a join t b on a.a = b.d group by b.b order by sum(a.d)",
		"select * from (select count(b.a) from t a join t b on a.a = b.d group by b.b having sum(a.d) < 0) tt",
		"select (select count(a) from t where b = k.a) from t k",
		"select exists (select count(*) from t where b = k.a) from t k",
		"select b = (select count(*) from t where b = k.a) from t k",
		"select exists (select count(a) from t where b = k.a group by b) from t k",
		"select a as c1, b as c2 from t order by 1, c1 + c2 + c",
		"select a from t where b < any (select c from t)",
		"select a from t where (b,a) != all (select c,d from t)",
		"select a from t where (b,a) in (select c,d from t)",
		"select a from t where a in (select a from t s group by t.b)",
	}

	for _, sql := range sqls {
		log.Println(sql)
	}
}

func TestAggPrune(t *testing.T) {
	sqls := []string {
		"select a, count(b) from t group by a",
		"select sum(b) from t group by c, d, e",
		"select tt.a, sum(tt.b) from (select a, b from t) tt group by tt.a",
		"select count(1) from (select count(1), a as b from t group by a) tt group by b",
		"select a, count(b) from t group by a",
		"select a, count(distinct a, b) from t group by a",
	}

	for _, sql := range sqls {
		log.Println(sql)
	}
}

func TestUnion(t *testing.T) {
	sqls := []string {
		"select a from t union select a from t",
		"select a from t union all select a from t",
		"select a from t union select a from t union all select a from t",
		"select a from t union select a from t union all select a from t union select a from t union select a from t",
		"select a from t union select a, b from t",
		"select * from (select 1 as a  union select 1 union all select 2) t order by a",
		"select * from (select 1 as a  union select 1 union all select 2) t order by (select a)",
	}

	for _, sql := range sqls {
		log.Println(sql)
	}
}

func TestTopNPushDown(t *testing.T) {
	sqls := []string {
		// Test TopN + Selection.
		"select * from t where a < 1 order by b limit 5",
		// Test Limit + Selection.
		"select * from t where a < 1 limit 5",
		// Test Limit + Agg + Proj .
		"select a, count(b) from t group by b limit 5",
		// Test TopN + Agg + Proj .
		"select a, count(b) from t group by b order by c limit 5",
		// Test TopN + Join + Proj.
		"select * from t, t s order by t.a limit 5",
		// Test Limit + Join + Proj.
		"select * from t, t s limit 5",
		// Test TopN + Left Join + Proj.
		"select * from t left outer join t s on t.a = s.a order by t.a limit 5",
		// Test TopN + Left Join + Proj.
		"select * from t left outer join t s on t.a = s.a order by t.a limit 5, 5",
		// Test Limit + Left Join + Proj.
		"select * from t left outer join t s on t.a = s.a limit 5",
		// Test Limit + Left Join Apply + Proj.
		"select (select s.a from t s where t.a = s.a) from t limit 5",
		// Test TopN + Left Join Apply + Proj.
		"select (select s.a from t s where t.a = s.a) from t order by t.a limit 5",
		// Test TopN + Left Semi Join Apply + Proj.
		"select exists (select s.a from t s where t.a = s.a) from t order by t.a limit 5",
		// Test TopN + Semi Join Apply + Proj.
		"select * from t where exists (select s.a from t s where t.a = s.a) order by t.a limit 5",
		// Test TopN + Right Join + Proj.
		"select * from t right outer join t s on t.a = s.a order by s.a limit 5",
		// Test Limit + Right Join + Proj.
		"select * from t right outer join t s on t.a = s.a order by s.a,t.b limit 5",
		// Test TopN + UA + Proj.
		"select * from t union all (select * from t s) order by a,b limit 5",
		// Test TopN + UA + Proj.
		"select * from t union all (select * from t s) order by a,b limit 5, 5",
		// Test Limit + UA + Proj + Sort.
		"select * from t union all (select * from t s order by a) limit 5",
		// Test `ByItem` containing column from both sides.
		"select ifnull(t1.b, t2.a) from t t1 left join t t2 on t1.e=t2.e order by ifnull(t1.b, t2.a) limit 5",
		// Test ifnull cannot be eliminated
		"select ifnull(t1.h, t2.b) from t t1 left join t t2 on t1.e=t2.e order by ifnull(t1.h, t2.b) limit 5",
	}

	for _, sql := range sqls {
		log.Println(sql)
	}
}

func TestOuterJoinEliminator(t *testing.T) {
	sqls := []string {
		// Test left outer join + distinct
		"select distinct t1.a, t1.b from t t1 left outer join t t2 on t1.b = t2.b",
		// Test right outer join + distinct
		"select distinct t2.a, t2.b from t t1 right outer join t t2 on t1.b = t2.b",
		// Test duplicate agnostic agg functions on join
		"select max(t1.a), min(t1.b) from t t1 left join t t2 on t1.b = t2.b",
		"select sum(distinct t1.a) from t t1 left join t t2 on t1.a = t2.a and t1.b = t2.b",
		"select count(distinct t1.a, t1.b) from t t1 left join t t2 on t1.b = t2.b",
		// Test left outer join
		"select t1.b from t t1 left outer join t t2 on t1.a = t2.a",
		// Test right outer join
		"select t2.b from t t1 right outer join t t2 on t1.a = t2.a",
		// For complex join query
		"select max(t3.b) from (t t1 left join t t2 on t1.a = t2.a) right join t t3 on t1.b = t3.b",
	}

	for _, sql := range sqls {
		log.Println(sql)
	}
}

func TestWindowFunction(t *testing.T) {
	sqls := []string {
		"select a, avg(a) over(partition by a) from t",
		"select a, avg(a) over(partition by b) from t",
		"select a, avg(a+1) over(partition by (a+1)) from t",
		"select a, avg(a) over(order by a asc, b desc) from t order by a asc, b desc",
		"select a, b as a, avg(a) over(partition by a) from t",
		"select a, b as z, sum(z) over() from t",
		"select a, b as z from t order by (sum(z) over())",
		"select sum(avg(a)) over() from t",
		"select b from t order by(sum(a) over())",
		"select b from t order by(sum(a) over(partition by a))",
		"select b from t order by(sum(avg(a)) over())",
		"select a from t having (select sum(a) over() as w from t tt where a > t.a)",
		"select avg(a) over() as w from t having w > 1",
		"select sum(a) over() as sum_a from t group by sum_a",
		"select sum(a) over() from t window w1 as (w2)",
		"select sum(a) over(w) from t",
		"select sum(a) over() from t window w1 as (w2), w2 as (w1)",
		"select sum(a) over(w partition by a) from t window w as ()",
		"select sum(a) over(w) from t window w as (rows between 1 preceding AND 1 following)",
		"select sum(a) over w from t window w as (rows between 1 preceding AND 1 following)",
		"select sum(a) over(w order by b) from t window w as (order by a)",
		"select sum(a) over() from t window w1 as (), w1 as ()",
		"select sum(a) over(w1), avg(a) over(w2) from t window w1 as (partition by a), w2 as (w1)",
		"select a from t window w1 as (partition by a) order by (sum(a) over(w1))",
		"select sum(a) over(groups 1 preceding) from t",
		"select sum(a) over(rows between unbounded following and 1 preceding) from t",
		"select sum(a) over(rows between current row and unbounded preceding) from t",
		"select sum(a) over(rows interval 1 MINUTE_SECOND preceding) from t",
		"select sum(a) over(rows between 1.0 preceding and 1 following) from t",
		"select sum(a) over(range between 1 preceding and 1 following) from t",
		"select sum(a) over(order by c_str range between 1 preceding and 1 following) from t",
		"select sum(a) over(order by a range interval 1 MINUTE_SECOND preceding) from t",
		"select sum(a) over(order by i_date range interval a MINUTE_SECOND preceding) from t",
		"select sum(a) over(order by i_date range interval -1 MINUTE_SECOND preceding) from t",
		"select sum(a) over(order by i_date range 1 preceding) from t",
		"select sum(a) over(order by a range between 1.0 preceding and 1 following) from t",
		"select row_number() over(rows between 1 preceding and 1 following) from t",
		"select avg(b), max(avg(b)) over(rows between 1 preceding and 1 following) max, min(avg(b)) over(rows between 1 preceding and 1 following) min from t group by c",
		"select nth_value(a, 1.0) over() from t",
		"select nth_value(a, 0) over() from t",
		"select ntile(0) over() from t",
		"select ntile(null) over() from t",
		"select avg(a) over w from t window w as(partition by b)",
		"select nth_value(i_date, 1) over() from t",
	}

	for _, sql := range sqls {
		log.Println(sql)
	}
}



