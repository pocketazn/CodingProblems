package _2021_08

// cons(a, b) constructs a pair, and car(pair) and cdr(pair)
//returns the first and last element of that pair. For example,
//car(cons(3, 4)) returns 3, and cdr(cons(3, 4)) returns 4.
//
//Given this implementation of cons:
//
//def cons(a, b):
//    def pair(f):
//        return f(a, b)
//    return pair
//Implement car and cdr.

// solve time. 3 Mins

type Pair struct {
	First  int
	Second int
}

func Cons(f int, s int) Pair {
	return Pair{
		First:  f,
		Second: s,
	}
}

func Car(p Pair) int {
	return p.First
}

func Cdr(p Pair) int {
	return p.Second
}
