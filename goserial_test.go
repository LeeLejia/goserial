package goserial

import (
	"testing"
	"fmt"
)

func TestGoSerial(t *testing.T) {
	m:=initTestData()
	obj:=SerializableObj{}
	obj.Serialize("string",m["string"])
	obj.Serialize("int",m["int"])
	obj.Serialize("bool",m["bool"])
	obj.Serialize("struct",m["struct"])
	obj.Serialize("C",m["C"])

	sv,err:=obj.Sum()
	if err!=nil{
		fmt.Println(err.Error())
		t.Fail()
		return
	}
	mm,err:=Deserialize(sv)
	if err!=nil{
		fmt.Println(err.Error())
		t.Fail()
		return
	}
	var all interface{}
	if err=mm.GetObj("struct",&all);err!=nil{
		fmt.Println(err.Error())
		t.Fail()
	}else{
		fmt.Println(all)
	}
	var c C
	if err=mm.GetObj("C",&c);err!=nil{
		fmt.Println(err.Error())
		t.Fail()
	}else{
		fmt.Println(fmt.Sprintf("c.D=%d",c.D))
	}
}
type C struct {
	D int
}

func initTestData() map[string]interface{}{
	m:=make(map[string]interface{})
	m["string"]="test string"
	m["int"]=122
	m["bool"]=true
	m["C"]=C{25}
	st:= struct {
		A string
		B int
		CC C
	}{A:"sb",B:123,CC:C{25}}
	m["struct"]=st
	return m
}