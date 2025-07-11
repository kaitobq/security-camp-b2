package main

import (
	"fmt"
	"image"
	"io"
	"math"
	"strings"
	"time"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Abser interface {
	Abs() float64
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describe2(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v2 := Vertex{3, 4}
	v2.Scale(10)
	fmt.Println(v2.Abs())

	var a Abser
	f2 := MyFloat(-math.Sqrt2)
	v3 := Vertex{3, 4}

	a = f2
	a = &v3

	a = v

	fmt.Println(a.Abs())

	var i I
	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()

	var i2 interface{}
	describe2(i2)

	i2 = 42
	describe2(i2)

	i2 = "hello"
	describe2(i2)

	var i3 interface{} = "hello"

	s := i3.(string)
	fmt.Println(s)

	s, ok := i3.(string)
	fmt.Println(s, ok)

	f3, ok := i3.(float64)
	fmt.Println(f3, ok)

	// f3 = i3.(float64) // panic
	// fmt.Println(f3)

	do(21)
	do("hello")
	do(true)

	a2 := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a2, z)

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	if err := run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))

	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
