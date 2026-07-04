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

1. 进入模块目录，使用 `make` 命令生成 **SDK** 相关版本的代码文件
   
   > `make` 支持的参数：
   > 
   > - VERSION=版本号（默认：v6.7.13）
   > 
   > - PLATFORM=平台名称（默认：future）
   > 
   > 

2. 命令执行后，将生成三种类型代码文件：
   
   - 数据定义代码：
     
     - ***thost*** 下生成结构体定义文件：*ctp_structs.go*
     
     - ***thost/types*** 下生成类型定义文件：*ctp_types.go*
       
       > 可手工执行类型定义的
   
   - 接口封装代码：
     
     > 封装代码生成在对应接口名模块文件夹下，可直接导入需要的版本实现使用
     > 
     > 以下以 `trader` 交易模块的 *v6.7.13* 版本为例
     
     - ***trader/v6.7.13*** 下生成对应版本接口封装代码：
       
       - *gen.go* 全部封装代码的 go generate 定义
         
         > 该定义文件存在的情况下，可重复手工执行生成代码：
         > 
         > ```bash
         > go generate .
         > ```
       
       - *api_helper.h* api 接口的 c 桥接代码定义
       
       - *api_helper.c* api 接口的 c 桥接代码实现
       
       - *api_impl.go* api 接口的 go 封装实现
       
       - *spi_helper.h* spi 接口的 c 桥接代码定义
       
       - *spi_helper.c* spi 接口的 c 桥接代码实现
       
       - *spi_impl.go* spi 接口的 go 封装实现
       
       - *consts_linux.go* linux系统特定的静态函数名
       
       - *consts_windows.go* windows系统特定的静态函数名
     
     - ***trader*** 下生成版本导入文件： *imp_6.7.13.go* 

3. 如需使用 `trader` 或 `mduser` 模块下更高抽象层级的接口，需在对应版本封装模块内实现 `init()` 初始化调用，完成 `thost` 的版本化模块注册。 
   
   > 目前由于 `TraderApi` 和 `MdApi` 的接口定义由于不同版本存在差异，接口的定义暂未做自动化生成，故而封装模块的初始化注册代码也未通过自动化生成
   
   ```go
   package v6_7_13
   
   import (
       "fmt"
   
       "github.com/frozenpine/ctp4go/thost"
       "github.com/frozenpine/ctp4go/thost/types"
   )
   
   type apiWrapper struct {
       *ThostFtdcTraderApi
   }
   
   func (api apiWrapper) SubscribePrivateTopic(
       nResumeType types.THOST_TE_RESUME_TYPE, nSeqNo ...int,
   ) {
       var seq int
       if len(nSeqNo) > 0 {
           seq = nSeqNo[0]
       }
       api.ThostFtdcTraderApi.SubscribePrivateTopic(
           int(nResumeType), seq,
       )
   }
   
   func (api apiWrapper) SubscribePublicTopic(
       nResumeType types.THOST_TE_RESUME_TYPE,
   ) {
       api.ThostFtdcTraderApi.SubscribePublicTopic(int(nResumeType))
   }
   
   func sdkMaker(
       libPath string,
       params ...thost.Param,
   ) func() (thost.TraderApi, error) {
       return func() (thost.TraderApi, error) {
           if libPath == "" {
               return nil, fmt.Errorf(
                   "%w: lib path is empty", thost.ErrInvalidArgs,
               )
           }
   
           var (
               FlowPath         string
               IsProductinoMode bool
   
               ok bool
           )
   
           for _, p := range params {
               switch p.Key {
               case thost.ParamFlowPath:
                   if FlowPath, ok = p.Value.(string); !ok {
                       return nil, fmt.Errorf(
                           "%w: invalid %s value %+v",
                           thost.ErrInvalidArgs, p.Key, p.Value,
                       )
                   }
               case thost.ParamIsProductionMode:
                   if IsProductinoMode, ok = p.Value.(bool); !ok {
                       return nil, fmt.Errorf(
                           "%w: invalid %s value %+v",
                           thost.ErrInvalidArgs, p.Key, p.Value,
                       )
                   }
               }
           }
   
           api, err := CreateThostFtdcTraderApi(
               libPath, FlowPath, IsProductinoMode,
           )
           if err != nil {
               return nil, err
           }
   
           return apiWrapper{api}, nil
       }
   }
   
   func init() {
       if err := thost.SetTraderMaker("v6.7.13", sdkMaker); err != nil {
           panic(err)
       }
   }
   ```
   
   ```go
   package v6_7_13
   
   import (
       "fmt"
   
       "github.com/frozenpine/ctp4go/thost"
   )
   
   func sdkMaker(
       libPath string,
       params ...thost.Param,
   ) func() (thost.MdApi, error) {
       return func() (thost.MdApi, error) {
           if libPath == "" {
               return nil, fmt.Errorf(
                   "%w: lib path is empty", thost.ErrInvalidArgs,
               )
           }
   
           var (
               FlowPath         string
               IsUsingUdp       bool
               IsMulticast      bool
               IsProductinoMode bool
   
               ok bool
           )
   
           for _, p := range params {
               switch p.Key {
               case thost.ParamFlowPath:
                   if FlowPath, ok = p.Value.(string); !ok {
                       return nil, fmt.Errorf(
                           "%w: invalid %s value %+v",
                           thost.ErrInvalidArgs, p.Key, p.Value,
                       )
                   }
               case thost.ParamIsProductionMode:
                   if IsProductinoMode, ok = p.Value.(bool); !ok {
                       return nil, fmt.Errorf(
                           "%w: invalid %s value %+v",
                           thost.ErrInvalidArgs, p.Key, p.Value,
                       )
                   }
               }
           }
   
           return CreateThostFtdcMdApi(
               libPath, FlowPath, IsUsingUdp, IsMulticast, IsProductinoMode,
           )
       }
   }
   
   func init() {
       if err := thost.SetMduserMaker("v6.7.13", sdkMaker); err != nil {
           panic(err)
       }
   }
   ```