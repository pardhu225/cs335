package main

type T struct{
	a rune
	b int 
	c rune
	d int16
	e float64
	name[10] rune
	f rune
}

func f(x T) {
	x.a = 'a'
	x.b = 47114711
	x.c = 'c'
	x.d = 1234
	x.e = 3.141592897932
  	x.f = '*'
  	x.name[0] = 'a'
	x.name[1] = 'b'
	x.name[2] = 'c'
}

func main() {
	var k T
	f(k)
}
