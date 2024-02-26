package main

import "fmt"

func main() {

	/*
		ascii
		American Standard Code for Information Interchange
		2^8 (1 byte) = 256 unique values

		unicode
		2^32 (4 bytes) = 4,294,967,296 unique values
		ore than enough to account for every character
		in every language

		utf-8
		(up to 4 bytes)
		stores unicode as binary
		If a character needs 1 byte that's all it will use.
		If a character needs 4 bytes it will use 4 bytes.
		variable length encoding = efficient memory use
		common characters like "C" take 8 bits,
		rare characters like "emoji" take 32 bits.
		Other characters take 16 or 24 bits.
	*/

	fmt.Println(`
	UTF-8 saves space.
	In UTF-8, common characters like "C" take 8 bits,
	while rare characters like "emojis" take 32 bits.
	Other characters take 16 or 24 bits.
	A blog post like this one takes about four times less space
	in UTF-8, than it would in UTF-32.

	So it loads four times faster.
	`)

}
