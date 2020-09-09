# 报到系统流程梳理

## 会议信息管理

管理员

* 用户角色
* 密码
* 

大会信息

* 编号
* 会议名称
* 开始时间
* 结束时间
* 管理员组编号
* 备注

会议信息

* 编号

* 大会编号
* 子会议信息
  * 序号
  * 子会议名称
  * 是否判断迟到
  * 迟到时间
* 报到规则
  * 是否设置会议人数线
  * 规则人员类型
    * 委员
    * 所有类型
  * 规则（>或>=）
  * 规则按人数设置或比例设置
  * 会议人数线
  * 按比例判断
    * 按照委员或所有类型的
    * 比例
* 达标文字

分组信息

* 编号

* 大会编号
* 子会议编号
* 分组详细信息

人员信息

* 人员基本信息
  * 编号
  * 大会编号
  * 子会议编号
  * 姓名
  * 界别
  * 类型
    * 工作人员
    * 委员
    * 陪同人员...
  * 级别
  * 性别
  * 民族
  * 坐席
  * 单位
  * 联系方式
  * 照片名称
  * 主席团成员
    * 是\否
  * 职务
  * 卡片编号
  * 备注
  * 关注
  * 分组信息
    * 分组信息下标
    * 分组表编号

写卡软件（卡片信息）

* 编号
* 人员编号
* 卡片信息
  * 主卡号
  * 备卡号
  * 测试卡号

会议人员信息表

* 编号
* 子会议编号
* 人员编号
* 座位编号
  * 行
  * 列

## 会议进行中

会议报到信息表

* 编号
* 子会议编号
* 人员编号
* 报到时间
* 报到状态



## 安装nodejs

官网下载安装包下载

配置环境变量

```
# 配置的是npm安装的全局模块所在的路径，以及缓存cache的路径，之所以要配置，是因为以后在执行类似：npm install express [-g] （后面的可选参数-g，g代表global全局安装的意思）的安装语句时，会将安装的模块安装到【C:\Users\用户名\AppData\Roaming\npm】路径中，占C盘空间。
npm config set cache "D:\softwares\nodejs\node_cache"
npm config set prefix "D:\softwares\nodejs\node_global"
# 修改默认的用户变量C:\Users\lenovo\AppData\Roaming\npm为D:\softwares\nodejs\node_global
# 增加默认的环境变量D:\softwares\nodejs\node_global
```

### 安装vue

https://www.jianshu.com/p/e215a33aed40

### vue目录

* build 项目构建
* config 配置目录
  * 端口号
* node_modules 加载的项目依赖模块
* src 开发目录
  * assets 放置一些图片
    * logo等
  * components 组件文件，可以不用
  * App.vue 项目入口文件，可以把组件直接写在这里，而不使用components目录
  * main.js项目核心文件
* static 静态资源目录：图片、字体
* text 初始化测试目录，可删除
* .xxxx 配置文件
* index.html 首页入口文件，添加meta信息或统计代码
* package.json 项目配置文件
* readme 项目说明文件