package token_test

import (
	"fmt"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/token"
)

func TestToken(t *testing.T) {
	for i, c := range []struct {
		tokens   token.Tokens
		standard string
		mySQL    string
	}{
		{
			token.Tokens{
				token.QuoteStart,
				token.Word("foo"),
				token.QuoteEnd,
			},
			`"foo"`,
			"`foo`",
		},
	} {
		t.Run(fmt.Sprintf("%d Standard", i), func(t *testing.T) {
			sql := c.tokens.Sprint(token.StandardFormat)
			if sql != c.standard {
				t.Error(diff.SQL(sql, c.standard))
			}
		})
		t.Run(fmt.Sprintf("%d MySQL", i), func(t *testing.T) {
			sql := c.tokens.Sprint(token.MySQLFormat)
			if sql != c.mySQL {
				t.Error(diff.SQL(sql, c.mySQL))
			}
		})
	}
}
