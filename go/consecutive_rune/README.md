For a given string s find the character c (or C) with longest consecutive repetition and return:

type Result struct {
    C rune // character
    L int  // count
}

where l (or L) is the length of the repetition. If there are two or more characters with the same l return the first in order of appearance.

For empty string return:

Result{}

https://www.codewars.com/kata/586d6cefbcc21eed7a001155/train/go