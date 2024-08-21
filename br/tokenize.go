package br

/*
#cgo CFLAGS: -finput-charset=UTF-8

#include <wchar.h>
int t_accepts1 (wchar_t *in,wchar_t *out,int max);
void tokenize_break (wchar_t *in, wchar_t *out, int *nchar, int *retv)
{
	size_t
		bufsize;
	bufsize = wcslen (in) * 2;
	*retv = t_accepts1 (in, out, bufsize);
	if (*retv == 1)
		*nchar = wcslen(out);
}
*/
import "C"

import (
	"github.com/pebbe/tokenize/internal"
)

func Dutch(text string) (string, error) {
	wc1 := make([]C.wchar_t, len(text)+1)
	i := 0
	for _, c := range text {
		wc1[i] = C.wchar_t(c)
		i++
	}
	wc1[i] = C.wchar_t(0)

	var nchar C.int
	var retv C.int
	wc2 := make([]C.wchar_t, 2*len(wc1)+2)

	C.tokenize_break(&wc1[0], &wc2[0], &nchar, &retv)

	switch retv {
	case 0:
		return "", internal.ErrImpossible
	case 2:
		return "", internal.ErrTooLong
	}

	r := make([]rune, int(nchar))
	for i := range r {
		r[i] = rune(wc2[i])
	}
	return internal.Post(string(r), true), nil
}
