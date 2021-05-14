package main

import (
    "fmt"
    str "strings" // giving alias
)

var print = fmt.Println // assigning a  fun to a var ~ giving alias name

func main() {
	print("\nHere are some useful functions' implementation from the strings library using a string - 'test'\n")

    print("Contains:  ", str.Contains("test", "es"))
    print("Count:     ", str.Count("test", "t"))
    print("HasPrefix: ", str.HasPrefix("test", "te"))
    print("HasSuffix: ", str.HasSuffix("test", "st"))
    print("Index:     ", str.Index("test", "e"))
    print("Join:      ", str.Join([]string{"a", "b"}, "-"))
    print("Repeat:    ", str.Repeat("a", 5))
    print("Replace:   ", str.Replace("foo", "o", "0", -1))
    print("Replace:   ", str.Replace("foo", "o", "0", 1))
    print("Split:     ", str.Split("a-b-c-d-e", "-"))
    print("ToLower:   ", str.ToLower("TEST"))
    print("ToUpper:   ", str.ToUpper("test"))
    print()

    print("Len: ", len("hello"))
    print("Char:", "hello"[1])
}