package main

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	complexpb "protobuf-example/src/complex"
	enumpb "protobuf-example/src/enum_example"
	simplepb "protobuf-example/src/simple"
)

func main(){
	sm := doSimple()
	readAndWriteDisk(sm)
	readAndWriteJSON(sm)

	doEnum()
	doComplex()
}

func doComplex(){
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id: 1,
			Name: "First Message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id: 2,
				Name: "Second Message",
			},
			&complexpb.DummyMessage{
				Id: 3,
				Name: "Third Message",
			},
		},
	}

	fmt.Println(cm)
}

func doEnum(){
	em := enumpb.EnumMessage{
		Id: 1234,
		DayOfTheWeek: enumpb.DayOfTheWeek_MONDAY,
	}
	em.DayOfTheWeek = enumpb.DayOfTheWeek_SATURDAY
	fmt.Println(em)
}

func readAndWriteJSON(sm proto.Message){
	smAsString := toJSON(sm)
	fmt.Println(smAsString)

	sm2 := &simplepb.SimpleMessage{}
	fromJSON(smAsString,  sm2)
	fmt.Println("Successfully created proto struc:", sm2)
}
func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)

	if err != nil{
		log.Fatalln("Can't convert to JSON", err)
		return ""
	}

	return out
}

func fromJSON(in string, pb proto.Message){
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil{
		log.Fatalln("Couldn't unmarchal the JSON into the pb struct", err)
	}
}

func readAndWriteDisk(sm proto.Message){
	writeToFile("simple.bin", sm)
	sm2 := &simplepb.SimpleMessage{}
	readfromFile("simple.bin", sm2)
	fmt.Println("Read the content:", sm2)
}

func writeToFile(fname string, pb proto.Message) error{
	out, err := proto.Marshal(pb)
	if err != nil{
		log.Fatalln("Can't serialize to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil{
		log.Fatalln("Can't write to file", err)
		return err
	}

	fmt.Println("Data has been written!")
	return nil
}

func readfromFile(fname string, pb proto.Message) error{
	in, err := ioutil.ReadFile(fname)

	if err != nil{
		log.Fatalln("Something went wrong when reading the file", err)
		return err
	}

	if err := proto.Unmarshal(in, pb); err != nil{
		log.Fatalln("Couldn't put the bytes into the protocol buffers struct", err)
		return err
	}

	return nil

}

func doSimple() *simplepb.SimpleMessage{
	sm := simplepb.SimpleMessage{
		Id: 1234,
		IsSimple: true,
		Name: "My Simple Message",
		SampleList: []int32{ 1, 5, 7, 8},
	}

	fmt.Println(sm)

	sm.Name = "Rename Simple Message"

	fmt.Println(sm.GetId())
	//GetId() nill체크해줌

	return &sm
}


