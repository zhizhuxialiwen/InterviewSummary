# gRPC微服务框架

## 1、gRPC介绍

　　gRPC是由Google公司开源的一款高性能的远程过程调用(RPC)框架，可以在任何环境下运行。该框架提供了负载均衡、跟踪、智能监控、身份验证等功能，可以实现系统间的高效连接。另外，在分布式系统中，gRPC框架也有有广泛应用，实现移动社会，浏览器等和服务器的连接。

　　gRPC开源库支持诸如：C++，C#，Dart，Go，Java，Node，Objective-C，PHP，Python，Ruby，WebJS等多种语言，开发者可以自行在gRPC的github主页库选择查看对应语言的实现。

## 2、gRPC调用执行过程

　　因为gRPC支持多种语言的实现，因此gRPC支持客户端与服务器在多种语言环境中部署运行和互相调用。多语言环境交互示例如下图所示：

![grpc9](../../images/grpc9.png)
 
 图1 多语言环境交互示例

**gRPC中默认采用的数据格式化方式是protocol buffers。**

## 3、grpc-go使用

### 3.1 定义服务

　　我们想要实现的是通过gRPC框架进行远程服务调用，首先第一步应该是要有服务。利用之前所掌握的内容，gRPC框架支持对服务的定义和生成。

　　gRPC框架默认使用`protocol buffers`作为接口定义语言，用于描述网络传输消息结构。除此之外，还可以使用`protobuf`定义服务接口。

- `OrderMessage.proto`

```go
syntax = "proto3";
package message;

// 订单请求参数
message OrderRequest {
    string orderId = 1;
    int64 timeStamp = 2;
}

// 订单信息
message OrderInfo {
    string OrderId = 1;
    string OrderName = 2;
    string OrderStatus = 3;
}

// 订单服务service定义
service OrderService {
    rpc GetOrderInfo(OrderRequest) returns (OrderInfo);
}
```

　　我们通过proto文件定义了数据结构的同时，还定义了要实现的服务接口，`GetOrderInfo`即是具体服务接口的定义，在`GetOrderInfo`接口定义中，`OrderRequest`表示是请求传递的参数，`OrderInfo`表示处理结果返回数据参数

### 3.2 编译.proto文件

　　如果定义的`.proto`文件，如本案例中所示，定义中包含了服务接口的定义，而我们想要使用gRPC框架实现RPC调用。开发者可以采用`protocol-gen-go`库提供的插件编译功能，生成兼容gRPC框架的golang语言代码。只需要在基本编译命令的基础上，指定插件的参数，告知protoc编译器即可。具体的编译生成兼容gRPC框架的服务代码的命令如下：

`protoc --go_out=plugins=grpc:. *.proto`
　
## 4、gRPC实现RPC编程

### 4.1 服务接口实现

　　在`.proto`定义好服务接口并生成对应的go语言文件后，需要对服务接口做具体的实现。定义服务接口具体由`OrderServiceImpl`进行实现，并实现`GetOrderInfo`详细内容，服务实现逻辑与前文所述内容相同。不同点是服务接口参数的变化。详细代码实现如下：

```go
package main

import (
    "context"
    "errors"
    "fmt"
    "gRPCProject/message"
    "google.golang.org/grpc"
    "net"
    "time"
)

type OrderServiceImpl struct {
}

// 具体的方法实现
func (this *OrderServiceImpl) GetOrderInfo(ctx context.Context, request *message.OrderRequest) (*message.OrderInfo, error) {
    orderMap := map[string] message.OrderInfo{
        "201907300001": message.OrderInfo{OrderId: "201907300001", OrderName: "衣服", OrderStatus: "已付款"},
        "201907310001": message.OrderInfo{OrderId: "201907310001", OrderName: "零食", OrderStatus: "已付款"},
        "201907310002": message.OrderInfo{OrderId: "201907310002", OrderName: "食品", OrderStatus: "未付款"},
    }

    var response *message.OrderInfo
    current := time.Now().Unix()
    if (request.TimeStamp > current) {
        *response = message.OrderInfo{OrderId: "0", OrderName: "", OrderStatus: "订单信息异常"}
    } else {
        result := orderMap[request.OrderId]
        if result.OrderId != "" {
            fmt.Println(result)
            return &result, nil
        } else {
            return nil, errors.New("server error")
        }
    }
    return response, nil
}

func main() {
    server := grpc.NewServer()
    message.RegisterOrderServiceServer(server, new(OrderServiceImpl))
    lis, err := net.Listen("tcp", ":8090")
    if err != nil {
        panic(err.Error())
    }
    server.Serve(lis)
}
```

### 4.2 gRPC实现客户端

　　实现完服务端以后，实现客户端程序。和服务端程序关系对应，调用gRPC框架的方法获取相应的客户端程序，并实现服务的调用，具体编程实现如下：

```go
package main

import (
    "context"
    "fmt"
    "gRPCProject/message"
    "google.golang.org/grpc"
    "time"
)

func main() {
    conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())
    if err != nil {
        panic(err.Error())
    }
    defer conn.Close()

    orderServiceClient := message.NewOrderServiceClient(conn)
    orderRequest := &message.OrderRequest{OrderId: "201907300001", TimeStamp: time.Now().Unix()}
    orderInfo, err := orderServiceClient.GetOrderInfo(context.Background(), orderRequest)
    if err == nil {
        fmt.Println(orderInfo.GetOrderId())
        fmt.Println(orderInfo.GetOrderName())
        fmt.Println(orderInfo.GetOrderStatus())
    }
}
```

运行程序:

```go
201907300001
衣服
已付款
```

## 5、gRPC调用

### 5.1、服务端流 RPC

　　在服务端流RPC实现中，服务端得到客户端请求后，处理结束返回一个数据应答流。在发送完所有的客户端请求的应答数据后，服务端的状态详情和可靠的跟踪元数据发送给客户端：

#### 5.1.1、服务接口定义

```go
//订单服务service定义
service OrderService {
    rpc GetOrderInfos (OrderRequest) returns (stream OrderInfo) {}; //服务端流模式
}
```

　　我们可以看到与之前简单模式下的数据作为服务接口的参数和返回值不同的是，此处服务接口的返回值使用了stream进行修饰。通过stream修饰的方式表示该接口调用时，服务端会以数据流的形式将数据返回给客户羰。

#### 5.1.2、编译.proto文件，生成pb.go文件

`protoc --go_out=plugins=grpc:. message.proto`

#### 5.1.3、自动生成文件的变化

　　与数据结构体发送携带数据实现不同的时，流模式下的数据发送和接收使用新的功能方法完成。在自动生成的go代码程序当中，每一个流模式对应的服务接口，都会自动生成对应的单独的client和server程序，以及对应的结构体实现。

1. 服务端自动生成

```go
type OrderService_GetOrderInfosServer interface {
    Send(*OrderInfo) error
    grpc.ServerStream
}

type orderServiceGetOrderInfosServer struct {
    grpc.ServerStream
}

func (x *orderServiceGetOrderInfosServer) Send(m *OrderInfo) error {
    return x.ServerStream.SendMsg(m)
}
```

流模式下，服务接口的服务端提供Send方法，将数据以流的形式进行发送　　　

2. 客户端自动生成

```go
type OrderService_GetOrderInfosClient interface {
    Recv() (*OrderInfo, error)
    grpc.ClientStream
}

type orderServiceGetOrderInfosClient struct {
    grpc.ClientStream
}

func (x *orderServiceGetOrderInfosClient) Recv() (*OrderInfo, error) {
    m := new(OrderInfo)
    if err := x.ClientStream.RecvMsg(m); err != nil {
        return nil, err
    }
    return m, nil
}
```

流模式下，服务接口的客户端提供Recv()方法接收服务端发送的流数据。

#### 5.1.4、服务编码实现

```go
package main

import (
    "fmt"
    "gRPCProject/demo2/message"
    "google.golang.org/grpc"
    "net"
    "time"
)

type OrderServiceImpl struct {
}

func (this *OrderServiceImpl) GetOrderInfo(request *message.OrderRequest, stream message.OrderService_GetOrderInfoServer) error {
    fmt.Println(" 服务端流RPC械 ")
    orderMap := map[string]message.OrderInfo{
        "201907300001": message.OrderInfo{OrderId: "201907300001", OrderName: "衣服", OrderStatus: "已付款"},
        "201907310001": message.OrderInfo{OrderId: "201907310001", OrderName: "零食", OrderStatus: "已付款"},
        "201907310002": message.OrderInfo{OrderId: "201907310002", OrderName: "食品", OrderStatus: "未付款"},
    }
    for id, info := range orderMap {
        if (time.Now().Unix() >= request.TimeStamp) {
            fmt.Println("订单序列号ID：", id)
            fmt.Println("订单详情：", info)
            //通过流模式发送给客户端
            stream.Send(&info)
        }
    }
    return nil
}

func main() {
    server := grpc.NewServer()
    message.RegisterOrderServiceServer(server, new(OrderServiceImpl))
    lis, err := net.Listen("tcp", ":8090")
    if err != nil {
        panic(err.Error())
    }
    server.Serve(lis)
}
```

　　`GetOrderInfos`方法就是服务接口的具体实现，因为是流模式开发，服务端将数据以流的形式进行发送，因此，该方法的第二个参数类型为OrderService_GetOrderInfosServer，该参数类型是一个接口，其中包含Send方法，允许发送流数据，Send方法的具体实现现在编译好的pb.go文件中，进一步调用grpc.SeverStream.SendMsg方法。

#### 5.1.5、客户端数据接收

　　服务端使用Send方法将数据以流的形式进行发送，客户端可以使用Recv()方法接收流数据,因为数据流是源源不断的，因此使用for无限循环实现数据流的读取，当读取到`io.EOF`时，表示流数据结束。客户端数据读取实现如下：

```go
package main

import (
    "context"
    "fmt"
    "gRPCProject/demo2/message"
    "google.golang.org/grpc"
    "io"
    "time"
)

func main() {
    conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())
    if err != nil {
        panic(err.Error())
    }
    defer conn.Close()

    orderClient := message.NewOrderServiceClient(conn)
    orderRequest := &message.OrderRequest{OrderId: "201907300001", TimeStamp: time.Now().Unix()}
    orderInfoClient, err := orderClient.GetOrderInfo(context.Background(), orderRequest)
    if err == nil {
        for {
            orderInfo, err := orderInfoClient.Recv()
            if err == io.EOF {
                fmt.Println("读取结束")
                return
            }
            if err != nil {
                panic(err.Error())
            }
            fmt.Println("读取到的信息：", orderInfo)
        }
    }
}
```

### 5.2 客户端流模式　
　
#### 5.2.1 服务接口的定义

　　与服务端同理,客户端流模式的RPC服务声明格式，就是使用stream修饰服务接口的接收参数,具体如下所示:

```go
//订单服务service定义
service OrderService {
    rpc AddOrderList (stream OrderRequest) returns (OrderInfo) {}; //客户端流模式
}
```　

#### 5.2.2 编译.proto文件

使用编译命令编译`.proto`文件。客户端流模式中也会自动生成服务接口的接口。

`protoc --go_out=plugins=grpc:. message.proto`

1. 自动生成的服务流接口实现

```go
type OrderService_AddOrderListServer interface {
    SendAndClose(*OrderInfo) error
    Recv() (*OrderRequest, error)
    grpc.ServerStream
}

type orderServiceAddOrderListServer struct {
    grpc.ServerStream
}

func (x *orderServiceAddOrderListServer) SendAndClose(m *OrderInfo) error {
    return x.ServerStream.SendMsg(m)
}

func (x *orderServiceAddOrderListServer) Recv() (*OrderRequest, error) {
    m := new(OrderRequest)
    if err := x.ServerStream.RecvMsg(m); err != nil {
        return nil, err
    }
    return m, nil
}
```

`SendAndClose`和`Recv`方法是客户端流模式下的服务端对象所拥有的方法。

2. 自动生成的客户端流接口实现

```go
type OrderService_AddOrderListClient interface {
    Send(*OrderRequest) error
    CloseAndRecv() (*OrderInfo, error)
    grpc.ClientStream
}

type orderServiceAddOrderListClient struct {
    grpc.ClientStream
}

func (x *orderServiceAddOrderListClient) Send(m *OrderRequest) error {
    return x.ClientStream.SendMsg(m)
}

func (x *orderServiceAddOrderListClient) CloseAndRecv() (*OrderInfo, error) {
    if err := x.ClientStream.CloseSend(); err != nil {
        return nil, err
    }
    m := new(OrderInfo)
    if err := x.ClientStream.RecvMsg(m); err != nil {
        return nil, err
    }
    return m, nil
}
```

`Send`和`CloseAndRecv`是客户端流模式下的客户端对象所拥有的方法。

#### 5.2.3 服务的实现

```go
package main

import (
    "WXProjectDemo/clientStream/message"
    "fmt"
    "google.golang.org/grpc"
    "io"
    "net"
)

type OrderServiceImpl struct {
}

func (this *OrderServiceImpl) AddOrderList(stream message.OrderService_AddOrderListServer) error {
    fmt.Println("客户端RPC模式")
    for {
        //从流中读取数据信息
        orderRequest, err := stream.Recv()
        if err == io.EOF {
            fmt.Println("读取数据结束")
            result := message.OrderInfo{OrderStatus: "读取数据结束"}
            return stream.SendAndClose(&result)
        }
        if err != nil {
            fmt.Println(err.Error())
            return err
        }
        fmt.Println(orderRequest)
    }
}

func main() {
    server := grpc.NewServer()
    message.RegisterOrderServiceServer(server, new(OrderServiceImpl))
    lis, err := net.Listen("tcp", ":8090")
    if err != nil {
        panic(err.Error())
    }
    server.Serve(lis)
}
```

#### 5.2.4 客户端实现

```go
package main

import (
    "WXProjectDemo/clientStream/message"
    "context"
    "fmt"
    "google.golang.org/grpc"
    "io"
)

func main() {
    conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())
    if err != nil {
        panic(err.Error())
    }
    defer conn.Close()
    orderMap := map[string]message.OrderInfo{
        "201907300001": message.OrderInfo{OrderId: "201907300001", OrderName: "衣服", OrderStatus: "已付款"},
        "201907310001": message.OrderInfo{OrderId: "201907310001", OrderName: "零食", OrderStatus: "已付款"},
        "201907310002": message.OrderInfo{OrderId: "201907310002", OrderName: "食品", OrderStatus: "未付款"},
    }

    orderServiceClient := message.NewOrderServiceClient(conn)
    addOrderListClient, err := orderServiceClient.AddOrderList(context.Background())
    if err != nil {
        panic(err.Error())
    }
    for _, info := range orderMap {
        err = addOrderListClient.Send(&info)
        if err != nil {
            panic(err.Error())
        }
    }
    for {
        orderInfo, err := addOrderListClient.CloseAndRecv()
        if err == io.EOF {
            fmt.Println(" 读取数据结束了 ")
            return
        }
        if err != nil {
            fmt.Println(err.Error())
        }
        fmt.Println(orderInfo.GetOrderStatus())
    }
}
```　

### 5.3 双向流模式　

#### 5.3.1 双向流服务的定义

```go
//订单服务service定义
service OrderService {
    rpc GetOrderInfos (stream OrderRequest) returns (stream OrderInfo) {}; //双向流模式
}
```
　　
#### 5.3.2 编译.proto文件

1. 服务端接口实现

```go
type OrderService_GetOrderInfosServer interface {
    Send(*OrderInfo) error
    Recv() (*OrderRequest, error)
    grpc.ServerStream
}

type orderServiceGetOrderInfosServer struct {
    grpc.ServerStream
}

func (x *orderServiceGetOrderInfosServer) Send(m *OrderInfo) error {
    return x.ServerStream.SendMsg(m)
}

func (x *orderServiceGetOrderInfosServer) Recv() (*OrderRequest, error) {
    m := new(OrderRequest)
    if err := x.ServerStream.RecvMsg(m); err != nil {
        return nil, err
    }
    return m, nil
}
```

2. 客户端接口实现

```go
type OrderService_GetOrderInfosClient interface {
    Send(*OrderRequest) error
    Recv() (*OrderInfo, error)
    grpc.ClientStream
}

type orderServiceGetOrderInfosClient struct {
    grpc.ClientStream
}

func (x *orderServiceGetOrderInfosClient) Send(m *OrderRequest) error {
    return x.ClientStream.SendMsg(m)
}

func (x *orderServiceGetOrderInfosClient) Recv() (*OrderInfo, error) {
    m := new(OrderInfo)
    if err := x.ClientStream.RecvMsg(m); err != nil {
        return nil, err
    }
    return m, nil
}
```

#### 5.3.3 服务实现

```go
package main

import (
    "WXProjectDemo/bothStream/message"
    "fmt"
    "google.golang.org/grpc"
    "io"
    "net"
)

type OrderServiceImpl struct {
}

func (this *OrderServiceImpl) GetOrderInfos(stream message.OrderService_GetOrderInfosServer) error {
    for {
        orderRequest, err := stream.Recv()
        if err == io.EOF {
            fmt.Println(" 数据读取结束 ")
            return err
        }
        if err != nil {
            panic(err.Error())
            return err
        }

        fmt.Println(orderRequest.GetOrderId())
        orderMap := map[string]message.OrderInfo{
            "201907300001": message.OrderInfo{OrderId: "201907300001", OrderName: "衣服", OrderStatus: "已付款"},
            "201907310001": message.OrderInfo{OrderId: "201907310001", OrderName: "零食", OrderStatus: "已付款"},
            "201907310002": message.OrderInfo{OrderId: "201907310002", OrderName: "食品", OrderStatus: "未付款"},
        }
        result := orderMap[orderRequest.GetOrderId()]
        err = stream.Send(&result)
        if err == io.EOF {
            fmt.Println(err)
            return err
        }
        if err != nil {
            fmt.Println(err.Error())
            return err
        }
    }
    return nil
}

func main() {
    server := grpc.NewServer()
    message.RegisterOrderServiceServer(server,new(OrderServiceImpl))
    lis, err := net.Listen("tcp", ":8090")
    if err != nil {
        panic(err.Error())
    }
    server.Serve(lis)
}
```

#### 5.3.4 客户端实现

```go
package main

import (
    "WXProjectDemo/bothStream/message"
    "context"
    "fmt"
    "google.golang.org/grpc"
    "io"
)

func main() {
    conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())
    if err != nil {
        panic(err.Error())
    }
    defer conn.Close()

    orderServiceClient := message.NewOrderServiceClient(conn)
    fmt.Println("客户端请求RPC调用：双向流模式")
    orderIDs := []string{"201907300001", "201907310001", "201907310002"}

    orderInfoClient, err := orderServiceClient.GetOrderInfos(context.Background())
    for _, orderID := range orderIDs {
        orderRequest := message.OrderRequest{OrderId: orderID}
        err := orderInfoClient.Send(&orderRequest)
        if err != nil {
            panic(err.Error())
        }
    }
    orderInfoClient.CloseSend()

    for {
        orderInfo, err := orderInfoClient.Recv()
        if err == io.EOF {
            fmt.Println("读取结束")
            return
        }
        if err != nil {
            return
        }
        fmt.Println("读取到的信息是： ", orderInfo)
    }
}
```

## 6、TLS验证和Token认证

- 授权认证

　　gRPC中默认支持两种授权方式,分别是：**SSL/TLS认证方式、基于Token的认证方式**。

### 6.1 SSL/TLS认证方式

　　SSL全称是Secure Sockets Layer，又被称之为安全套接字层，是一种标准安全协议，用于在通信过程中建立客户端与服务器之间的加密链接。

　　TLS的全称是Transport Layer Security，TLS是SSL的升级版。在使用的过程中，往往习惯于将SSL和TLS组合在一起写作SSL/TLS。

　　简而言之，SSL/TLS是一种用于网络通信中加密的安全协议。

#### 6.1.1 SSL/TLS工作原理

　　使用SSL/TLS协议对通信连接进行安全加密，是通过非对称加密的方式来实现的。所谓非对称加密方式又称之为公钥加密，密钥对由公钥和私钥两种密钥组成。私钥和公钥成对存在，先生成私钥，通过私钥生成对应的公钥。公钥可以公开，私钥进行妥善保存。

　　在加密过程中：客户端想要向服务器发起链接，首先会先向服务端请求要加密的公钥。获取到公钥后客户端使用公钥将信息进行加密，服务端接收到加密信息，使用私钥对信息进行解密并进行其他后续处理，完成整个信道加密并实现数据传输的过程。

#### 6.1.2 制作证书

　　可以自己在本机计算机上安装openssl，并生成相应的证书。

`openssl ecparam -genkey -name secp384r1 -out server.key`
`openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650`

#### 6.1.3 编程实现服务端

1. `message.proto`

```go
syntax = "proto3";
package message;


message RequestArgs {
    int32 args1 = 1;
    int32 args2 = 2;
}

message Response {
    int32 code = 1;
    string message = 2;
}

//服务
service MathService {
    //服务
    rpc AddMethod (RequestArgs) returns (Response) {
    };
}
```

2. 编程实现服务端

```go
package main

import (
    "context"
    "fmt"
    "gRPCProject/opensslDemo/message"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials"
    "google.golang.org/grpc/grpclog"
    "net"
    "os"
)

type MathManager struct {
}

func (this *MathManager) AddMethod(ctx context.Context, request *message.RequestArgs) (response *message.Response, err error) {
    fmt.Println(" 服务端Add方法 ")
    result := request.Args1 + request.Args2
    fmt.Println("计算的结果是： ", result)
    response = new(message.Response)
    response.Code = 1;
    response.Message = "执行成功"
    return response, nil
}

func main() {
    // TLS认证
    dir,_ := os.Getwd()
    fmt.Println("当前路径：",dir)

    creds, err := credentials.NewServerTLSFromFile("opensslDemo/keys/server.pem", "opensslDemo/keys/server.key")
    if err != nil {
        grpclog.Fatal(" 加载证书文件失败", err)
    }

    // 实例化grpc server,开启TLS认证
    server := grpc.NewServer(grpc.Creds(creds))
    message.RegisterMathServiceServer(server, new(MathManager))
    lis, err := net.Listen("tcp", ":8090")
    if err != nil {
        panic(err.Error())
    }
    server.Serve(lis)
}
```

3. 编程实现客户端

```go
package main

import (
    "context"
    "fmt"
    "gRPCProject/opensslDemo/message"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials"
    "google.golang.org/grpc/grpclog"
)

func main() {
    // TLS连接
    creds, err := credentials.NewClientTLSFromFile("opensslDemo/keys/server.pem", "go-grpc-example")
    if err != nil {
        panic(err.Error())
    }
    // 1.Dail连接
    conn, err := grpc.Dial("localhost:8090", grpc.WithTransportCredentials(creds))
    if err != nil {
        panic(err.Error())
    }
    defer conn.Close()

    serviceClient := message.NewMathServiceClient(conn)
    addArgs := message.RequestArgs{Args1: 3, Args2: 5}
    response, err := serviceClient.AddMethod(context.Background(), &addArgs)
    if err != nil {
        grpclog.Fatal(err.Error())
    }
    fmt.Println(response.GetCode(), response.GetMessage())
}
```

### 6.2 基于Token认证方式

#### 6.2.1 Token认证介绍

　　在web应用的开发过程中，我们往往还会使用另外一种认证方式进行身份验证，那就是：Token认证。基于Token的身份验证是无状态，不需要将用户信息服务存在服务器或者session中。

#### 6.2.2 Token认证过程

　　基于Token认证的身份验证主要过程是：客户端在发送请求前，首先向服务器发起请求，服务器返回一个生成的token给客户端。客户端将token保存下来，用于后续每次请求时，携带着token参数。服务端在进行处理请求之前，会首先对token进行验证，只有token验证成功了，才会处理并返回相关的数据。

#### 6.2.3 gRPC的自定义Token认证

　　在gRPC中，允许开发者自定义自己的认证规则，通过

`grpc.WithPerRPCCredentials()`

设置自定义的认证规则。`WithPerRPCCredentials` 方法接收一个 `PerRPCCredentials` 类型的参数，进一步查看可以知道`PerRPCCredentials`是一个接口，定义如下：

```go
type PerRPCCredentials interface {
    GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)
    RequireTransportSecurity() bool
}
```

因此，开发者可以实现以上接口，来定义自己的token信息。

#### 6.2.4 服务端

```go
package main

import (
    "context"
    "fmt"
    "gRPCProject/tokenDemo/message"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/credentials"
    "google.golang.org/grpc/grpclog"
    "google.golang.org/grpc/metadata"
    "google.golang.org/grpc/status"
    "net"
)

type MathManager struct {
}

func (this *MathManager) AddMethod(ctx context.Context, request *message.RequestArgs) (response *message.Response, err error) {
    // 通过metadata
    md, exist := metadata.FromIncomingContext(ctx)
    if !exist {
        return nil, status.Errorf(codes.Unauthenticated, "无Token认证信息")
    }
    var appKey string
    var appSecret string

    if key, ok := md["appid"]; ok {
        appKey = key[0]
    }

    if secret, ok := md["appkey"]; ok {
        appSecret = secret[0]
    }
    if appKey != "hello" || appSecret != "20200430" {
        return nil, status.Error(codes.Unauthenticated, "Token不合法")
    }
    fmt.Println(" 服务端 Add方法 ")
    result := request.Args1 + request.Args2
    fmt.Println(" 计算结果是：", result)
    response = new(message.Response)
    response.Code = 1;
    response.Message = "执行成功"
    return response, nil
}

func main() {
    creds, err := credentials.NewServerTLSFromFile("tokenDemo/keys/server.pem", "tokenDemo/keys/server.key")
    if err != nil {
        grpclog.Fatal("加载证书文件失败", err)
    }
    server := grpc.NewServer(grpc.Creds(creds))
    message.RegisterMathServiceServer(server, new(MathManager))
    lis, err := net.Listen("tcp", ":8090")
    if err != nil {
        panic(err.Error())
    }
    server.Serve(lis)
}
```

#### 6.2.5 客户端

```go
package main

import (
    "context"
    "fmt"
    "gRPCProject/tokenDemo/message"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials"
    "google.golang.org/grpc/grpclog"
)

//token认证
type TokenAuthentication struct {
    AppKey    string
    AppSecret string
}

//组织token认证的metadata信息
func (this *TokenAuthentication) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
    return map[string]string{
        "appid":  this.AppKey,
        "appkey": this.AppSecret,
    }, nil
}

//是否基于TLS认证进行安全传输
func (this *TokenAuthentication) RequireTransportSecurity() bool {
    return true
}

func main() {
    creds, err := credentials.NewClientTLSFromFile("tokenDemo/keys/server.pem", "go-grpc-example")
    if err != nil {
        panic(err.Error())
    }
    auth := TokenAuthentication{
        AppKey:    "hello",
        AppSecret: "20200430",
    }
    //1、Dail连接
    conn, err := grpc.Dial("localhost:8090", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&auth))
    if err != nil {
        panic(err.Error())
    }
    defer conn.Close()

    serviceClient := message.NewMathServiceClient(conn)
    addArgs := message.RequestArgs{Args1: 89, Args2: 566}
    response, err := serviceClient.AddMethod(context.Background(), &addArgs)
    if err != nil {
        grpclog.Fatal(err.Error())
    }
    fmt.Println(response.GetCode(), response.GetMessage())
}
```

#### 6.2.6 拦截器的使用

　　在服务端的方法中，每个方法都要进行token的判断。程序效率太低，可以优化一下处理逻辑，在调用服务端的具体方法之前，先进行拦截，并进行token验证判断，这种方式称之为拦截器处理。

　　除了此处的token验证判断处理以外，还可以进行日志处理等。

1. `Interceptor`

使用拦截器，首先需要注册。
在grpc中编程实现中，可以在NewSever时添加拦截器设置，grpc框架中可以通过`UnaryInterceptor`方法设置自定义的拦截器，并返回ServerOption。具体代码如下：

```go
grpc.UnaryInterceptor()
```
　　
`UnaryInterceptor()`接收一个`UnaryServerInterceptor`类型，继续查看源码定义，可以发现`UnaryServerInterceptor`是一个func，定义如下：

```go
type UnaryServerInterceptor func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler UnaryHandler) (resp interface{}, err error)
```

通过以上代码，如果开发者需要注册自定义拦截器，需要自定义实现`UnaryServerInterceptor`的定义。

2. 自定义`UnaryServerInterceptor`

　　接下来就自定义实现func,符合UnaryServerInterceptor的标准，在该func的定义中实现对token的验证逻辑。自定义实现func如下：

```go
func TokenInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

    //通过metadata
    md, exist := metadata.FromIncomingContext(ctx)
    if !exist {
        return nil, status.Errorf(codes.Unauthenticated, "无Token认证信息")
    }

    var appKey string
    var appSecret string
    if key, ok := md["appid"]; ok {
        appKey = key[0]
    }
    if secret, ok := md["appkey"]; ok {
        appSecret = secret[0]
    }

    if appKey != "hello" || appSecret != "20190812" {
        return nil, status.Errorf(codes.Unauthenticated, "Token 不合法")
    }
    //通过token验证，继续处理请求
    return handler(ctx, req)
}
```

　　在自定义的`TokenInterceptor`方法定义中,和之前在服务的方法调用的验证逻辑一致，从`metadata`中取出请求头中携带的`token`认证信息，并进行验证是否正确。如果`token`验证通过，则继续处理请求后续逻辑，后续继续处理可以由`grpc.UnaryHandler`进行处理。`grpc.UnaryHandler`同样是一个方法，其具体的实现就是开发者自定义实现的服务方法。`grpc.UnaryHandler`接口定义源码定义如下：

```go
type UnaryHandler func(ctx context.Context, req interface{}) (interface{}, error)
```

3. 拦截器注册

在服务端调用`grpc.NewServer`时进行拦截器的注册。详细如下：

```go
server := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(TokenInterceptor))
```

- server.go

```go
package main

import (
    "WXProjectDemo/InterceptorDemo/message"
    "context"
    "fmt"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/credentials"
    "google.golang.org/grpc/grpclog"
    "google.golang.org/grpc/metadata"
    "google.golang.org/grpc/status"
    "net"
)

type MathManager struct {
}

func (this *MathManager) AddMethod(ctx context.Context, request *message.RequestArgs) (response *message.Response, err error) {
    fmt.Println("服务端Add方法")
    result := request.Args1 + request.Args2
    fmt.Println(" 计算结果是： ", result)
    response = new(message.Response)
    response.Code = 1
    response.Message = "执行成功"
    return response, nil
}

func main() {
    // TLS认证
    creds, err := credentials.NewServerTLSFromFile("InterceptorDemo/keys/server.pem", "InterceptorDemo/keys/server.key")
    if err != nil {
        grpclog.Fatal("加载证书文件失败", err)
    }
    // 实例化grpc server,开启TLS认证
    server := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(TokenInterceptor))
    message.RegisterMathServiceServer(server, new(MathManager))
    lis, err := net.Listen("tcp", ":8090")
    if err != nil {
        panic(err.Error())
    }
    server.Serve(lis)
}
```

- 自定义拦截器实现

```go
func TokenInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
    // 通过metadata
    md, exist := metadata.FromIncomingContext(ctx)
    if !exist {
        return nil, status.Errorf(codes.Unauthenticated, "无Token认证信息")
    }
    var appKey string
    var appSecret string
    if key, ok := md["appid"]; ok {
        appKey = key[0]
    }
    if secret, ok := md["appkey"]; ok {
        appSecret = secret[0]
    }
    if appKey != "hello" || appSecret != "20200502" {
        return nil, status.Errorf(codes.Unauthenticated, "Token 不合法")
    }
    // 通过token验证，继续处理请求
    return handler(ctx, req)
}
```

- client.go

```go
package main

import (
    "WXProjectDemo/InterceptorDemo/message"
    "context"
    "fmt"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials"
    "google.golang.org/grpc/grpclog"
)

// token认证
type TokenAuthentication struct {
    AppKey    string
    AppSecret string
}

// 组织token认证的metadata信息
func (this *TokenAuthentication) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
    return map[string]string{
        "appid":  this.AppKey,
        "appkey": this.AppSecret,
    }, nil
}

// 是否基于TLS认证进行安全传输
func (this *TokenAuthentication) RequireTransportSecurity() bool {
    return true
}

func main() {
    // TLS连接
    creds, err := credentials.NewClientTLSFromFile("InterceptorDemo/keys/server.pem", "go-grpc-example")
    if err != nil {
        panic(err.Error())
    }
    auth := TokenAuthentication{
        AppKey:    "hello",
        AppSecret: "20200502",
    }
    // 1.Dail连接
    conn, err := grpc.Dial("localhost:8090", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&auth))
    if err != nil {
        panic(err.Error())
    }
    defer conn.Close()

    serviceClent := message.NewMathServiceClient(conn)
    addArgs := message.RequestArgs{Args1: 3, Args2: 97}
    response, err := serviceClent.AddMethod(context.Background(), &addArgs)
    if err != nil {
        grpclog.Fatal(err.Error())
    }
    fmt.Println(response.GetCode(), response.GetMessage())
}
```

- 项目运行

　　依次运行server.go程序和client.go程序，可以得到程序运行的正确结果。修改token携带值，可以验证token非法情况的拦截器效果。