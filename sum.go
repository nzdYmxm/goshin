package goshin

import (
	"fmt"
)

var say = fmt.Println

func ShowUp() {
	say("goshin is a wrapper above genshinAPI to help developers use it more clearly")
	say("To clarify with character build just call build/<character>")
}
