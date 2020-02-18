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