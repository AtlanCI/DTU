# DTU
微诺时代物联网项目框架参考微服务DDD和kratos工程设计概念
# 目录结构
```bash
├── apis   //基于porotocl buffer设计的api http或grpc
│   └── 1.txt
├── cmd   // 程序生命周期管理 启动，停止等 使用依赖注入
│   ├── 1.txt
│   └── DTU //app名称
├── configs //配置文件
│   └── 1.txt
├── go.mod
├── internal  //各服务 
│   ├── devices_health   //app名称
│   │   ├── biz  //业务逻辑层，data需要的数据接口定义在这里
│   │   ├── data //数据层  依赖倒置原则 实现业务的数据接口
│   │   └── service //简单业务逻辑组装，实现apis目录定义的接口
│   ├── devices_list
│   │   ├── biz
│   │   ├── data
│   │   └── service
│   ├── devies_event
│   │   ├── biz
│   │   ├── data
│   │   └── service
│   └── pkg //internal包内各服务共用组件
├── LICENSE
└── README.md
```
# 目录构建命令
```bash
mkdir -pv {configs,apis,cmd/DTU,internal/{pkg,devices_list/{biz,data,service},devices_health/{biz,data,service},devies_event/{biz,data,service}}}
```
