package internal

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

func Post(s string, withLineBreaks bool) string {
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

	return s
}
