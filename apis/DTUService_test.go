package DTU

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"testing"
)

func TestSayHello_Descriptor(t *testing.T) {
	info := &SayHello{
		Name:  "Atlan",
		Int:   1,
		Email: "1@unix.com",
		Info: &SayHello_HeroInfo{
			Age:    20,
			Gender: 1,
		},
	}

	out, err := proto.Marshal(info)
	if err != nil {
		log.Fatalln("Failed to encode SayHello:", err)
	}
	fmt.Println("success value:", string(out))
	if err := ioutil.WriteFile("F:\\github\\DTU\\apis\\test.txt", out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
}

func Test_SayhelloRead(t *testing.T) {
	in, err := ioutil.ReadFile("F:\\github\\DTU\\apis\\test.txt")
	if err != nil {
		log.Fatalln("error reading file:", err)
	}
	info := &SayHello{}
	if err := proto.Unmarshal(in, info); err != nil {
		log.Fatalln("failed to parse info", err)
	}
	fmt.Println(info)
}
