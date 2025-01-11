digital-mall
==

基于gin+gorm+mysql 的电子商城
-
本项目实现了一些基本的用户注册，登录，信息修改，绑定邮箱等一些基本操作，还有商品系列的创建，搜索，获取，分类，以及地址，购物车，还有订单的一些功能，对于支付功能尚未实现，仍在改进中，收藏夹功能也有点小问题，其他后端接口都已实现

如何测试并导入接口
=
！[导入接口，将 doc 文件夹下的 .json 文件导入 postman 中 import 即可](doc/导入.png)
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
# 项目结构
test_mysql/ ├── dao/ │ ├── order_dao.go │ ├── user_dao.go │ └── product_dao.go ├── model/ │ ├── order.go │ ├── user.go │ └── product.go ├── pkg/ │ ├── e/ │ │ └── error.go │ └── util/ │ ├── encrypt.go │ └── logger.go ├── serializer/ │ └── response.go ├── service/ │ └── order_pay.go └── main.go






简要说明 
=
1. mysql 是存储主要数据
2. redis 使用来存储商品浏览次数
3. 用户创建默认金额为 1w 

