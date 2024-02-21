# KubeZ - 多集群k8s管理平台
KubeZ是一个多集群k8s管理平台,采用前后端分离架构,使用vue和go编写,旨在简化k8s操作难度
演示地址：（暂无）
部署文档：[deploy.md](doc/deploy.md)

前端栈：vue3，element-plus，axios，pinia，vue-route
后端栈：go，gins，go-client，logrous

目前实现的功能如下:

* 添加、管理多个集群
* 查看、创建、删除常用资源：如Pod，Deployment，Svc，ConfigMap等
* 交互式yaml生成器：通过表单生成完整资源yaml
* 部分资源的编辑

未来可能会更新的功能有：

* 管理员登录
* 用户组鉴权
* 跨集群资源复制
* 自动调度
* ……

### 预览图
![alt text](doc/image.png)

![alt text](doc/image2.png)

![alt text](doc/image3.png)
