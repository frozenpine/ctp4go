# CTP 交易&行情接口封装

> 采用动态导入方式进行封装
> 
> 使用C结构体实现了Api&Spi的虚表结构，无需Cpp封装层再导出C接口的传统模式，减少了封装层级，仅有回调虚表函数到Cgo导出函数的一层薄封装。

## 代码组织结构

- **dependencies** 存放SDK头文件及库文件
  
  > 目录内按CTP系统类型 -> 版本号组织
  
  - *future* 期货期权交易系统SDK，
  
  - *mini* Mini交易系统SDK（待扩展）
  
  - *etf* 股票期权交易系统SDK（待扩展）

- **gen** CTP SDK头文件解析模块
  
  > 解析CTP的 *ThostFtdcTraderApi.h* 和 *ThostFtdcMdApi.h* 头文件，并生成对应版本SDK的 `go` 封装代码和 `c` 桥接代码
  > 
  > 封装目前仅支持 **动态加载模式**
  
  - **internal** 基于 `go-clang` 的 `cpp` 头文件解析模块
    
    > 使用 `github.com/go-clang/clang-v15` 版本，需要代码生成环境具备 `llvm` , `clang` 的头文件及对应动态库（ `clang` 15.0.0版本及以上），同时模块采用 `pkg-config` 查找 `clang` 的头文件及链接路径，需具备 `llvm.pc` 配置文件，如不存在对应配置文件，可在 `pkg-config` 的配置路径内生成一个
    > 
    > ```bash
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

- **thost** CTP的接口数据类型/结构体封装，Go API/SPI接口签名，公共回调实现及版本注册入口点

- **trader** CTP各个版本的具体封装代码及交易接口的API + SPI的完整实现
  
  > ***imp_xxx.go*** 以tag条件编译方式导入具体的版本包，通过包内的 `init` 函数向 `thost` 注册具体版本的API创建实例，如使用 `trader` 模块的抽象，则编译时需使用 `-tag ${具体版本号}` 指定导入的实现版本
  
  - *v6_7_13* v6.7.13的具体封装实现，可脱离上层抽象，直接导入使用
  
  - *其他版本* 待更新实现

- **mduser** CTP各个版本的具体封装代码及行情接口的API + SPI的完整实现
  
  > ***imp_xxx.go*** 以tag条件编译方式导入具体的版本包，通过包内的 `init` 函数向 `thost` 注册具体版本的API创建实例，如使用 `mduser` 模块的抽象，则编译时需使用 `-tag ${具体版本号}` 指定导入的实现版本
  
  - *v6_7_13* v6.7.13的具体封装实现，可脱离上层抽象，直接导入使用
  
  - *其他版本* 待更新实现

- **state** 状态管理模块