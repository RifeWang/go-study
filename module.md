### 包管理 -- 使用 go.module 摆脱对 GOPATH 的依赖。
### https://github.com/golang/go/wiki/Modules

module 包含多个 package , go.mod 记录 module path, import path, dependency requirements.

对于不同的 major version , module path 必须指明。

包引用：`import "<module_name>/<package_name>"`

第三方包被下载到了 `$GOPATH/pkg/mod` 路径下。
- `go mod init <module_name>` : 初始化模块，一般将整个项目视为一个 module 。
- `go list -m all` : 列出所有依赖包。
- `go get golang.org/x/text@version` : 安装依赖包会自动记录并更新 go.mod, go.sum 文件。
- `import "rsc.io/quote/v3"` : module path 需要声明 major version 。
- `go mod tidy` : 清理依赖包。
- `go mod vendor` : 将所有依赖包下载到 vendor 目录下。
- `go build -mod=vendor` : 会使用本地顶级目录下的 vendor 下的依赖包。

如果被墙了，设置 GOPROXY 环境变量使用代理即可 `export GOPROXY=https://goproxy.io` .


### go build

`go build -o <output> <input>`: `go build -o bin/mybinfile os/os.go` 以当前路径为基准，编译 os 目录下的 os.go ，输出目录为 bin , 可执行文件为 m 。

交叉编译：通过环境变量 GOOS , GOARCH 分别指定目标环境的操作系统及架构。https://golang.org/doc/install/source#environment 。
`GOOS=linux GOARCH=amd64 go build -o bin/mybinfile os/os.go`

#### package names
包名应该简单明了，全小写，不带下划线，不允许大小写混合。
熟悉的部分可以缩写，但如果缩写造成表达的意义不明确则应该避免这样做。
不要从用户那里窃取名称，buffered I/O package is called bufio, not buf。

包名称与其内容是耦合的。
避免口吃，The HTTP server provided by the http package is called Server, not HTTPServer. 
简化函数名称，pkg 包中的函数返回的数据类型是 pkg.Pkg (或者 *pkg.Pkg) ，函数名称通常可以省略类型名称而不会产生混淆。
不同包中的类型可以同名，因为通过包就可以区分。

package paths，导入路径的最后一个元素是 package name ，通过 package 关键字定义的。
目录只提供了排列方式，包含于其中的多个 packages 之间并没有实际关联。
不同目录的包名也可以相同。

不要将所有的 APIs 放到一个 package 中，避免不受限制地增长，不向用户提供指导，累积依赖关系以及其他导入冲突。
避免不必要的同名。

