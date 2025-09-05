# 声明 
源码只供学习使用，如用于商业活动与本人无关，请勿将系统用于非法业务

## golangpay 聚合支付系统 四方支付系统 四方源码 golang源码 第四方支付平台 三方平台，uid平台，页游平台

go语言支付系统在线体验 演示地址：

运营端：https://homea.golangpay.com 

账号:xitong
登录密码：123456789

商户端：https://sha.golangpay.com 

账号：test 

登录密码：123456789

测试账号密码会定期更改，如不对请联系客服发新的账号密码

开发文档 https://www.golangpay.com

go语言支付系统演示地址：

运营端：http://home.golangpay.com 

帐号：xitong

登录密码：123456789

商户端：http://sh.golangpay.com 

帐号：mchtest 登录密码：mch123123

系统均采用golangpay语言开发，会golangpay的技术人员可以自行二次开发

golangpay是一套开箱即用、适合拿来直接运营的聚合支付系统。系统适合有技术团队的企业购买，我司可提供程序源码、技术文档和售后技术支持服务。

程序源码和文档包括哪些？ 源码包括：所有golang服务端源码和Layui前端源码，可二次开发，想怎么改就怎么改，So Easy !

文档包括：开发说明、系统部署、通道对接、API接口、线上运维、系统业务等。

技术支持有哪些服务？ 针对每个购买的客户，我司会单独创建群，至少指定一名技术支持人员单独提供售后技术支持。

技术支持内容包括：系统部署指导、二次开发指导、反馈Bug的修复、需求的收集等。

注：不提供软件开发环境搭建、不提供golang基础辅导、仅限该系统业务技术交流。

如需要最新完整商业版本请联系 飞机(Telegram)：@tianxiex

系统描述 golangpay为golang开发版，使用golang + vue + beego架构开发。包括运营平台、代理商系统、商户系统、支付系统，结算系统、对账系统等。

模块说明 golang-service 所有核心业务方法封装，供其他模块引用后调用

golang-core 核心包，包括golang服务接口以及实体beego,以及公用引用及常用工具类等

golang-manage 运营平台（接口和管理界面，前后端分离）

golang-merchant 商户系统（接口和管理界面，前后端分离）

golang-agent 代理商系统（接口和管理界面，前后端分离）

golang-pay 支付网关，提供商户访问的支付接口及对接所有支付通道实现

golang-task定时任务，包括对账服务、结算服务，部署时需单节点部署

golang-writeoff，核销端、提供核销API，和核销商提交话费，电费，油卡等核销户号，部署时需单节点部署

golang-z-api-base 支付接口的基础包
# 四方支付系统介绍
巨海四方支付系统是go语言开发支持多商户、多通道、能够自由进行对接配置、集成telegram机器人服务的聚合支付系统

## 商务合作请联系飞机：[tianxiex](https://t.me/tianxiex)

## 系统简介
## 运营端：四方运营人员、系统购买方
1.开户（给商户开通账号）

2.对接通道（通道供应商对接）

3.测试下单、补单

4.配置TG机器人、群发TG消息

5.其他功能

## 商户端：四方运营人员的客户、系统购买方的客户
1.查看开户信息（商户ID,密钥等）

2.查看对接文档

3.测试下单

4.其他功能

## 对接：四方运营人员、四方技术团队、系统购买方
1.提供四方对接文档给商户。（由四方技术团队协助、系统购买方）

2.对接通道。（由四方技术团队对接、系统购买方）

3.测试对接情况。（由四方技术团队、四方运营人员协助、系统购买方）

##TG机器人：运营方
1.四方运营人员（系统购买方）在系统配置运营人员

2.四方运营人员（系统购买方）在系统配置机器人相关信息

3.四方运营人员（系统购买方）在TG中创建针对商户的业务群，将已经配置好的机器人拉进群，并设置群管理角色

## 运营端部分演示
1.商户管理
​ 商户管理主要用来，创建商户，也就是开户，配置产品费率，以及配置产品对应跑的通道。

1.1 商户列表
![商户列表](https://raw.githubusercontent.com/golangpay/1/refs/heads/img/pay/3.png)



1.2 应用列表
![应用列表](https://raw.githubusercontent.com/golangpay/1/refs/heads/img/pay/4.png)

1.2.3 产品列表
image

产品配置通道（一个产品可以配置多个通道，这里就是运营人员需要重点关注的点）

![产品列表](https://raw.githubusercontent.com/golangpay/1/refs/heads/img/pay/5.png)

1.2.4 商户下单测试界面
商户下单测试可以快速通过不同商户，选择不同的产品，以及不同的通道，来进行测试或者正式下单，可以帮助大家快速获取下单的结果。

这里既可以针对商户对接，测试；也可以针对通道对接，测试。

![6](https://raw.githubusercontent.com/golangpay/1/refs/heads/img/pay/6.png)

## 2.订单管理
订单列表，可以实时查看当前系统订单情况。

1.可以查看每个商户，每个产品，每个通道，商户费率，通道费率，平台成本，平台利润，支付状态，回调状态，补单状态等。

![7](https://raw.githubusercontent.com/golangpay/1/refs/heads/img/pay/7.png)

2.打开实时统计，可以查看当前商户或产品或通道的统计情况

![8](https://raw.githubusercontent.com/golangpay/1/refs/heads/img/pay/8.png)

3.当供应商已经明确支付成功（有时金额不对的情况下），但没有回调成功时，可以通过人工补单操作，此时可以在备注信息中说明情况。补单密码是跟系统分开的。补单并不会修改订单金额，只是会将系统订单状态从支付中，改为已支付。补单有些情况需要跟商户沟通，是否需要进行补单。补单的目的，除了改变系统状态之外，也会向商户发起回调通知。

![9](https://raw.githubusercontent.com/golangpay/1/refs/heads/img/pay/9.png)

![10](https://raw.githubusercontent.com/golangpay/1/refs/heads/img/pay/10.png)

## 3.支付配置
### 3.1 供应商管理
​ 这里主要是用来对接通道的配置。开发这块的目的是为了：存在有优势的通道，需要我们去对接通道，但是，往往技术人员响应不及时，导致我们运营团队错失了很多赚钱的机会。那我们想：如果对接通道，不需要技术参与，我们运营人员能够自己对接，那该多好呀。在此应用场景下，供应商管理对接，就由此而生。 

![11](https://raw.githubusercontent.com/golangpay/1/refs/heads/img/pay/11.png)
参数映射配置：

第一步，供应商提供信息。包括（供应商开我们，开的户，密钥，下单地址，查单地址，等信息），具体还得看供应商的接口文档。 image 第二步：映射配置。包括下单、查单、回调，这三部分。下单包含，请求参数，返回参数配置；查单也一样，回调只有接收参数配置，而且每个配置都有说明。前期需要我们技术支持一下，带着一起配置几个通道方，以后运营人员就可以自己配置了。 

![12](https://raw.githubusercontent.com/golangpay/1/refs/heads/img/pay/12.png)
![13](https://raw.githubusercontent.com/golangpay/1/refs/heads/img/pay/13.png)
![14](https://raw.githubusercontent.com/golangpay/1/refs/heads/img/pay/14.png)
![15](https://raw.githubusercontent.com/golangpay/1/refs/heads/img/pay/15.png)
### 3.2 通道管理
​ 上述通过供应商对接后，我们需要将供应商提供的通道信息添加进来即可。选择指定的供应商，供应商提供的通道编码、通道名称、通道费率、收款类型（分为固定额度只能是指定的额度，比如：30|50|100|200和浮动额度，区间范围比如：1-9999）、是否启用。这里加了一个回收站：主要是方便管理，只展示有用的通道。暂时不需要的可以进行回收操作，到回收站中的通道管理进行查看。 


![16](https://raw.githubusercontent.com/golangpay/1/refs/heads/img/pay/16.png)
![17](https://raw.githubusercontent.com/golangpay/1/refs/heads/img/pay/17.png)
### 3.3 产品管理
​ 产品，在四方系统可以认为是一个大通道，我们给到商户的产品编码（也可以叫通道编码）和产品名称（也可以叫通道名称）。也就是对外推广和营销的大通道。这里由我们运营人员自定义一套产品（大通道）规则。
​ 可以提前自定义好，产品编码（大通道编码）和产品名称（大通道名称）。 

![18](https://raw.githubusercontent.com/golangpay/1/refs/heads/img/pay/18.png)

## 4.机器人管理
### 4.1 配置机器人
如果第一次接触TG机器人，需要到TG上注册一个机器人，获取token，进行设置，同时需要注册webhook方式，将机器人绑定到系统中去。请注意：如果是新部署的服务器域名必须要安装证书，也就是需要支持https

![19](https://raw.githubusercontent.com/golangpay/1/refs/heads/img/pay/19.png)
![20](https://raw.githubusercontent.com/golangpay/1/refs/heads/img/pay/20.png)
![21](https://raw.githubusercontent.com/golangpay/1/refs/heads/img/pay/21.png)

### 4.2 群组管理
这里的群组由运营人员，建好群后，并将机器人拉入群中，作为管理员权限，就会自动将群信息推送到后台系统中。如果运营人员通过命令绑定商户，就会自动关联群组。如果没有绑定，则显示暂未关联商户。同时，也可以自己建好群（频道）。建好群后，也可以添加群成员，设定权限。

![22](https://raw.githubusercontent.com/golangpay/1/refs/heads/img/pay/22.png)
![23](https://raw.githubusercontent.com/golangpay/1/refs/heads/img/pay/23.png)

### 4.3 消息管理
这里可以新增消息，选择对应的群，进行转发

![24](https://raw.githubusercontent.com/golangpay/1/refs/heads/img/pay/24.png)
## 商户系统部分演示
### 1.商户首页
包括商户的基本信息，对接信息等

![25](https://raw.githubusercontent.com/golangpay/1/refs/heads/img/pay/25.png)

### 2.商户收银台
这里可以测试下单

![26](https://raw.githubusercontent.com/golangpay/1/refs/heads/img/pay/26.png)
对接
对接在运营端中支付配置下的供应商管理界面。

TG机器人
TG机器人在运营端中机器人管理界面

日常机器人常见运营截图展示：

### 1.绑定商户

![26](https://raw.githubusercontent.com/golangpay/1/refs/heads/img/pay/27.png)

### 2.新增预付，及预付历史详情
![28](https://raw.githubusercontent.com/golangpay/1/refs/heads/img/pay/28.png)
![29](https://raw.githubusercontent.com/golangpay/1/refs/heads/img/pay/29.png)



## 合作模式
### 方式1：出租
服 务： 我们代买 服务器 + 域名 + 域名免费升级https + 搭建部署免费 + 功能更新不收费

通道对接：24+7技术在线服务 免费对接

机 器 人：带查单/对账机器人（免费）



### 方式2：出售
只出售源码，不代买(服务器 + 域名 + 域名免费升级https)，不帮忙搭建部署免费 + 不帮忙功能更新

### 商务合作

请联系飞机：[tianxiex](https://t.me/tianxiex)



