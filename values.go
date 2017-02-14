package sqlabble

import (
	"fmt"

	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/keyword"
)

type values struct {
	prevClause clause
	prev       vals
	vals       []interface{}
}

func newValues(vals ...interface{}) values {
	return values{
		vals: vals,
	}
}

func (v values) node() generator.Node {
	vs := valuesNodes(v)
	es := make([]generator.Node, len(vs))
	for i, v := range vs {
		es[i] = v.expression()
	}
	g := generator.NewContainer(
		generator.NewExpression(keyword.Values),
		generator.NewComma(es...),
	)

	if len(vs) > 0 {
		if c := vs[0].clause(); c != nil {
			return generator.NewNodes(
				c.node(),
				g,
			)
		}
	}

	return g
}

func (v values) expression() generator.Expression {
	return generator.NewExpression(
		fmt.Sprintf("(%s)", placeholders(len(v.vals))),
		v.vals...,
	)
}

func (v values) clause() clause {
	return v.prevClause
}

func (v values) previous() vals {
	return v.prev
}

func (v values) Values(vals ...interface{}) values {
	f := newValues(vals...)
	f.prev = v
	return f
}