#### 前言

    由于golang语言特性，在服务器程序开发中我们难以实现热更新，导致每次的  
    系统更新都要对服务器进行重启。这样同时导致了服务器中间数据的丢失。如
    session和一些分散的状态值。
    我们采取的一种方式，是在将服务器数据进行序列化，在重启前保存，并在重启  
    后对数据进行恢复。
    
    
#### 使用
* 导入包
```go
    import "github.com/cjwddz/goserial"
```
* 序列化
```go
    // 序列化对象
    obj:=goserial.SerializableObj{}
    
    var a = "test string"
    var b = 123
    var d = struct {
            		A string
            		B int
            		CC C
            	}{A:"sb",B:123,CC:C{25}}
            	
   // 将a,b,c添加到序列化对象
   obj.Serialize("a",a)
   obj.Serialize("b",b)
   obj.Serialize("c",c)
   // 序列化，得到了sv为[]byte,是序列化后的结果
   sv,err:=obj.Sum()
   if err!=nil{
       fmt.Println(err.Error())
       return
   }
   // 写到文件
   ioutil.WriteFile("./objs.bin",sv,0666)

```

* 反序列化

```go
    // 将[]byte解码为对象
    mm,err:=goserial.Deserialize(sv)
    if err!=nil{
        fmt.Println(err.Error())
        return
    }
    // 解析a对象到string
    var a string
    if err=mm.GetObj("a",&a);err!=nil{
        fmt.Println(err.Error())
    }else{
        fmt.Println(all)
    }
    // 解析C对象到结构体C
    var c C
    if err=mm.GetObj("C",&c);err!=nil{
        fmt.Println(err.Error())
    }else{
        fmt.Println(fmt.Sprintf("c.D=%d",c.D))
    }

```