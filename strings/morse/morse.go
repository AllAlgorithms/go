package morse

import (
	"fmt"
	"strings"
)

var morseAlphabet = map[rune]string{
	'a': ".-",
	'b': "-...",
	'c': "-.-.",
	'd': "-..",
	'e': ".",
	'f': "..-.",
	'g': "--.",
	'h': "....",
	'i': "..",
	'j': ".---",
	'k': "-.-",
	'l': ".-..",
	'm': "--",
	'n': "-.",
	'o': "---",
	'p': ".--.",
	'q': "--.-",
	'r': ".-.",
	's': "...",
	't': "-",
	'u': "..-",
	'v': "...-",
	'w': ".--",
	'x': "-..-",
	'y': "-.--",
	'z': "--..",
	'1': ".----",
	'2': "..---",
	'3': "...--",
	'4': "....-",
	'5': ".....",
	'6': "-....",
	'7': "--...",
	'8': "---..",
	'9': "----.",
	'0': "-----",
}

var inverseAlphabet map[string]rune

func init() {
	inverseAlphabet = make(map[string]rune)
	for r, c := range morseAlphabet {
		inverseAlphabet[c] = r
	}
}

func Encode(s string) (string, error) {
	var ret []string
	for _, r := range strings.ToLower(s) {
		if _, ok := morseAlphabet[r]; !ok {
			return "", fmt.Errorf("unknown char %v", string(r))
		}

		ret = append(ret, morseAlphabet[r])
	}

	return strings.Join(ret, " "), nil
}

func Decode(s string) (string, error) {
	ret := ""
	for _, c := range strings.Split(s, " ") {
		if _, ok := inverseAlphabet[c]; !ok {
			return "", fmt.Errorf("unknown code %v", c)
		}
		ret += string(inverseAlphabet[c])
	}

	return ret, nil
}
