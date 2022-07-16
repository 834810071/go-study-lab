# Go单元测试从入门到放弃—0.单元测试基础
```shell
# 在执行go test命令的时候可以添加-run参数，它对应一个正则表达式，只有函数名匹配上的测试函数才会被go test命令执行
go test -run=Sep -v
```

```go
func TestTimeConsuming(t *testing.T) {
    if testing.Short() {
        t.Skip("short模式下会跳过该测试用例")
    }
    ...
}
// 当执行go test -short时就不会执行上面的TestTimeConsuming测试用例。

// Go1.7+ 中新增了子测试，支持在测试函数中使用t.Run执行一组测试用例，这样就不需要为不同的测试数据定义多个测试函数了。
func TestXXX(t *testing.T) {
    t.Run("case1", func(t *testing.T){...})
    t.Run("case2", func(t *testing.T){...})
    t.Run("case3", func(t *testing.T){...})
}
```

```shell
# 社区里有很多自动生成表格驱动测试函数的工具，比如gotests等
go get -u github.com/cweill/gotests/...
gotests -all -w split.go
```

```shell
# 查看测试覆盖率
go test -cover
```

```shell
# Go还提供了一个额外的-coverprofile参数，用来将覆盖率相关的记录信息输出到一个文件
go test -cover -coverprofile=c.out

# 使用cover工具来处理生成的记录信息，该命令会打开本地的浏览器窗口生成一个HTML报告
go tool cover -html=c.out
```

```shell
# testify是一个社区非常流行的Go单元测试工具包，其中使用最多的功能就是它提供的断言工具——testify/assert或testify/require
go get github.com/stretchr/testify
```

# Go单元测试--模拟服务请求和接口返回
```shell
go get -u gopkg.in/h2non/gock.v1
```

# Go单测测试 — 数据库 CRUD 的 Mock 测试
```shell
# sqlmock 是一个实现 sql/driver 的mock库。
go get github.com/DATA-DOG/go-sqlmock
# miniredis是一个纯go实现的用于单元测试的redis server
go get github.com/alicebob/miniredis/v2
```

# 在项目里怎么给 GORM 做单元测试


