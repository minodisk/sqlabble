package stmt_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/stmt"
)

func TestColumnAsType(t *testing.T) {
	t.Parallel()
	for _, c := range []interface{}{
		stmt.ColumnAlias{},
	} {
		c := c
		t.Run(fmt.Sprintf("%T", c), func(t *testing.T) {
			t.Parallel()
			if _, ok := c.(stmt.ColOrAliasOrSub); !ok {
				t.Errorf("%T should implement stmt.ColumnOrColumnAs", c)
			}
		})
	}
}

func TestColumnAsSQL(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewColumnAlias("foo"),
			`AS "foo"`,
			`> AS "foo"
`,
			nil,
		},
	} {
		c := c
		t.Run(fmt.Sprintf("%d Build", i), func(t *testing.T) {
			t.Parallel()
			sql, values := b.Build(c.stmt)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			t.Parallel()
			sql, values := bi.Build(c.stmt)
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}
