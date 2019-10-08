A Go tokenizer for Dutch text.

This is a wrapper in Go for the tokenizer that is part of [Alpino](http://www.let.rug.nl/vannoord/alp/Alpino/).

The file `libtok_no_breaks.c` is a copy from the Alpino source.

The file `libtok1.c` is derived from the Alpino source by this command:

    perl -p -e 's/QDATUM|new_t_accepts|qentry|qinit|qinsert|qpeek|qremove|queue|replace_from_queue|resize_buf|t_accepts|transition_struct|trans|unknown_symbol/$&1/g' \
        libtok.c > libtok1.c

## Install

    go get github.com/pebbe/tokenize

## Docs

 * [package help](http://godoc.org/github.com/pebbe/tokenize)
