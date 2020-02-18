# gonsq
nsq消息队列测试

# 部署步骤和命令
PS：后台启动使用nohup即可，下面只是为了说明启动方式和命令参数

# 第一步需要启动nsqlookupd
./nsqlookupd
默认占用4161和4160两个端口
使用-http-address和-tcp-address可以修改

# 第二步启动两个nsqd
./nsqd -lookupd-tcp-address=192.168.1.102:4160 -broadcast-address=192.168.1.103 -data-path="/temp/nsq"
其中
-lookupd-tcp-address为上面nsqlookupd的IP和tcp的端口4160
-broadcast-address我填写的是自己的IP，这个IP官网上写的是会注册到nsqlookupd
-data-path为消息持久化的位置

# 第三步启动nsqadmin
./nsqadmin -lookupd-http-address=192.168.4.102:4161
同样需要指定-lookupd-http-address但是这次是http的端口也就是4161因为admin通过http请求来查询相关信息



# 组件
NSQ主要包含3个组件：

nsqd：在服务端运行的守护进程，负责接收，排队，投递消息给客户端。能够独立运行，不过通常是由 nsqlookupd 实例所在集群配置的
nsqlookup：为守护进程，负责管理拓扑信息并提供发现服务。客户端通过查询 nsqlookupd 来发现指定话题（topic）的生产者，并且 nsqd 节点广播话题（topic）和通道（channel）信息
nsqadmin：一套WEB UI，用来汇集集群的实时统计，并执行不同的管理任务


# 特性
消息默认不可持久化，虽然系统支持消息持久化存储在磁盘中（通过设置 –mem-queue-size 为零），不过默认情况下消息都在内存中
消息最少会被投递一次，假设成立于 nsqd 节点没有错误
消息无序，是由重新队列(requeues)，内存和磁盘存储的混合导致的，实际上，节点间不会共享任何信息。它是相对的简单完成疏松队列
支持无 SPOF 的分布式拓扑，nsqd 和 nsqadmin 有一个节点故障不会影响到整个系统的正常运行
支持requeue，延迟消费机制
消息push给消费者


# 流程
单个nsqd可以有多个Topic，每个Topic又可以有多个Channel。Channel能够接收Topic所有消息的副本，从而实现了消息多播分发；而Channel上的每个消息被分发给它的订阅者，从而实现负载均衡，所有这些就组成了一个可以表示各种简单和复杂拓扑结构的强大框架。