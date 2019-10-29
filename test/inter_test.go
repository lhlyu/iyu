package test

import "fmt"

type Human interface {
	Say() string
}

type Man struct {
}

func (m *Man) Say() string {
	return "man"
}

func IsNil(h interface{}) bool {
	return h == nil
}

func main() {
	var a interface{}
	var b *Man
	var c *Man
	var d Human
	var e interface{}
	a = b
	e = a

	/**
	  两条公理
	  1. 只有interface才需要考虑eface和iface,其他只要考虑零值
	  2. (_type && data == nil || tab && data == nil) == nil

	  a 实现类型 eface :  _type 指向 *Man, data = nil    =>  not nil
	  b 类型是 *Man , data = nil  => nil
	  c 类型是 *Man , data = nil  => nil
	  d 实现类型是 iface： tab 没有具体类型  nil , data = nil  => nil
	  e 实现类型 eface :  _type 指向 *Man, data = nil    =>  not nil

	  b == nil   ==> true
	  e == nil   ==> false
	  a == c     ==> true   这里类型相同 比较 data
	  a == d     ==> false  类型不同
	  c == d     ==> false  类型不同
	  e == b     ==> true   类型相同 data 相同
	  IsNil(c)   ->  参数 i 是 eface , _type => *Man , data = nil ,只有 _type && data 都为 nil 才为 nil  ==> false
	  IsNil(d)   ->  nil  ==> true

	*/

	fmt.Println(b == nil) // (1)
	/**
	  b data = nil                 => true
	*/
	fmt.Println(e == nil) // (2)
	/**
	  e eface
	  _type = *Man data = nil      => false
	*/
	fmt.Println(a == c) // (3)
	/**
	  c 类型是
	  a eface
	  _type = *Man data = nil       => true X
	*/
	fmt.Println(a == d) // (4)
	/**
	  a eface
	  _type = *Man data = nil
	  d iface
	  tab = nil data = nil            => false
	*/
	fmt.Println(c == d) // (5)
	/**
	  c  data = nil
	  d iface
	  tab = nil data = nil  类型不同  => false
	*/
	fmt.Println(e == b) // (6)
	/**
	  e eface
	  _type = *Man data = nil
	  b  data = nil   类型相同     => true
	*/
	fmt.Println(IsNil(c)) // (7)
	/**
	  c 的类型是 *Man 参数 传入 i 是 eface, _type = *Man  data = nil   => false
	*/
	fmt.Println(IsNil(d)) // (8)
	/**
	  d iface
	  tab = nil data = nil  true
	*/
	return

	fmt.Println(a == nil)
	// (1) false
	// a是eface类型，_type指向的是*Man类型，
	// data指向的是nil，所以此题为false

	fmt.Println(e == nil)
	// (2) false
	// 同理，e为eface类型，_type也是指向的*Man类型

	fmt.Println(a == c)
	// (3) true
	// a的_type是*Man类型，data是nil
	// c的data也是nil

	fmt.Println(a == d)
	// (4) false
	// a为eface类型，d为iface类型，而且d的itab指向的是nil，data也是nil
	// 因为d没有具体到哪种数据类型

	fmt.Println(c == d)
	// (5) false
	// c和d其实是两种不同的数据类型

	fmt.Println(e == b)
	// (6) true
	// 分析见(4)

	fmt.Println(IsNil(c))
	// (7) false
	// c是*Man类型，以参数的形式传入IsNil方法
	// 虽然c指向的是nil，但是参数i的_type指向的是*Man，所以i不为nil

	fmt.Println(IsNil(d))
	// (8) true
	// d没有指定具体的类型，所以d的itab指向的是nil，data也是nil                   ==> true
}
