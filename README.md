digital-mall
==

基于gin+gorm+mysql 的电子商城
-
本项目实现了一些基本的用户注册，登录，信息修改，绑定邮箱等一些基本操作，还有商品系列的创建，搜索，获取，分类，以及地址，购物车，还有订单的一些功能，对于支付功能尚未实现，仍在改进中，收藏夹功能也有点小问题，其他后端接口都已实现

如何测试并导入接口
=
![导入接口，将 doc 文件夹下的 .json 文件导入 postman 中 import 即可](https://github.com/yehan5555/digital-mall/blob/update/doc/%E5%AF%BC%E5%85%A5.png)
需要注意要改变一下环境变量
url =http://127.0.0.1:3000/api/v1/
对于token的话， 注册，登入即可在postman 中生成，复制到 环境变量中即可

主要功能
=
用户注册登录,信息修改（jwt）  
邮箱的绑定与解绑，修改密码   
商品的发布，浏览   
购物车的加入，删除，浏览   
地址的创建，删除，删除   
订单的创建，删除  
对支付密码的对称加密   
支持事务，可以进行回滚  

主要依赖
=
|名称|版本|
|---|----|
|golang|1.23|
|gorm|1.25.12|
|gin|1.10.0|
|mysql|1.5.7|
|dbresolver|1.5.3|
|jwt-go|3.2.0|
|logrus|1.9.0|
|redis|6.15.9|

项目结构
=
gin-mall  
api         #用于定义接口函数，也就是controller的作用  
cache       #redis 缓存  
cmd         #程序入口   
conf        #配置文件  
dao         # 对数据库进行操作  
doc         #文档  
logs        #日志文件  
middleware  #中间件  
model       #数据库模型  
pkg         # e 存放错误码，util 存放工具函数  
routes      #路由配置  
serializer  #将数据库序列化  
service     #接口函数的实现  
static      #存放静态文件  
 


简要说明 
=
1. mysql 是存储主要数据
2. redis 使用来存储商品浏览次数
3. 用户创建默认金额为 1w 

postman 的一些接口演示
=

用户登录
-
![用户登录](https://github.com/yehan5555/digital-mall/blob/update/doc/%E7%94%A8%E6%88%B7%E7%99%BB%E5%BD%95.png)

 创建商品
 -
![创建商品](https://github.com/yehan5555/digital-mall/blob/update/doc/%E5%88%9B%E5%BB%BA%E5%95%86%E5%93%81.png)

创建地址
-
 ![创建地址](https://github.com/yehan5555/digital-mall/blob/update/doc/%E5%88%9B%E5%BB%BA%E5%9C%B0%E5%9D%80.png)

创建订单
-
![创建订单](https://github.com/yehan5555/digital-mall/blob/update/doc/%E5%88%9B%E5%BB%BA%E8%AE%A2%E5%8D%95.png)

创建购物车
-
![创建购物车](https://github.com/yehan5555/digital-mall/blob/update/doc/%E5%88%9B%E5%BB%BA%E8%B4%AD%E7%89%A9%E8%BD%A6.png)

搜索商品
-
![搜索商品](https://github.com/yehan5555/digital-mall/blob/update/doc/%E6%90%9C%E7%B4%A2%E5%95%86%E5%93%81.png)

验证邮箱
-
![验证邮箱](https://github.com/yehan5555/digital-mall/blob/update/doc/%E9%AA%8C%E8%AF%81%E9%82%AE%E7%AE%B1.png)





