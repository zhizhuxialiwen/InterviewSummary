# redis数据库

## 1、介绍

Redis: remote dictionary server(远程字典服务器）。
性能非常高，单机达到15qps，通常适合做缓存，也可以持久化。开源、高性能分布式内存数据库，支持持久化的NoSQL数据库。
Redis操作基本原理示意图：

![redis1](../../images/redis1.PNG)

## 2、安装与使用

Redis 指令：http://redisdoc.com/string/getset.html

### 2.1 安装与启动命令

安装：https://blog.csdn.net/linghugoolge/article/details/86608897

使用：
客户端：D:\software_installation\redis>redis-cli.exe或 redis-cli.exe -h 127.0.0.1 -p 6379
服务端：D:\software_installation\redis>redis-server.exe redis.windows.conf


五大数据类型： String字符串、Hash哈希、List列表、Set集合、zset（sort set：有序集合）

## 3、Redis的基本使用

说明：redis安装好之后，默认有16个数据库，初始化使用0号库，编号0,...,15
1.添加key-val[set]
2.查看当前redis的所有key[key ...]
3.获取key对应的值[get key]
4.切换redis数据库[select index]
5.如何查看当前数据库的key-val数量[dbsize]
6.清空当前数据库的key-val和清空所有数据库的key-val[flushdb flushall]

redis数据类型
五大数据类型： String字符串、Hash哈希、List列表、Set集合、zset（sort set：有序集合）
命令指南：http://redisdoc.com/hash/index.html

### 3.1 string字符串

String类型可以是普通字符串，也可以是图片数据。字符串value最大值时512M。
若set (若存在，则修改)：设置；get：获取；del :删除。
基本操作：

```redis
1、Set key value：设置字符串
2、Get key ： 查找字符串
3、Dek key: 删除字符串
```

![redis2](../../images/redis2.PNG)

注意：

```redis
1、 setex(set with expire) 超时秒值：
Setex add 10 hello.go  //10秒后超时，自动删除
2、mset（muti set）[t同时设置一个或多个key-value对]
Mset key1 value1 key2 value2 ...
3、Mget
Mget key1 key2
```

![redis3](../../images/redis3.PNG)

### 3.2 Hash （哈希，类型golang的map）

Redis hash类似一个键值对集合。User map[string] string，一个string类型的filed的映射表，hash特别适合存储对象。
存储本质为字符串

![redis4](../../images/redis4.PNG)

Hash的基本操作

```redis
（1）hset/hget/hgetall /hdel
（1.1）hset key(表名) filed value
（1.2）hget key filed
（1.3）hgetall key
（1.4）hdel key 
```

![redis5](../../images/redis5.PNG)

（2）Hmset/hmget 设置多个

```redis
（2.1）Hmset key1 filed1 value1 filed2 value2  filed3 value2
（2.2）Hmget key1 filed1 filed2 filed3
```

![redis6](../../images/redis6.PNG)

注意事项：
1、`Hlen key`: 多少个键值对
2、`hexists key filed` : 1表示filed存在，0表示filed不存在

![redis7](../../images/redis7.PNG)

案例： 存放一个student信息：`stu1 name zhangsan age 30 address beijing`
![redis8](../../images/redis8.PNG)

### 3.3 list(列表)--先进先出原则

链表是简单的字符串链表，按照插入顺序排序，可以在头或尾插入数据。本质是一个链表，list的元素是有序的，元素的值是可以重复的。

List 基本操作：

```redis
（1）lpush/rpush/lrange/lpop/rpop/del
（1.1）lpush key value1 value2 value3 : 左边插入
（1.2）rpush key value1 value2 value3 : 右边插入
（1.3）lrange key start stop : 左边查看链表数据
（1.4）lpop key: 左边弹出
（1.5）rpop key：右边弹出
（1.6）del key :删除
```

![redis9](../../images/redis9.PNG)

注意事项：
1、`index`,按照索引下标获取元素（从左到右，编号从0开始）
2、`Llen key`: 返回列表key的长度，如不存在，则返回空列表0。
3、从左边插入或从右边插入
4、把所有数据移除lpop列表，对应键就删除了

![redis10](../../images/redis10.PNG)

### 3.4 set无序集合

Set是string类型的无序集合，底层是Hash table 数据结构，set也是存放很多字符串元素，字符串是无序，且元素的值不能重复。

![redis11](../../images/redis11.PNG)

Set基本使用:
（1）`sadd key value1 value2 ...`：添加set集合元素（重复添加为0）
（2）`smemebers key` :显示合元素
（3）`sismember key value1` : 判断是否是成员；若成功，则为0；失败为1
（4）`srem key value1`：删除指定值 

![redis12](../../images/redis12.PNG)

### 3.5 zset 有序集合

zset（sorted set：有序集合）：zset和set一样，都是存储string类型的集合，且都不允许重复；
但是区别是zset是为每一个元素都关联一个double类型的分数，并使用该分数对集合成员进行从小到大的排序。 
（1）添加元素操作：`zadd key score member`  //score:分数值可以是整数值或双精度浮点数。
（2）取zset元素：`zrange key score`

![redis13](../../images/redis13.PNG)

## 4、go 安装redis第三方插件

安装：go get github.com/gomodule/redigo/redis

![redis14](../../images/redis14.PNG)

Redis网络连接

```redis
package main

import (
    "fmt"
    "github.com/gomodule/redigo/redis"
)
func main(){
    //1 连接到redis: 建立tcp网络连接
    conn, err := redis.Dial("tcp","127.0.0.1:6379")
    if err != nil{
        fmt.Println("redis.Dial err = ", err)
        return
    }
    fmt.Println("conn succussfuly",conn)
}
```

## 5、 go与redis连用

### 5.1 字符串string与redis连用

```redis
package main

import (
    "fmt"
    "github.com/gomodule/redigo/redis"
)
func main(){
    //1 连接到redis: 建立tcp网络连接
    conn, err1 := redis.Dial("tcp","127.0.0.1:6379")
    if err1 != nil{
        fmt.Println("redis.Dial err = ", err1)
        return
    }
    defer conn.Close() //4关闭数据库
    //2.set通过go向redis写入数据string\
    //2.1 一个设置set
    _,err2 := conn.Do("set","name","liwen")
    if err2 != nil{
        fmt.Println("set err = ", err2)
        return
    }
    //2.2批量设置mset
    _,err22 := conn.Do("mset","name","zhangsan","id",324156, "address","beijing")
    if err22 != nil{
        fmt.Println("mset err = ", err22)
        return
    }
    //3. get 查找数据
    //3.1 一个获取get:String
    str,err3 := redis.String(conn.Do("get","name"))
    if err3 != nil{
        fmt.Println("get err = ", err3)
        return
    }
    fmt.Println(str)
    //3.2批量获取mget:Strings
    str3,err33 := redis.Strings(conn.Do("mget","name","id","address"))
    if err33 != nil{
        fmt.Println("mget err = ", err3)
        return
    }
    fmt.Println(str3)
}
```

### 5.2 哈希hash--redis

```redis
package main

import (
    "fmt"
    "github.com/gomodule/redigo/redis"
    //"strconv"
)
func main(){
    //1 连接到redis: 建立tcp网络连接
    conn, err1 := redis.Dial("tcp","127.0.0.1:6379")
    if err1 != nil{
        fmt.Println("redis.Dial err = ", err1)
        return
    }
    defer conn.Close() //4关闭数据库
    //2.哈希hash: hset通过go向redis写入数据hash
    _,err2 := conn.Do("hset","user","name","liwen")
    if err2 != nil{
        fmt.Println("hset err = ", err2)
        return
    }
    _, err22 := conn.Do("hmset","user","age",20,"id",123456,"phone",18513278676)
    if err22 != nil{
        fmt.Println("hmset err = ", err22)
        return
    }
    //3. 哈希hash: hget 查找数据
    //redis.String()
    str,err3 := redis.String(conn.Do("hget","user","name"))
    if err3 != nil{
        fmt.Println("hget err = ", err3)
        return
    }
    fmt.Println(str)
    //redis.Strings
    str3,err33 := redis.Strings(conn.Do("hmget","user","name","id","phone"))
    if err33 != nil{
        fmt.Println("hmget err = ", err33)
        return
    }
    fmt.Println(str3)
    for k,v := range str3{
        fmt.Printf("hmget key: %v, value:%v \n",k,v)
    }
}
```

### 5.3 列表list--redis

```redis
package main

import (
    "fmt"
    "github.com/gomodule/redigo/redis"
)
func main(){
    //1 连接到redis: 建立tcp网络连接
    conn, err1 := redis.Dial("tcp","127.0.0.1:6379")
    if err1 != nil{
        fmt.Println("redis.Dial err = ", err1)
        return
    }
    defer conn.Close() //4关闭数据库
    //2.push
    // //2.1 左输入数据lpush
    // if _, err := conn.Do("lpush", "listDemo", "world"); err != nil {
    //  fmt.Println(err)
    //  return
    // }
    if _, err := conn.Do("lpush", "listDemo", "hello","china","USA"); err != nil {
        fmt.Println(err)
        return
    }
    // // 1.2 linsert 左插入： 
    // if value, err := conn.Do("linsert", "listDemo", "BEFORE", "world", "---"); err != nil {
    //  fmt.Println(err)
    //  return
    // } else {
    //  fmt.Println("LINSERT=",value)
    // }
    //2.pop
    //2.1 lpop
    if value,err := conn.Do("lpop","listDemo"); err != nil{
        fmt.Println(err)
        return
    }else{
        fmt.Println("lpop value=",value)
    }
    //2.2 rpop
    if value,err := conn.Do("rpop","listDemo"); err != nil{
        fmt.Println(err)
        return
    }else{
        fmt.Println("rpop value=",value)
    }
   
    //2 显示列表数据
    //2.1 lrange
    if values, err := redis.Values(conn.Do("lrange", "listDemo", "0", "-1")); err != nil {
        fmt.Println(err)
        return
    } else {
        for _, v := range values {
            value := v.([]uint8)
            fmt.Println(string(value))
        }
    }
}
```

### 5.4 无序集合set--redis

```redis
package main

import (
    "fmt"
    "github.com/gomodule/redigo/redis"
)

func main(){
    //1 连接到redis: 建立tcp网络连接
    conn, err1 := redis.Dial("tcp","127.0.0.1:6379")
    if err1 != nil{
        fmt.Println("redis.Dial err = ", err1)
        return
    }
    defer conn.Close() //4关闭数据库
    //2无序集合 set通过go向redis写入数据string\
    //2.1 一个设置set
    _, err := conn.Do("sadd", "myset", "mobike", "foo", "ofo", "bluegogo")
    if err != nil {
        fmt.Println("set add failed", err.Error())
    }
    //查看set集合
    value, err := redis.Values(conn.Do("smembers", "myset"))
    if err != nil {
        fmt.Println("set get members failed", err.Error())
    } else {
        fmt.Printf("myset members :")
        for _,v := range value {
            fmt.Printf("%s \n", v.([]byte))
        }
        fmt.Printf("\n")
    }
    fmt.Println(value)
    
}
```

### 5.5 给数据设置有效时间

给字段Name数据设置有效时间为10s
_,err := conn.Do(“expire”,”name”,10)

### 5.6 Pipelining(管道)

管道操作可以理解为并发操作，并通过Send()，Flush()，Receive()三个方法实现。客户端可以使用send()方法一次性向服务器发送一个或多个命令，命令发送完毕时，使用flush()方法将缓冲区的命令输入一次性发送到服务器，客户端再使用Receive()方法依次按照先进先出的顺序读取所有命令操作结果。
`Send(commandName string, args ...interface{}) errorFlush() errorReceive() (reply interface{}, err error)`

Send：发送命令至缓冲区
Flush：清空缓冲区，将命令一次性发送至服务器
Recevie：依次读取服务器响应结果，当读取的命令未响应时，该操作会阻塞。
示例：

```redis
package main
import (
    "github.com/garyburd/redigo/redis"
    "fmt"
)
func main() {
    conn,err := redis.Dial("tcp","localhost:6379")
    if err != nil {
        fmt.Println("connect redis error :",err)
        return
    }
    defer conn.Close()
    //1 send发送到缓存
    
    conn.Send("HSET", "student","Score","100")
    conn.Send("HGET", "student","Score")

    conn.Send("HMSET", "student","name", "wd","age","22")
    conn.Send("HMGET", "student","name","age")
    //2一次性发送服务器
    conn.Flush()
    //3依次接收相应结果
    if res1, err := conn.Receive(); err != nil{
        fmt.Println("res1 err1 = ",err)
    }else{
        fmt.Printf("Receive res1:%v \n", res1) //Receive res1:0
    }
    if res2, err := conn.Receive(); err != nil{
        fmt.Println("res2 err1 = ",err)
    }else{
        fmt.Printf("Receive res2:%s \n", res2) //Receive res2:100
    }
    if res3, err := conn.Receive(); err != nil{
        fmt.Println("res3 err1 = ",err)
    }else{
        fmt.Printf("Receive res3: %v \n", res3) //Receive res3: OK
    }
    if res4, err := conn.Receive(); err != nil{
        fmt.Println("res4 err1 = ",err)
    }else{
        fmt.Printf("Receive res4: %s \n", res4)  //Receive res4: [wd 22]
    }
    
   
}//Receive res1:0 //Receive res2:0//Receive res3:22
```

### 5.7 发布/订阅

redis本身具有发布订阅的功能，其发布订阅功能通过命令SUBSCRIBE(订阅)／PUBLISH(发布)实现，并且发布订阅模式可以是多对多模式还可支持正则表达式，发布者可以向一个或多个频道发送消息，订阅者可订阅一个或者多个频道接受消息。
示意图：
发布者：

![redis15](../../images/redis15.PNG)

订阅者：

![redis16](../../images/redis16.PNG)

操作示例，示例中将使用两个goroutine分别担任发布者和订阅者角色进行演示：

```redis
package main
import (
    "github.com/garyburd/redigo/redis"
    "fmt"
    "time"
)

func Subs() {  //订阅者
    conn, err := redis.Dial("tcp", "localhost:6379")
    if err != nil {
        fmt.Println("connect redis error :", err)
        return
    }
    defer conn.Close()
    psc := redis.PubSubConn{conn}
    psc.Subscribe("channel1") //订阅channel1频道
    for {
        switch v := psc.Receive().(type) {
        case redis.Message:
            //channel1: message: this is wd
            fmt.Printf("%s: message: %s\n", v.Channel, v.Data) 
        case redis.Subscription:
            //channel1: subscribe 1
            fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count) 
        case error:
            fmt.Println(v)
            return
        }
    }
}
func Push(message string)  { //发布者
    conn, _ := redis.Dial("tcp", "localhost:6379")
    _,err1 := conn.Do("PUBLISH", "channel1", message)
    if err1 != nil {
        fmt.Println("pub err: ", err1)
        return
    }
}
func main()  {
    go Subs()
    go Push("this is wd")
    time.Sleep(time.Second*3)
}//channel1: subscribe 1//channel1: message: this is wd
```

### 5.8 事务操作

MULTI, EXEC,DISCARD和WATCH是构成Redis事务的基础，当然我们使用go语言对redis进行事务操作的时候本质也是使用这些命令。
MULTI：开启事务
EXEC：执行事务
DISCARD：取消事务
WATCH：监视事务中的键变化，一旦有改变则取消事务。
示例：

```redis
package main
import (
    //"github.com/garyburd/redigo/redis"
    "github.com/gomodule/redigo/redis"
    "fmt"
)

func main() {
    conn,err := redis.Dial("tcp","localhost:6379")
    if err != nil {
        fmt.Println("connect redis error :",err)
        return
    }
    defer conn.Close()
    //事务：多条语句一起执行，要么成功，要么失败
    conn.Send("MULTI") //事务启动
    conn.Send("INCR", "foo")
    conn.Send("INCR", "bar")
    // conn.Send("set", "student","liwen")
    // conn.send("set", "student","lisi")
    if r, err := conn.Do("EXEC"); err != nil{ //事务执行
        fmt.Println("err=" ,err)
        return
    }else{
        fmt.Println(r)//[1, 1]
    }
   
}
```

### 5.9 redis 连接池

Redis连接池流程：
1）事先初始化一定数量的连接，放入到连接池
2）当Go需要操作redis时，直接从Redis连接池取出即可；
3）这样可以节省时获取redis连接时间，获取获取效率

```redis
package main

import (
    "fmt"
    "github.com/gomodule/redigo/redis"
)
var pool * redis.Pool
//初始化连接池
func init(){
    pool = &redis.Pool{
        MaxIdle:8,//最大空闲连接
        MaxActive:0,//表示和数据库的最大连接数，0表示没有限制
        IdleTimeout:100, //最大空闲时间
        Dial:func() (redis.Conn,error) { //初始化链接代码，链接哪个ip的redis
            return redis.Dial("tcp","localhost:6379")
            //return  redis.Dial("tcp","127.0.0.1:6379")
        },
    }
}
func main(){
    conn := pool.Get()//从连接池获取一个连接
    defer pool.Close()//关闭一个连接池，不能从连接池取出链接
    if _,err := conn.Do("set","name","liwen");err != nil{
        fmt.Println("set err=",err)
        return
    }
    if res,err := redis.String(conn.Do("get","name")); err != nil{
        fmt.Println("get err=",err)
        return
    }else{
        fmt.Println("res=",res)
    }
    
    
}
```
