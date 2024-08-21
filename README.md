A Go tokenizer for Dutch text.

This is a wrapper in Go for the tokenizer that is part of [Alpino](http://www.let.rug.nl/vannoord/alp/Alpino/).

The file `nobr/libtok_no_breaks.c` is a copy from the Alpino source.

The file `br/libtok1.c` is derived from the Alpino source by this command:

    perl -p -e 's/QDATUM|new_t_accepts|qentry|qinit|qinsert|qpeek|qremove|queue|replace_from_queue|resize_buf|t_accepts|transition_struct|trans|unknown_symbol/$&1/g' \
        libtok.c > br/libtok1.c

## Install

    go get github.com/pebbe/tokenize

If you get an error like this:

    # github.com/pebbe/tokenize
    libtok1.c:1:1: error: unknown type name ‘version’
     version https://git-lfs.github.com/spec/v1
     ^
    libtok1.c:1:14: error: expected ‘=’, ‘,’, ‘;’, ‘asm’ or ‘__attribute__’ before ‘:’ token
     version https://git-lfs.github.com/spec/v1
                  ^

... run this command, and try again: `git lfs install`

If you get an error like this:

    git: 'lfs' is not a git command. See 'git --help'.

... install `git-lfs` from https://git-lfs.github.com/



## Docs

 * [package help](http://godoc.org/github.com/pebbe/tokenize)
