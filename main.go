package main

import (
	_ "github.com/lib/pq"
	"go_postgresql/GROUP_BY"
	"go_postgresql/JOINS"
	"go_postgresql/SELECT"
	"go_postgresql/assessmentTest"
)

func selectStatement() {
	//SELECT.Select()
	//SELECT.ChallengeSelect()
	//SELECT.ChallengeDistinct()
	//SELECT.Count()
	//SELECT.ChallengeWhere()
	//SELECT.OrderByLimit()
	//SELECT.ChallengeOrderBy()
	//SELECT.Between()
	//SELECT.In()
	//SELECT.LikeILike()
	SELECT.FinalChallenge()
}

func groupBy() {
	//GROUP_BY.Aggregation()
	//GROUP_BY.GroupBy()
	//GROUP_BY.ChallengeGroupBy()
	//GROUP_BY.Having()
	GROUP_BY.ChallengeHaving()
}

func test() {
	assessmentTest.AssessmentTest1()
}

func joins() {
	//JOINS.As()
	//JOINS.InnerJoin()
	//JOINS.FullOuterJoin()
	JOINS.LeftOuterJoin()
}

// main関数を実行するには実行構成の編集でディレクトリを選択し、goファイルのパッケージ名をmainにし、main関数を作成しないといけない
func main() {
	//selectStatement()
	//groupBy()
	//test()
	joins()
}
