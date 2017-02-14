package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/keyword"
)

type groupBy struct {
	prev    clause
	columns []column
}

func newGroupBy(col column, columns ...column) groupBy {
	return groupBy{
		columns: append([]column{col}, columns...),
	}
}

func (g groupBy) node() generator.Node {
	cs := clauseNodes(g)
	gs := make([]generator.Node, len(cs))
	for i, c := range cs {
		gs[i] = c.nodeMine()
	}
	return generator.NewNodes(gs...)
}

func (g groupBy) nodeMine() generator.Node {
	gs := make([]generator.Node, len(g.columns))
	for i, c := range g.columns {
		gs[i] = c.node()
	}
	return generator.NewContainer(
		generator.NewExpression(string(keyword.GroupBy)),
		generator.NewComma(gs...),
	)
}

func (g groupBy) previous() clause {
	return g.prev
}

func (g groupBy) Having(operation comparisonOrLogicalOperation) having {
	l := newHaving(operation)
	l.prev = g
	return l
}

func (g groupBy) Limit(offset, lim int) limit {
	l := newLimit(offset, lim)
	l.prev = g
	return l
}
