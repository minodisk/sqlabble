package token

// Generate converts the tokens to text according to format.
func Generate(tokens Tokens, format Format) string {
	if format.IsBreaking {
		return scanWithBreaking(tokens, format).Sprint(format)
	}
	return scanWithoutBreaking(tokens, format).Sprint(format)
}

func scanWithBreaking(tokens Tokens, format Format) Tokens {
	ts := Tokens{}
	last := len(tokens) - 1
	for i, t := range tokens {
		var next Token
		if i < last {
			next = tokens[i+1]
		}
		switch t {
		case LineStart, Indent, LineEnd:
			ts = append(ts, t)
			continue
		case Dot, ParenStart, FuncParenStart, QuoteStart:
			ts = append(ts, t)
			continue
		default:
			switch next {
			case Dot, Comma, QuoteEnd, ParenEnd, FuncParenStart, FuncParenEnd, LineEnd:
				ts = append(ts, t)
				continue
			default:
				ts = append(ts, t, Space)
				continue
			}
		}
	}
	return ts
}

func scanWithoutBreaking(tokens Tokens, format Format) Tokens {
	ts1 := Tokens{}
	for _, t := range tokens {
		switch t {
		case LineStart, Indent, LineEnd:
			continue
		default:
			ts1 = append(ts1, t)
			continue
		}
	}

	ts2 := Tokens{}
	last := len(ts1) - 1
	for i, t := range ts1 {
		var next Token
		if i < last {
			next = ts1[i+1]
		}
		switch t {
		case Dot, ParenStart, FuncParenStart, QuoteStart:
			ts2 = append(ts2, t)
			continue
		default:
			if i == last {
				ts2 = append(ts2, t)
				continue
			}
			switch next {
			case Dot, Comma, QuoteEnd, ParenEnd, FuncParenStart, FuncParenEnd:
				ts2 = append(ts2, t)
				continue
			default:
				ts2 = append(ts2, t, Space)
				continue
			}
		}
	}

	return ts2
}
