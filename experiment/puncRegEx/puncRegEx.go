package main

import (
	reloaded "reloaded/reloaded"
	"fmt"
)

func main() {
	seanStr	:= 
` ' hello world 'A ' apple .' ... ???' ??? ' ??? 'assd 
a orange '  
`  
/* ' She 'said ' hello ' world' 
101 (bin)
1A (hex)
hello (cap)
HELLO (low)
hello (up)
this is a test (cap, 2)
If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?
I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure
Don not be sad ,because sad backwards is das . And das not good
harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.
Punctuation tests are ... kinda boring ,what do you think ?
`
*/
	fmt.Println(seanStr)
	// seanStrOldPunk	:= reloaded.PunctuationFormat(seanStr)
	//fmt.Println("after punct  ", seanStrOldPunk)
	seanStr	= reloaded.PunctuationRegEx(seanStr)
	fmt.Println(seanStr)
	seanStrAtoAn	:= reloaded.AtoAnRegEx(seanStr)
	fmt.Println(" after a to A", seanStrAtoAn)
	fmt.Println(seanStr)
	seanStr	= reloaded.PunctuationRegEx(seanStr)
	fmt.Println(seanStr)
}