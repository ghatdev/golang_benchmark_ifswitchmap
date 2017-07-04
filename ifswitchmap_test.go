package main

import (
	"testing"
)

var words = []string{"super", "califragi", "listic", "expiali", "docius"}
var wlen = len(words)

func f() int {
	return 1
}

func BenchmarkSwitchIntShort(b *testing.B) {
	m := 0
	for i := 0; i < b.N; i++ {
		switch words[i%wlen] {
		case "super":
			m++
		case "califragi":
			m++
		case "listic":
			m++
		case "expiali":
			m++
		case "docius":
			m++
		}
	}

}

func BenchmarkIfIntShort(b *testing.B) {
	m := 0
	for i := 0; i < b.N; i++ {
		w := words[i%wlen]
		if w == "super" {
			m++
		} else if w == "califragi" {
			m++
		} else if w == "listic" {
			m++
		} else if w == "expiali" {
			m++
		} else if w == "docius" {
			m++
		}
	}

}

func BenchmarkMapIntShort(b *testing.B) {
	m := 0
	a := map[string]int{
		"super":     1,
		"califragi": 1,
		"listic":    1,
		"expiali":   1,
		"docius":    1,
	}

	for i := 0; i < b.N; i++ {
		w := words[i%wlen]
		m += a[w]
	}

}

func BenchmarkSwitchFuncShort(b *testing.B) {
	m := 0
	for i := 0; i < b.N; i++ {
		switch words[i%wlen] {
		case "super":
			m += f()
		case "califragi":
			m += f()
		case "listic":
			m += f()
		case "expiali":
			m += f()
		case "docius":
			m += f()
		}
	}

}

func BenchmarkIfFuncShort(b *testing.B) {
	m := 0
	for i := 0; i < b.N; i++ {
		w := words[i%wlen]
		if w == "super" {
			m += f()
		} else if w == "califragi" {
			m += f()
		} else if w == "listic" {
			m += f()
		} else if w == "expiali" {
			m += f()
		} else if w == "docius" {
			m += f()
		}
	}

}

func BenchmarkMapFuncShort(b *testing.B) {
	m := 0
	a := map[string]func() int{
		"super":     f,
		"califragi": f,
		"listic":    f,
		"expiali":   f,
		"docius":    f,
	}

	for i := 0; i < b.N; i++ {
		w := words[i%wlen]
		m += a[w]()
	}

}

var wordsLong = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
var wordsLongLen = len(wordsLong)

func BenchmarkSwitchIntLong(b *testing.B) {
	m := 0
	for i := 0; i < b.N; i++ {
		switch wordsLong[i%wordsLongLen] {
		case "a":
			m++
		case "b":
			m++
		case "c":
			m++
		case "d":
			m++
		case "e":
			m++
		case "f":
			m++
		case "g":
			m++
		case "h":
			m++
		case "i":
			m++
		case "j":
			m++
		case "k":
			m++
		case "l":
			m++
		case "m":
			m++
		case "n":
			m++
		case "o":
			m++
		case "p":
			m++
		case "q":
			m++
		case "r":
			m++
		case "s":
			m++
		case "t":
			m++
		case "u":
			m++
		case "v":
			m++
		case "w":
			m++
		case "x":
			m++
		case "y":
			m++
		case "z":
			m++
		}
	}
}

func BenchmarkIfIntLong(b *testing.B) {
	m := 0
	for i := 0; i < b.N; i++ {
		w := wordsLong[i%wordsLongLen]
		if w == "a" {
			m++
		} else if w == "b" {
			m++
		} else if w == "c" {
			m++
		} else if w == "d" {
			m++
		} else if w == "e" {
			m++
		} else if w == "f" {
			m++
		} else if w == "g" {
			m++
		} else if w == "h" {
			m++
		} else if w == "i" {
			m++
		} else if w == "j" {
			m++
		} else if w == "k" {
			m++
		} else if w == "l" {
			m++
		} else if w == "m" {
			m++
		} else if w == "n" {
			m++
		} else if w == "o" {
			m++
		} else if w == "p" {
			m++
		} else if w == "q" {
			m++
		} else if w == "r" {
			m++
		} else if w == "s" {
			m++
		} else if w == "t" {
			m++
		} else if w == "u" {
			m++
		} else if w == "v" {
			m++
		} else if w == "w" {
			m++
		} else if w == "x" {
			m++
		} else if w == "y" {
			m++
		} else if w == "z" {
			m++
		}
	}
}

func BenchmarkMapIntLong(b *testing.B) {
	m := 0
	a := map[string]int{
		"a": 1,
		"b": 1,
		"c": 1,
		"d": 1,
		"e": 1,
		"f": 1,
		"g": 1,
		"h": 1,
		"i": 1,
		"j": 1,
		"k": 1,
		"l": 1,
		"m": 1,
		"n": 1,
		"o": 1,
		"p": 1,
		"q": 1,
		"r": 1,
		"s": 1,
		"t": 1,
		"u": 1,
		"v": 1,
		"w": 1,
		"x": 1,
		"y": 1,
		"z": 1,
	}

	for i := 0; i < b.N; i++ {
		w := wordsLong[i%wordsLongLen]
		m += a[w]
	}

}

func BenchmarkSwitchFuncLong(b *testing.B) {
	m := 0
	for i := 0; i < b.N; i++ {
		switch wordsLong[i%wordsLongLen] {
		case "a":
			m += f()
		case "b":
			m += f()
		case "c":
			m += f()
		case "d":
			m += f()
		case "e":
			m += f()
		case "f":
			m += f()
		case "g":
			m += f()
		case "h":
			m += f()
		case "i":
			m += f()
		case "j":
			m += f()
		case "k":
			m += f()
		case "l":
			m += f()
		case "m":
			m += f()
		case "n":
			m += f()
		case "o":
			m += f()
		case "p":
			m += f()
		case "q":
			m += f()
		case "r":
			m += f()
		case "s":
			m += f()
		case "t":
			m += f()
		case "u":
			m += f()
		case "v":
			m += f()
		case "w":
			m += f()
		case "x":
			m += f()
		case "y":
			m += f()
		case "z":
			m += f()
		}
	}
}

func BenchmarkIfFuncLong(b *testing.B) {
	m := 0
	for i := 0; i < b.N; i++ {
		w := wordsLong[i%wordsLongLen]
		if w == "a" {
			m += f()
		} else if w == "b" {
			m += f()
		} else if w == "c" {
			m += f()
		} else if w == "d" {
			m += f()
		} else if w == "e" {
			m += f()
		} else if w == "f" {
			m += f()
		} else if w == "g" {
			m += f()
		} else if w == "h" {
			m += f()
		} else if w == "i" {
			m += f()
		} else if w == "j" {
			m += f()
		} else if w == "k" {
			m += f()
		} else if w == "l" {
			m += f()
		} else if w == "m" {
			m += f()
		} else if w == "n" {
			m += f()
		} else if w == "o" {
			m += f()
		} else if w == "p" {
			m += f()
		} else if w == "q" {
			m += f()
		} else if w == "r" {
			m += f()
		} else if w == "s" {
			m += f()
		} else if w == "t" {
			m += f()
		} else if w == "u" {
			m += f()
		} else if w == "v" {
			m += f()
		} else if w == "w" {
			m += f()
		} else if w == "x" {
			m += f()
		} else if w == "y" {
			m += f()
		} else if w == "z" {
			m += f()
		}
	}
}

func BenchmarkMapFuncLong(b *testing.B) {
	m := 0
	a := map[string]func() int{
		"a": f,
		"b": f,
		"c": f,
		"d": f,
		"e": f,
		"f": f,
		"g": f,
		"h": f,
		"i": f,
		"j": f,
		"k": f,
		"l": f,
		"m": f,
		"n": f,
		"o": f,
		"p": f,
		"q": f,
		"r": f,
		"s": f,
		"t": f,
		"u": f,
		"v": f,
		"w": f,
		"x": f,
		"y": f,
		"z": f,
	}

	for i := 0; i < b.N; i++ {
		w := wordsLong[i%wordsLongLen]
		m += a[w]()
	}
}
