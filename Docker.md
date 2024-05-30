**场景：使用Docker创建两个Mysql容器**

问题：创建第一个容器，默认端口都是3306，但是第二个容器创建默认都是3307的时候，连不上**（经过证实可以连上）**

原因：第一个Mysql内部端的端口默认是3306，容器端口和宿主机端口也是3306，所以没问题，但是第二个由于没有指定Mysql内部端口，即使容器端口和宿主机端口都是3307，但是第二个Mysql配置文件里面的默认端口还是3306，所以有两种方法修改Mysql内部端口方法。

方法一：创建时指定Mysql内部端口（推荐）

```
//推荐使用这种  创建时指定Mysql内部端口：-e  MYSQL_TCP_PORT=3309
//这样好处就是不用登录进去里面该配置文件端口
docker run  --name mysql-4 -d -p 3309:3309 -e MYSQL_ROOT_PASSWORD=root -e  MYSQL_TCP_PORT=3309 mysql:8.0
```

如果不指定端口的话

```
docker run  --name mysql-3 -d -p 3308:3308 -e MYSQL_ROOT_PASSWORD=root mysql:8.0
不指定Mysql内部端口，默认是3306，但是做映射容器端口3308，所以连不上
```

**验证容器端口都是3306情况创建两个mysql容器，映射宿主机的端口是3307和3308**

```
docker run --name db-mysql -e MYSQL_ROOT_PASSWORD=root -d -p 3307:3306 mysql:8.0
docker run --name db-mysql2 -e MYSQL_ROOT_PASSWORD=root -d -p 3308:3306 mysql:8.0
```

参数解析：

**3307**：容器映射宿主机的端口，

**3308**：容器映射宿主机的端口，

**3306**：docker中mysql的运行端口（每个相互独立）

上面使用 `MYSQL_TCP_PORT=3309`只是修改当前容器启动mysq默认的

方法二：（麻烦不推荐）

进入容器里面，然后修改Mysql配置文件

#### 一、Docker基础

![image-20240529095550902](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240529095550902.png)

##### 1、registry镜像仓库

registry可以理解为镜像仓库，用于保存docker image。

Docker Hub是docker官方的镜像仓库，docker命令默认从docker hub中拉取镜像。我们也可以搭建自己的镜像仓库。

##### 2、image镜像

image可以理解为一个只读的应用模版。image包含了应用程序及其所需的依赖环境，例如可执行文件、环境变量、初始化脚本、启动命令等。

###### docker pull拉取镜像

从镜像仓库拉取镜像到本地

`docker pull nginx`不写默认是latest

`docker pull nginx:latest`

`docker pull nginx:1.22`

`docker pull nginx:1.22.0-aplpine`

**一般不建议使用latest，因为最新的镜像是滚动更新的，过一段时间，可能跟你本地的不是同一个。**

使用`docker images`查看本地镜像

****

##### 3、container容器

容器时image的一个运行实例。当我们运行一个image，就创建了一个容器。

可以使用以下命令操作容器：

**`docker start db-mysql`** 启动容器

**`docker stop`**   关闭容器

**`docker restart`** 重启容器

**`docker rm`** 删除容器

****

###### docker run启动容器

**`docker run [可选参数] 镜像名:版本 []`**

**启动一个容器示例：**

**`docker run --name db-mysql -e MYSQL_ROOT_PASSWORD=root -d -p 3307:3306 mysql:8.0`**

![image-20240529213903100](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240529213903100.png)

创建成功会返回一个容器ID，是有64个十六进制字符组成的字符串。

当使用**`docker ps`**查看正在运行的容器时

![image-20240529214627284](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240529214627284.png)

容器ID可能会缩短为12字符，以提高可读性，这个短ID是容器ID的前12个字符。

上面头列说明：

- **CONTAINER ID**:容器的唯一表示ID。
- **IMAGE**:创建容器时使用的镜像。
- **COMMAND**:容器最后运行的命令。
- **CREATED**:创建容器的时间。
- **STATUS**:容器状态。
- **PORTS**:对外开放的端口。
- **NAMES**:容器名。可以和容器ID一样唯一标识容器，同一台宿主机上不允许有同名容器存在，否则会冲突。

****

###### 停止/关闭容器

- **docker stop [NAME]/[CONTAINER ID]**:将容器退出。
- **docker kill [NAME]/[CONTAINER ID]**:强制停止一个容器。

![image-20240529221608607](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240529221608607.png)

成功关闭一个容器会返回对应容器的容器名。

****

###### 删除容器

**命令：`docker rm [NAME]/[CONTAINER ID]`**

**`docker rm [NAME]/[CONTAINER ID]`:**不能够删除一个正在运行的容器，会报错。需要先停止容器。

- **删除正在运行的容器，实例如图所示：**

![image-20240529222054450](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240529222054450.png)

删除正在运行的容器会报错误。

- **删除已经停止运行的容器：**

![image-20240529222330876](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240529222330876.png)

删除成功会返回删除的容器名。

****

###### 公开端口(-p)

**`docker run --name some-nginx -d -p 8080:80 ningx:1.22`**

默认情况下，容器无法通过外部网络访问。

需要使用**`-p`**参数将容器的端口映射到宿主机端口，才可以通过宿主机IP进行访问。

**参数说明：**

- **`--name`** ：容器的名称
- **`-d`** ：**Detached 模式**，容器启动后将在后台运行
- **`-p`**： 将容器端口映射到宿主机端口，`-p 8080:80`，左边8080代表的是宿主机端口，右边的是容器端口。

浏览器打开http://217.0.0.1:8080或者自己对应的IP

![image-20240529201454612](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240529201454612.png)

**-p 8080-8090:8080-8090**公开端口范围，前后必须对应

**-p 192.168.56.106:8080:80**如果宿主机有多个ip，可以指定绑定到哪个ip

****

###### 前台交互运行

创建一个新的容器，使用mysql客户端

```
docker run -it --rm mysql:5.7 mysql -h172.17.0.2 -uroot -p
```

**`-it`** 使用交互模式，可以在控制台里输入、输出

**`--rm`**在容器**退出时自动删除容器**。一般在使用客户端程序时使用此参数。

如果每次使用客户端都创建一个新的容器，这样将占用大量的系统空间。

`mysql -h172.17.0.2 -uroot -p`表示启动容器时执行的命令。

**`docker exec`**在运行的容器中执行命令，一般配合`-it`参数使用交互模式

下面实例图的例子对应上面创建的mysql容器：

![image-20240529224745093](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240529224745093.png)

****

###### 常用命令

- `docker ps` 查看正在运行的容器
- `docker ps -a` 查看所有容器，包括正在运行和停止的
- `docker inspect` 查看容器的信息
- `docker logs`查看日志
- `docker cp` 在容器和宿主机间复制文件

#### 二、docker网络

默认网络

docker会自动创建三个网络，**`bridge`**、**`host`**、**`none`**

![image-20240530101500011](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240530101500011.png)

- **bridge(桥接网络)：**如果不指定容器连接的网络，新创建的容器默认将连接到bridge网络。
  默认情况下，使用bridge网络，宿主机可以ping通容器ip，容器中也能ping通宿主机。
  容器之间只能通过IP地址相互访问，由于容器的IP会随着启动顺序发生变化，因此不推荐使用IP访问。

- **host：**容器与宿主机共享网络，不需要映射端口即可通过主机ip访问。（-p选项会被忽略）
  主机模式网络可用于优化性能，在容器需要处理大量端口的情况下，它不需要网络地址切换（NAT），并且不会为每个端口创建“用户空间代理”。

  > 慎用，可能会有安全问题。

- **none：**禁用容器中所用网络，在启动容器时使用，提供了一种空白的网络配置，方便用户排除其他干扰，用于自定义网络。