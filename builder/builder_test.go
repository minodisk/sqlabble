package builder_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/builder"
	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/statement"
	"github.com/minodisk/sqlabble/token"
)

var (
	b  = builder.Standard
	bi = builder.NewBuilder(token.NewFormat("> ", "  ", `"`, "\n"))
)

func TestBuilder(t *testing.T) {
	for i, c := range []struct {
		statement statement.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			statement.NewSelect(
				statement.NewColumn("created_at"),
				statement.NewColumn("name").As("n"),
				statement.NewColumn("gender").As("g"),
				statement.NewColumn("age"),
			).From(
				statement.NewTable("users"),
			).Where(
				statement.NewAnd(
					statement.NewColumn("g").Eq(statement.NewParam("male")),
					statement.NewOr(
						statement.NewColumn("age").Lt(statement.NewParam(20)),
						statement.NewColumn("age").Eq(statement.NewParam(30)),
						statement.NewColumn("age").Gte(statement.NewParam(50)),
					),
					statement.NewColumn("created_at").Between(statement.NewParam("2016-01-01"), statement.NewParam("2016-12-31")),
				),
			).OrderBy(
				statement.NewColumn("created_at").Desc(),
				statement.NewColumn("id").Asc(),
			).Limit(
				20,
			).Offset(
				20 * 5,
			),
			`SELECT created_at, name AS "n", gender AS "g", age FROM users WHERE g = ? AND (age < ? OR age = ? OR age >= ?) AND created_at BETWEEN ? AND ? ORDER BY created_at DESC, id ASC LIMIT ? OFFSET ?`,
			`> SELECT
>   created_at
>   , name AS "n"
>   , gender AS "g"
>   , age
> FROM
>   users
> WHERE
>   g = ?
>   AND (
>     age < ?
>     OR age = ?
>     OR age >= ?
>   )
>   AND created_at BETWEEN ? AND ?
> ORDER BY
>   created_at DESC
>   , id ASC
> LIMIT
>   ?
> OFFSET
>   ?
`,
			[]interface{}{
				`male`,
				20,
				30,
				50,
				`2016-01-01`,
				`2016-12-31`,
				20,
				100,
			},
		},
		{
			statement.NewInsertInto(
				statement.NewTable("foo"),
				statement.NewColumn("name"),
				statement.NewColumn("age"),
			).Values(
				statement.NewParams(`Obi-Wan Kenobi`, 63),
				statement.NewParams(`Luke Skywalker`, 19),
			),
			`INSERT INTO foo (name, age) VALUES (?, ?), (?, ?)`,
			`> INSERT INTO
>   foo (
>     name
>     , age
>   )
> VALUES
>   (?, ?)
>   , (?, ?)
`,
			[]interface{}{
				`Obi-Wan Kenobi`,
				63,
				`Luke Skywalker`,
				19,
			},
		},
		{
			statement.NewDelete().From(
				statement.NewTable("login_history"),
			).Where(
				statement.NewColumn("login_date").Lt(statement.NewParam("2004-07-02 09:00:00")),
			),
			`DELETE FROM login_history WHERE login_date < ?`,
			`> DELETE
> FROM
>   login_history
> WHERE
>   login_date < ?
`,
			[]interface{}{
				`2004-07-02 09:00:00`,
			},
		},
		{
			statement.NewUnion(
				statement.NewSelect(
					statement.NewColumn("emp_id"),
				).From(
					statement.NewTable("employee"),
				).Where(
					statement.NewAnd(
						statement.NewColumn("assigned_branch_id").Eq(statement.NewParam(2)),
						statement.NewOr(
							statement.NewColumn("title").Eq(statement.NewParam("Teller")),
							statement.NewColumn("title").Eq(statement.NewParam("Head Teller")),
						),
					),
				),
				statement.NewSelectDistinct(
					statement.NewColumn("open_emp_id"),
				).From(
					statement.NewTable("account"),
				).Where(
					statement.NewColumn("open_branch_id").Eq(statement.NewParam(2)),
				),
			),
			`(SELECT emp_id FROM employee WHERE assigned_branch_id = ? AND (title = ? OR title = ?)) UNION (SELECT DISTINCT open_emp_id FROM account WHERE open_branch_id = ?)`,
			`> (
>   SELECT
>     emp_id
>   FROM
>     employee
>   WHERE
>     assigned_branch_id = ?
>     AND (
>       title = ?
>       OR title = ?
>     )
> )
> UNION (
>   SELECT DISTINCT
>     open_emp_id
>   FROM
>     account
>   WHERE
>     open_branch_id = ?
> )
`,
			[]interface{}{
				2,
				`Teller`,
				`Head Teller`,
				2,
			},
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
