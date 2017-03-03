package token_test

import (
	"fmt"
	"testing"

	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/token"
)

func TestToken(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		token                    token.Token
		sprintWithStandardFormat string
		sprintWithMySQLFormat    string
		string                   string
	}{
		{
			token.QuoteStart,
			`"`,
			"`",
			"QuoteStart",
		},
		{
			token.QuoteEnd,
			`"`,
			"`",
			"QuoteEnd",
		},
		{
			token.LineStart,
			"",
			"",
			"LineStart",
		},
		{
			token.LineEnd,
			"\n",
			"\n",
			"LineEnd",
		},
		{
			token.Indent,
			"  ",
			"  ",
			"Indent",
		},
	} {
		c := c
		t.Run(fmt.Sprintf("%d Sprint(StandardFormat)", i), func(t *testing.T) {
			t.Parallel()
			got := c.token.Sprint(token.StandardIndentedFormat)
			if got != c.sprintWithStandardFormat {
				t.Error(diff.Values(got, c.sprintWithStandardFormat))
			}
		})
		t.Run(fmt.Sprintf("%d Sprint(MySQL)", i), func(t *testing.T) {
			t.Parallel()
			got := c.token.Sprint(token.MySQLIndentedFormat)
			if got != c.sprintWithMySQLFormat {
				t.Error(diff.Values(got, c.sprintWithMySQLFormat))
			}
		})
		t.Run(fmt.Sprintf("%d String()", i), func(t *testing.T) {
			t.Parallel()
			got := c.token.String()
			if got != c.string {
				t.Error(diff.Values(got, c.string))
			}
		})
	}
}
