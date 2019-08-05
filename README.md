# awesomePet
<code>golang 1.11+</code> &nbsp; <code>echo V4</code> &nbsp; <code>proto3</code> &nbsp; <code>mysql 5.x+</code>

一个使用 Golang 编写的后端服务 学习案例——在线宠物分享

## Use
#### Requirements
Go version >= 1.11 and GO111MODULE=on； mysql >= 8.0；配置 conf.yaml 文件
#### Build & Run
```bash
$ git clone https://github.com/
$ cd Wade && go build
$ export GOPROXY=https://goproxy.io //存在网络环境问题的可以设置代理
$ chmod a+x Wade //linux下赋予文件执行权限
$ ./Wade
```
打开浏览器访问：[https://localhost:443/info] 查看 http 请求信息

## Features
* 跨平台开发与交叉编译（启用CGO需关闭交叉编译）；
* 远程程序调用,跨平台、跨语言进程间通信；
* 数据可视化库，后台编译生成html；
#### awesomePet有哪些功能？
* [Echo](https://echo.labstack.com/)：高性能、可扩展、简约的的Go Web框架;
* [GORM](https://gorm.io/)：全功能数据库 orm 引擎；
* [go-echarts](https://go-echarts.chenjiandongx.com/)：Golang 数据可视化第三方库；
* [gRPC](https://grpc.io/)：高性能、开源的RPC框架，可跨语言；
* [Protocol Buffers](https://github.com/protocolbuffers/protobuf)：配合gRPC使用；

## others
#### About Author——追求 源于热爱
欢迎反馈使用过程中遇到的问题，可用以下联系方式跟我交流：
* 邮箱：beihai@wingsxdu.com；
* QQ：1844848686；
* blog：<https://www.wingsxdu.com> @beihai

#### Noticed
* TLS 使用自签名证书，有效期为一年，更新证书指令：
```bash
$ cd Wade && go run $GOROOT/src/crypto/tls/generate_cert.go --host localhost
```
* IDE: Goland ,为方便作者在多设备上编程，.idea 文件夹一并上传到 github 上，可自行删除。

#### why awesomePet?
* 只是想做一个golang 后端开发的案例，并设计(脑部）了一个应用场景——在线宠物分享
* 搬砖民工也会建成自己的罗马帝国。

#### Thanks