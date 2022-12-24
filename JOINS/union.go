package JOINS

// UNION例
// TableA       TableB
// name amount, name amount
// David  100   David 200
// Claire 50    Claire 100
// というテーブルたちに対して
//
// SELECT * FROM TableA
// UNION
// SELECT * FROM TableB
//
// を実行すると
// name amount
// David 100
// Claire 50
// David 200
// Claire 100
//
// という結果が得られる
