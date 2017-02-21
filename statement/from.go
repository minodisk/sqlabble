package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
)

type From struct {
	prev   Clause
	joiner Joiner
}

func NewFrom(joiner Joiner) From {
	return From{
		joiner: joiner,
	}
}

func (f From) nodeize() (token.Tokenizer, []interface{}) {
	return nodeizeClauses(f)
}

func (f From) self() (token.Tokenizer, []interface{}) {
	middle, values := f.joiner.nodeize()
	return token.NewContainer(
		token.NewLine(token.Word(keyword.From)),
	).SetMiddle(
		middle,
	), values
}

func (f From) previous() Clause {
	return f.prev
}

func (f From) Where(op ComparisonOrLogicalOperation) Where {
	w := NewWhere(op)
	w.prev = f
	return w
}

func (f From) GroupBy(col Column, columns ...Column) GroupBy {
	g := NewGroupBy(col, columns...)
	g.prev = f
	return g
}

func (f From) OrderBy(orders ...Order) OrderBy {
	o := NewOrderBy(orders...)
	o.prev = f
	return o
}

func (f From) Limit(count int) Limit {
	l := NewLimit(count)
	l.prev = f
	return l
}
