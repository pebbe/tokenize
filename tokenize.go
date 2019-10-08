package tokenize

/*
#cgo CFLAGS: -finput-charset=UTF-8

#include <wchar.h>
int t_accepts (wchar_t *in,wchar_t *out,int max);
int t_accepts1 (wchar_t *in,wchar_t *out,int max);
void tokenize (wchar_t *in, wchar_t *out, int *nchar, int *retv, int with_line_breaks)
{
	size_t
		bufsize;
	bufsize = wcslen (in) * 2;
	if (with_line_breaks)
		*retv = t_accepts1 (in, out, bufsize);
	else
		*retv = t_accepts (in, out, bufsize);
	if (*retv == 1)
		*nchar = wcslen(out);
}
*/
import "C"

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrImpossible = errors.New("no transduction possible")
	ErrTooLong    = errors.New("length of transduction would be > max")

	reTuut         = regexp.MustCompile("['`’\"] \\pL+['`’\"]-")
	reBuitenGewoon = regexp.MustCompile(`\( (\pL+-?\))`)
	reFeit         = regexp.MustCompile("(?m)^(\\([^\n)]*\\)) (\\p{Lu})")
	reHuisTuin     = regexp.MustCompile("\\b(\\pL+) -([^ \n][^-\n]*[^ \n])- (\\pL+)\\b")
	reEndSpace     = regexp.MustCompile(" +\n")
)

func Dutch(text string, withLineBreaks bool) (string, error) {
	wc1 := make([]C.wchar_t, len(text)+1)
	i := 0
	for _, c := range text {
		wc1[i] = C.wchar_t(c)
		i++
	}
	wc1[i] = C.wchar_t(0)

	wb := C.int(0)
	if withLineBreaks {
		wb = C.int(1)
	}

	var nchar C.int
	var retv C.int
	wc2 := make([]C.wchar_t, 2*len(wc1)+2)

	C.tokenize(&wc1[0], &wc2[0], &nchar, &retv, wb)

	switch retv {
	case 0:
		return "", ErrImpossible
	case 2:
		return "", ErrTooLong
	}

	r := make([]rune, int(nchar))
	for i := range r {
		r[i] = rune(wc2[i])
	}
	s := string(r)

	//
	// post-processing copied from Perl script 'tokenize_more' in Alpino
	//

	// ## ' tuut'-vorm    --> 'tuut'-vorm
	// s/(['`’"]) (\p{L}+\g1-)/$1$2/g;
	s = reTuut.ReplaceAllStringFunc(s, func(s string) string {
		r := []rune(s)
		n := len(r)
		if r[0] == r[n-2] {
			return string(r[:1]) + string(r[2:])
		}
		return s
	})

	// ## ( buiten)gewoon --> (buitengewoon)
	// s/[(] (\p{L}+-?[)])/($1/g;
	s = reBuitenGewoon.ReplaceAllString(s, `($1`)

	// ## ( Dat is een feit ) Ik ...
	// } elsif (/^[(][^)]*[)] (?=\p{Lu})/o) {
	//     $_= $` . $& . "\n" . $'; #'
	// }
	if withLineBreaks {
		s = reFeit.ReplaceAllString(s, "$1\n$2")
	}

	// ## attempts to distinguish various use of -
	// ## "huis- tuin- en keuken"  should be left alone
	// ## "ik ga -zoals gezegd- naar huis" will be rewritten into
	// ## "ik ga - zoals gezegd - naar huis"
	// if(/[ ][-]([^ ][^-]*[^ ])[-][ ]/) {
	//     $prefix=$`;
	//     $middle=$1;
	//     $suffix=$';   # '
	//     if ($prefix !~ /(en |of )$/ &&
	//         $suffix !~ /^(en |of )/) {
	//         $_ = "$prefix - $middle - $suffix";
	//     }
	// }
	s = reHuisTuin.ReplaceAllStringFunc(s, func(m string) string {
		if strings.HasPrefix(m, "en ") || strings.HasPrefix(m, "of ") {
			return m
		}
		if strings.HasSuffix(m, " en") || strings.HasSuffix(m, " of") {
			return m
		}
		return reHuisTuin.ReplaceAllString(m, "$1 - $2 - $3")
	})

	// ## remove spaces at end of line
	s = reEndSpace.ReplaceAllString(s, "\n")

	return s, nil
}
