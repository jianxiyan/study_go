# go 编译

#### go build

    go build表示将源代码编译成可执行文件。
    在hello目录下执行：

    go build
    或者在其他目录执行以下命令：

    go build hello
    go编译器会去 GOPATH的src目录下查找你要编译的hello项目

    编译得到的可执行文件会保存在执行编译命令的当前目录下，如果是windows平台会在当前目录下找到hello.exe可执行文件。

    可在终端直接执行该hello.exe文件：

    c:\desktop\hello>hello.exe
    Hello World!
    我们还可以使用-o参数来指定编译后得到的可执行文件的名字。

    go build -o heiheihei.exe

#### go install

    go install表示安装的意思，它先编译源代码得到可执行文件，然后将可执行文件
    移动到GOPATH的bin目录下。因为我们的环境变量中配置了GOPATH下的bin目录，
    所以我们就可以在任意地方直接执行可执行文件了。

#### 跨平台编译

    默认我们go build的可执行文件都是当前操作系统可执行的文件，如果我想在windows下编译一个linux下可执行文件，那需要怎么做呢？

    只需要指定目标操作系统的平台和处理器架构即可：

    SET CGO_ENABLED=0  // 禁用CGO
    SET GOOS=linux  // 目标平台是linux
    SET GOARCH=amd64  // 目标处理器架构是amd64
    使用了cgo的代码是不支持跨平台编译的

    然后再执行go build命令，得到的就是能够在Linux平台运行的可执行文件了。

    Mac 下编译 Linux 和 Windows平台 64位 可执行程序：

    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
    Linux 下编译 Mac 和 Windows 平台64位可执行程序：

    CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
    Windows下编译Mac平台64位可执行程序：

    SET CGO_ENABLED=0
    SET GOOS=darwin
    SET GOARCH=amd64
    go build

# 变量常量

# 基本数据类型

    Go语言中有丰富的数据类型，除了基本的整型、浮点型、布尔型、字符串外，还有数组、切片、结构体、函数、map、通道（channel）等。Go 语言的基本类型和其他语言大同小异。

#### 整型

整型分为以下两个大类： 按长度分为：int8、int16、int32、int64 对应的无符号整型：uint8、uint16、uint32、uint64
其中，uint8 就是我们熟知的 byte 型，int16 对应 C 语言中的 short 型，int64 对应 C 语言中的 long 型。

    类型	描述
    uint8	无符号 8位整型 (0 到 255)
    uint16	无符号 16位整型 (0 到 65535)
    uint32	无符号 32位整型 (0 到 4294967295)
    uint64	无符号 64位整型 (0 到 18446744073709551615)
    int8	有符号 8位整型 (-128 到 127)
    int16	有符号 16位整型 (-32768 到 32767)
    int32	有符号 32位整型 (-2147483648 到 2147483647)
    int64	有符号 64位整型 (-9223372036854775808 到 9223372036854775807)
