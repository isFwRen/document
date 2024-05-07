#### Docker

场景：使用Docker创建两个Mysql容器

问题：创建第一个容器，默认端口都是3306，但是第二个容器创建默认都是3307的时候，连不上

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

方法二：（麻烦不推荐）

进入容器里面，然后修改Mysql配置文件