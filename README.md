# CTP 交易&行情接口封装

> 采用动态导入方式进行封装
> 
> 使用C结构体实现了Api&Spi的虚表结构，无需Cpp封装层再导出C接口的传统模式，减少了封装层级，仅有回调虚表函数到Cgo导出函数的一层薄封装。

## 代码组织结构

- **dependencies** 存放SDK头文件及库文件
  
  > 目录内按CTP系统类型 -> 版本号组织
  
  - *future* 期货期权交易系统SDK，
  
  - *mini* Mini交易系统SDK
  
  - *etf* 股票期权交易系统SDK（待扩展）

- **gen** CTP SDK头文件解析模块
  
  > 解析CTP的 *ThostFtdcTraderApi.h* 和 *ThostFtdcMdApi.h* 头文件，并生成对应版本SDK的 `go` 封装代码和 `c` 桥接代码
  > 
  > 封装目前仅支持 **动态加载模式**
  
  - **parser** 基于 `go-clang` 的 `cpp` 头文件解析模块
    
    > 使用 `github.com/go-clang/clang-v15` 版本，需要代码生成环境具备 `llvm` , `clang` 的头文件及对应动态库（ `clang` 15.0.0版本及以上），同时模块采用 `pkg-config` 查找 `clang` 的头文件及链接路径，需具备 `llvm.pc` 配置文件，如不存在对应配置文件，可在 `pkg-config` 的配置路径内生成一个
    > 
    > ```bash
    > # 安装 llvm clang，其他系统环境请自行查找安装命令
    > apt update -y
    > apt install -y llvm clang libclang-dev
    > 
    > # 生成 pkg-config 配置文件
    > cat <<EOF > llvm.pc
    > Name: LLVM
    > Description: Low Level Virtual Machine
    > Version: $(llvm-config --version)
    > Cflags: -I$(llvm-config --includedir)
    > Libs: -L$(llvm-config --libdir) $(llvm-config --system-libs)
    > EOF
    > 
    > # 查看 pkg-config 配置内容
    > pkg-config --libs --cflags llvm
    > # -I/usr/lib/llvm-18/include -L/usr/lib/llvm-18/lib
    > ```
  
  - **handlers** `clang` 解析后的AST对应的 `c` `go` `cgo` 的参数名及类型名处理
    
    > 目前的处理逻辑特化于 CTP 各个接口的参数类型，理论上可以扩展为 `cpp` 支持的所有参数类型
  
  - **templates** 基于 `text/template` 模块的代码生成模板文件

- **thost** CTP的接口数据类型/结构体封装，Go API/SPI接口签名，公共回调实现及版本注册入口点

- **trader** CTP各个版本的具体封装代码及交易接口的API + SPI的完整实现
  
  > ***imp_xxx.go*** 以tag条件编译方式导入具体的版本包，通过包内的 `init` 函数向 `thost` 注册具体版本的API创建实例，如使用 `trader` 模块的抽象，则编译时需使用 `-tag ${具体版本号}` 指定导入的实现版本
  
  - *v6.7.13* v6.7.13的具体封装实现，可脱离上层抽象，直接导入使用
  
  - *其他版本* 待更新实现

- **mduser** CTP各个版本的具体封装代码及行情接口的API + SPI的完整实现
  
  > ***imp_xxx.go*** 以tag条件编译方式导入具体的版本包，通过包内的 `init` 函数向 `thost` 注册具体版本的API创建实例，如使用 `mduser` 模块的抽象，则编译时需使用 `-tag ${具体版本号}` 指定导入的实现版本
  
  - *v6.7.13* v6.7.13的具体封装实现，可脱离上层抽象，直接导入使用
  
  - *其他版本* 待更新实现

- **state** 状态管理模块

## 代码生成

工程根目录可执行 `make` 命令生成对应模块，也可进入对应模块目录下执行 `make`

> ```bash
> # 工程根目录
> make trader
> 
> # trader 模块目录下
> cd trader
> make
> 
> # 以上两种执行方式等价
> ```

1. `thost` 数据定义代码：
   
   > ```bash
   > # 生成数据定义代码
   > make thost
   > 
   > # 生成v6.5.1版的模块代码，trader & mduser 同理
   > # 需在 dependencies 对应的平台目录下存在对应版本的头文件定义
   > make VERSION=v6.5.1 thost
   > ```
   
   - ***thost/${平台名}*** 下生成结构体定义文件：*ctp_structs.go*
   
   - ***thost/${平台名}/types*** 下生成类型定义文件：*ctp_types.go

2. `trader` | `mduser` 接口封装代码 `make trader` | `make mduser` ：
   
   > ```bash
   > # 生成 v1.7.5 版本的 mini 交易接口
   > make PLATFORM=mini VERSION=v1.7.5 trader
   > 
   > # 生成 v6.5.1 版本的 future 行情接口
   > make VERSION=v6.5.1 mduser 
   > ```
   > 
   > 封装代码生成在对应版本号文件夹下，可直接导入需要的版本实现使用
   
   - ***trader|mduser/${平台名}/v6.7.13*** 包含对应版本接口封装代码
     
     - *api_helper.h* api 接口的 c 桥接代码定义
     
     - *api_helper.c* api 接口的 c 桥接代码实现
     
     - *api_impl.go* api 接口的 go 封装实现
     
     - *spi_helper.h* spi 接口的 c 桥接代码定义
     
     - *spi_helper.c* spi 接口的 c 桥接代码实现
     
     - *spi_impl.go* spi 接口的 go 封装实现
     
     - *consts_linux.go* linux系统特定的静态函数名
     
     - *consts_windows.go* windows系统特定的静态函数名
   
   - ***trader|mduser/${平台名}/imp_v6.7.13.go*** `trader` | `mduser` 模块版本特化的导入代码，编译时需指定 `-tags v6.7.13` 参数

3. `clean` 清除全部生成代码
   
   > 可在 `clean` 后加连字符 **-** 跟 **target** 名（无空格分隔），清理指定 **target** 的生成代码
   > 
   > ```bash
   > # 清除全部指定平台版本的模块代码
   > make clean
   > 
   > # 清除指定平台版本的 thost 代码
   > make clean-thost
   > 
   > # 清除 mini 版本 v1.7.5 全部代码
   > make PLATFORM=mini VERSION=v1.7.5 clean
   > ```