### 简介
  一个开源的bug管理系统，IT人员开发全过程，文件存储，接口文档, 单接口测试功能  
 

文档地址： http://itflow.doc.hyahm.com  
路由接口文档地址： http://hyahm.com:10001/docs （配合新前端不断更新中）

更改项目路由为：  https://github.com/hyahm/xmux   
新前端地址：  https://github.com/jiandanzhiyun/vue-elementAdmin （开发完成后会替换掉老的前端）    

### 功能
- [x] 增加bug，改变bug状态，转交bug 
- [x] 部门管理
- [x] 显示bug列表,搜索、分页
- [x] 用户创建及其操作  
- [x] 上传个人头像  
- [x] 增加缓存机制 
- [x] 增加邮件通知功能  
- [x] 增加admin用户的信息重置接口  
   admin用户有且只有一个，注册admin账户建议直接操作数据库，然后修改密码即可
   如果忘记admin的密码，可以执行下面命令重置密码，如下所示，只能在go服务器那台机器上执行
```
   curl http://127.0.0.1:10001/admin/reset?password=123
```
- [x] 增加修改邮箱，昵称，姓名页面  
- [x] 只允许修改自己部门的账号权限   增加用户修改权限功能  
- [x] bug可以指定多人，自己的bug才可以转交，删除bug内部转交功能，增加缓存,增加查看所有bug的权限  
- [x] 增加用户禁用功能，当此用户存在bug时，无法被删除  
- [x] 禁用用户，此用户的所有发布的bug都将移动至垃圾箱，垃圾箱里面的bug只有管理员才能查看，启用用户会将此用户的bug改为非垃圾箱  
- [x] 增加操作日志，只有管理员才能查看   
- [x] 状态实时保存，增加缓存  
- [x] 数据库表自动更新
- [x] 接口文档自定义数据类型， 无法定义基础类型  
- [x] 增加接口测试  
- [x] go代码优化
- [x] vue代码优化
- [x] 为了更好的使用，增加消息提示
- [ ] 增加docker-compose一条命令启动服务

### 展示页面： 
   展示页面会更新为最新可使用的代码  
   [ITflow](http://bug.hyahm.com "ITflow")  
   
 

### 项目优势   
1. 部署简单,使用简单    
2. 因为后端是go语言写的，跨平台，速度快（虽然这个没什么卵用的样子）  
3. vue平台用的vue-element-admin框架上写的  
4. 永久开源  可以自己二次开发 
  
### QQ群  
    928790087@qq.com  
