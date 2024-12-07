# Go build and install

## Build

注意创建项目时，要有：
```bash
go mod init booking-app
```

编译命令如下：

```bash
cd booking-app
go build
```

你将会看到一个和包名相同的可执行文件: booking-app, 通过下面命令你可执行这个文件:

```bash
./booking-app
```

## Install

### 可执行文件的路径

添加一个我们自定义的go可执行文件的安装地方

```bash
go env -w GOBIN=/Users/lbb/workspace/software/mygo  
```

将自定义Go可执行文件的安装路径，添加到你对应的shell的路径下, macos可以选择 ~/.zshrc，添加如下内容

```
export GOBIN="/Users/lbb/workspace/software/mygo"
export PATH=$GOBIN:$PATH
```

然后记得: source  ~/.zshrc 

### install go program

```bash
cd booking-app
go install
```

即可在$GOBIN看到可执行文件 booking-app. 同时在shell中可以使用booking-app命令，执行booking-app程序.
