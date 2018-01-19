package stmt_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/sqlabble/sqlabble/internal/diff"
	"github.com/sqlabble/sqlabble/stmt"
)

func TestOrderType(t *testing.T) {
	t.Parallel()
	o := stmt.Order{}
	if _, ok := interface{}(o).(stmt.Statement); !ok {
		t.Errorf("%T should implement stmt.Statement", o)
	}
}

func TestOrderSQL(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		stmt      stmt.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			stmt.NewAsc(),
			`ASC`,
			`> ASC
`,
			nil,
		},
		{
			stmt.NewDesc(),
			`DESC`,
			`> DESC
`,
			nil,
		},
		{
			stmt.NewColumn("foo").Asc(),
			`"foo" ASC`,
			`> "foo" ASC
`,
			nil,
		},
		{
			stmt.NewColumn("foo").Desc(),
			`"foo" DESC`,
			`> "foo" DESC
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
