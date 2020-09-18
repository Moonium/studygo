# studygo
根据老男孩的Go语言培训课程分日编写的代码

大项目为实现一个Log Agent日志文件收集系统：

将分散的机器上的日志实时收集，统一储存到中心系统，再对这些日志建立索引，最终通过友好的web页面对其进行检索

信息流：

Logs -> Log Agent (with etcd monitoring) -> kafka -> Log Transfer (with etcd monitoring) -> Elastic Search -> Kibana
