package tokenize

import (
	"github.com/pebbe/tokenize/br"
	"github.com/pebbe/tokenize/nobr"
)

func Dutch(text string, withLineBreaks bool) (string, error) {
	if withLineBreaks {
		return br.Dutch(text)
	}
	return nobr.Dutch(text)
}
