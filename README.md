# **dnsmasq-admin**
#### golang开发的dnsmasq开发的简洁web管理界面
#### 源码安装请下载golang1.11以上版本

### **获取**
```shell
git clone https://github.com/cnns-jd/dnsmasq-admin
```
### **编译服务端**
```shell
cd dnsmasq-admin/cmd/dnsmasq-admin
go build
```
### 某些地区无法使用golang包请设置goproxy后再编译
```shell
export goproxy=https://goproxy.cn
```
### **启动服务端**
```shell
# 服务端tcp绑定地址
export LISTEN_ADDRESS=0.0.0.0:9001
# 身份验证密码
export WORK_DIR=/etc/dnsmasq.host.d
./dnsmasq-admin
```
* [Issue]

[Issue]:     https://github.com/cnns-jd/dnsmasq-admin/issues?state=open

免责声明
------

请遵守当地国家法律法规

本代码仅供学习使用，使用本代码一切后果与作者无关
