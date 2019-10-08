package test

import (
	"bytes"
	"fmt"
)

//https://juejin.im/post/5bf909cb51882521c8114523

func TestBytes()  {
	TestBytesBuffer()
	TestBytesBytes()
	TestBytesReader()
}


func TestBytesBuffer()  {
	//bytes.buffer是一个缓冲byte类型的缓冲器

	buf1 := bytes.NewBufferString("hello")
	buf2 := bytes.NewBuffer([]byte("hello"))
	buf3:= bytes.NewBuffer([]byte{'h','e','l','l','o'})
	//以上三者等效,输出//hello

	buf4 := bytes.NewBufferString("")
	buf5 := bytes.NewBuffer([]byte{})
	//以上两者等效,输出//""

	fmt.Println(buf1.String(),buf2.String(),buf3.String(),buf4,buf5,1)



	s := []byte("world")
	buf1.Write(s)                 //将s这个slice添加到buf的尾部
	fmt.Println(buf1.String())   //helloworld

	s2 := "world"
	buf2.WriteString(s2)           //将string写入到buf的尾部
	fmt.Println(buf2.String())   //helloworld

	var s3 byte = '?'    //将一个byte类型的数据放到缓冲器的尾部
	buf3.WriteByte(s3)         //将s写到buf的尾部
	fmt.Println(buf3.String()) //hello？

	//WriteRune方法，将一个rune类型的数据放到缓冲器的尾部
	var s4 rune  =  '好'
	buf1.WriteRune(s4)
	fmt.Println(buf1.String()) //helloworld好

	s5 := "你好"
	buf2.WriteString(s5)
	fmt.Println(buf2.String()) //helloworld你好
	for _,b := range buf2.Bytes(){
		fmt.Println(b)
	}
	/*没法解析到 中文 rune
	104
	101
	108
	108
	111
	119
	111
	114
	108
	100
	228
	189
	160
	229
	165
	189*/

	b,n,_ := buf2.ReadRune()  //取出第一个rune
	fmt.Println(b,n) //104 1
}

func TestBytesBytes() {
	buf1 := []byte("hello世界")
	fmt.Println(len(buf1)) //11
	fmt.Println(len(bytes.Runes(buf1)))//7
}

func TestBytesReader()  {
	buf1 := []byte("hello世界")
	r := bytes.NewReader(buf1)
	fmt.Println(r.Len()) //11
	fmt.Println(r.Size()) //11

}

