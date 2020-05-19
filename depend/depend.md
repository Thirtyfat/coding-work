### go语言的依赖管理
* 依赖管理的三个阶段：GOPATH,GOVENDOR,go mod
* GOPATH 每个项目建立自己的vendor目录，存放第三方库文件 (glide，dep)


#### go mod
> 拉取第三方库(最新)
```go
go get -u go.uber.org/zap 
```
> 切换版本(指定版本)
```go
go get -u go.uber.org/zap@v1.12
```
> 清理历史版本
```go
go mod tidy
```
> 编译
```go
go build ./...
```
* 由go命令进行统一管理，用户不必关心目录结构

##### 项目迁移
* go mod : go mod init,go build ./...

### 目录
* 一个package下面只能存在一个main函数
```go
go build ./...
go install ./...
```