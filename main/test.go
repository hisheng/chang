package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	example "github.com/hisheng/chang"
	"log"
	"time"
)


/*test 异步的时间，这个程序异步 执行时间是 10秒，如果同步的话，需要是 20秒*/
func testYibu()  {
	start := time.Now().Unix()

	sum := make(chan int)

	go func() {
		time.Sleep(time.Second*10)
		sum <- 2
	}()


	time.Sleep(time.Second*10)
	fmt.Println("sum is :")
	fmt.Println(<-sum)

	end := time.Now().Unix()
	fmt.Println(end-start)

}

/*测试 proto*/
func testProto()  {
	test := &example.Test{
		Label: proto.String("hello"),
		Type:  proto.Int32(17),
		Reps:  []int64{1, 2, 3},
	}
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	newTest := &example.Test{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	// Now test and newTest contain the same data.
	if test.GetLabel() != newTest.GetLabel() {
		log.Fatalf("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
	}
}
