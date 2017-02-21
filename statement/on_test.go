package statement_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/statement"
)

func TestOn(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement: statement.NewOn(
				statement.NewColumn("f.id"),
				statement.NewColumn("b.id"),
			),
			sql: "ON f.id = b.id",
			sqlIndent: `> ON f.id = b.id
`,
			values: nil,
		},
		{
			statement: statement.NewOn(
				statement.NewColumn("f.id"),
				statement.NewColumn("b.id"),
			).Join(
				statement.NewTable("bar"),
			),
			sql: "ON f.id = b.id JOIN bar",
			sqlIndent: `> ON f.id = b.id
> JOIN bar
`,
			values: nil,
		},
		{
			statement: statement.NewOn(
				statement.NewColumn("f.id"),
				statement.NewColumn("b.id"),
			).InnerJoin(
				statement.NewTable("bar"),
			),
			sql: "ON f.id = b.id INNER JOIN bar",
			sqlIndent: `> ON f.id = b.id
> INNER JOIN bar
`,
			values: nil,
		},
		{
			statement: statement.NewOn(
				statement.NewColumn("f.id"),
				statement.NewColumn("b.id"),
			).LeftJoin(
				statement.NewTable("bar"),
			),
			sql: "ON f.id = b.id LEFT JOIN bar",
			sqlIndent: `> ON f.id = b.id
> LEFT JOIN bar
`,
			values: nil,
		},
		{
			statement: statement.NewOn(
				statement.NewColumn("f.id"),
				statement.NewColumn("b.id"),
			).RightJoin(
				statement.NewTable("bar"),
			),
			sql: "ON f.id = b.id RIGHT JOIN bar",
			sqlIndent: `> ON f.id = b.id
> RIGHT JOIN bar
`,
			values: nil,
		},
	} {
		t.Run(fmt.Sprintf("%d Build", i), func(t *testing.T) {
			sql, values := b.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			sql, values := bi.Build(c.statement)
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}
