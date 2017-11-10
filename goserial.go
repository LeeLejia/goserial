package goserial
import (
	"encoding/json"
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
)

type SerializableObj struct {
	length uint32
	objKey uint32
	data   map[string]interface{}
}
/**
	反序列化
 */
func Deserialize(data []byte) (obj SerializableObj,err error){
	buf1 := bytes.NewBuffer(data)
	err = binary.Read(buf1, binary.LittleEndian, &(obj.length))
	if err!=nil{
		return
	}
	err = binary.Read(buf1, binary.LittleEndian, &(obj.objKey))
	if err!=nil{
		return
	}
	vs:=data[8:]
	crc := crc32.ChecksumIEEE(vs)
	if crc !=obj.objKey {
		err=fmt.Errorf(" crc not check")
		return
	}
	err=json.Unmarshal(vs,&obj.data)
	return
}
/**
	添加一个待序列化对象
 */
func (obj *SerializableObj)Serialize(key string,v interface{})(err error) {
	if obj.data==nil{
		obj.data=make(map[string]interface{})
	}
	obj.data[key]=v
	return
}
/**
	获取序列化后数据
 */
func (obj *SerializableObj)Sum() (rs []byte,err error){
	bs,err:=json.Marshal(obj.data)
	if err!=nil{
		return bs,err
	}
	obj.objKey = crc32.ChecksumIEEE(bs)
	// 开始序列化
	buf1 := new(bytes.Buffer)
	err = binary.Write(buf1, binary.LittleEndian,bs)
	if err!=nil{
		return
	}
	buf2 := new(bytes.Buffer)
	obj.length = uint32(buf1.Len() + 8)
	err = binary.Write(buf2, binary.LittleEndian, obj.length)
	if err!=nil{
		return
	}
	// 写入hash校验
	err = binary.Write(buf2, binary.LittleEndian, obj.objKey)
	if err!=nil{
		return
	}
	// 写入数据
	err = binary.Write(buf2, binary.LittleEndian, buf1.Bytes())
	if err!=nil{
		return
	}
	return buf2.Bytes(),nil
}

/**
	获取未序列化的对象
 */
func (obj *SerializableObj)GetObj(key string,v interface{}) (err error){
	if obj.data==nil{
		return fmt.Errorf("对象内容为空！")
	}
	d:=obj.data[key]
	if d==nil{
		return fmt.Errorf("不存在该key对象！")
	}
	r,err:=json.Marshal(d)
	if err!=nil{
		return err
	}
	return json.Unmarshal(r,v)
}