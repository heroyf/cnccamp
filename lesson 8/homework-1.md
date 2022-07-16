# 优雅启动

通过liveprobe 和 readness probe来完成检测

![img.png](screenshots/img.png)

# 优雅终止

这里使用tini来做为容器的1号进程，所以需要对原先的Dockerfile加以修改

![img_1.png](screenshots/img_1.png)

需要注意的是alpine缺少tini依赖的静态库，所以需要下载tini-static，验证docker 查看command

![img_2.png](screenshots/img_2.png)

可以看到是由tini拉起

# 资源需求和QOS

QOS类型为：Burstable

![img_3.png](screenshots/img_3.png)

# 代码和配置分离

采用了一个环境变量来控制logLevel 

![img_4.png](screenshots/img_4.png)

K8s 中注入env

![img_6.png](screenshots/img_6.png)

验证loglevel级别，k8s 管理的pod中已经是1，而Dockerfile中指定的是2

![img_5.png](screenshots/img_5.png)

完整yaml参考lesson 8下的其他yaml文件