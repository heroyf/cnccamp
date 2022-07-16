# 优雅启动

通过liveprobe 和 readness probe来完成检测

![image-20220717001046069](/Users/heroyf/Library/Application Support/typora-user-images/image-20220717001046069.png)

# 优雅终止

这里使用tini来做为容器的1号进程，所以需要对原先的Dockerfile加以修改

![image-20220717001233868](/Users/heroyf/Library/Application Support/typora-user-images/image-20220717001233868.png)

需要注意的是alpine缺少tini依赖的静态库，所以需要下载tini-static，验证docker 查看command

![image-20220717001400910](/Users/heroyf/Library/Application Support/typora-user-images/image-20220717001400910.png)

可以看到是由tini拉起

# 资源需求和QOS

QOS类型为：Burstable

![image-20220717001512255](/Users/heroyf/Library/Application Support/typora-user-images/image-20220717001512255.png)

# 代码和配置分离

采用了一个环境变量来控制logLevel 

![image-20220717001656055](/Users/heroyf/Library/Application Support/typora-user-images/image-20220717001656055.png)

K8s 中注入env

![image-20220717001716317](/Users/heroyf/Library/Application Support/typora-user-images/image-20220717001716317.png)

验证loglevel级别，k8s 管理的pod中已经是1，而Dockerfile中指定的是2

![image-20220717001758546](/Users/heroyf/Library/Application Support/typora-user-images/image-20220717001758546.png)

完整yaml参考lesson 8下的其他yaml文件