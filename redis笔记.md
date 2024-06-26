### 一、Redis配置

Redis 的配置文件位于 Redis 安装目录下，文件名为 **redis.conf**(Windows 名为 redis.windows.conf)。

你可以通过 **CONFIG** 命令查看或设置配置项。

------

**语法Redis CONFIG 命令格式如下：`redis 127.0.0.1:6379> CONFIG GET CONFIG_SETTING_NAME`**

**实例**

```
redis 127.0.0.1:6379> CONFIG GET loglevel

1) "loglevel"
2) "notice"
```

使用 ***** 号获取所有配置项：

**实例**

```
redis 127.0.0.1:6379> CONFIG GET *

  1) "dbfilename"
  2) "dump.rdb"
  3) "requirepass"
  4) ""
  5) "masterauth"
  6) ""
  7) "unixsocket"
  8) ""
  9) "logfile"
 10) ""
 11) "pidfile"
 12) "/var/run/redis.pid"
 13) "maxmemory"
 14) "0"
 15) "maxmemory-samples"
 16) "3"
 17) "timeout"
 18) "0"
 19) "tcp-keepalive"
 20) "0"
 21) "auto-aof-rewrite-percentage"
 22) "100"
 23) "auto-aof-rewrite-min-size"
 24) "67108864"
 25) "hash-max-ziplist-entries"
 26) "512"
 27) "hash-max-ziplist-value"
 28) "64"
 29) "list-max-ziplist-entries"
 30) "512"
 31) "list-max-ziplist-value"
 32) "64"
 33) "set-max-intset-entries"
 34) "512"
 35) "zset-max-ziplist-entries"
 36) "128"
 37) "zset-max-ziplist-value"
 38) "64"
 39) "hll-sparse-max-bytes"
 40) "3000"
 41) "lua-time-limit"
 42) "5000"
 43) "slowlog-log-slower-than"
 44) "10000"
 45) "latency-monitor-threshold"
 46) "0"
 47) "slowlog-max-len"
 48) "128"
 49) "port"
 50) "6379"
 51) "tcp-backlog"
 52) "511"
 53) "databases"
 54) "16"
 55) "repl-ping-slave-period"
 56) "10"
 57) "repl-timeout"
 58) "60"
 59) "repl-backlog-size"
 60) "1048576"
 61) "repl-backlog-ttl"
 62) "3600"
 63) "maxclients"
 64) "4064"
 65) "watchdog-period"
 66) "0"
 67) "slave-priority"
 68) "100"
 69) "min-slaves-to-write"
 70) "0"
 71) "min-slaves-max-lag"
 72) "10"
 73) "hz"
 74) "10"
 75) "no-appendfsync-on-rewrite"
 76) "no"
 77) "slave-serve-stale-data"
 78) "yes"
 79) "slave-read-only"
 80) "yes"
 81) "stop-writes-on-bgsave-error"
 82) "yes"
 83) "daemonize"
 84) "no"
 85) "rdbcompression"
 86) "yes"
 87) "rdbchecksum"
 88) "yes"
 89) "activerehashing"
 90) "yes"
 91) "repl-disable-tcp-nodelay"
 92) "no"
 93) "aof-rewrite-incremental-fsync"
 94) "yes"
 95) "appendonly"
 96) "no"
 97) "dir"
 98) "/home/deepak/Downloads/redis-2.8.13/src"
 99) "maxmemory-policy"
100) "volatile-lru"
101) "appendfsync"
102) "everysec"
103) "save"
104) "3600 1 300 100 60 10000"
105) "loglevel"
106) "notice"
107) "client-output-buffer-limit"
108) "normal 0 0 0 slave 268435456 67108864 60 pubsub 33554432 8388608 60"
109) "unixsocketperm"
110) "0"
111) "slaveof"
112) ""
113) "notify-keyspace-events"
114) ""
115) "bind"
116) ""
```

------

**编辑配置**

你可以通过修改 redis.conf 文件或使用 **CONFIG set** 命令来修改配置。

**语法**

**CONFIG SET** 命令基本语法：

```
redis 127.0.0.1:6379> CONFIG SET CONFIG_SETTING_NAME NEW_CONFIG_VALUE
```

**实例**

```
redis 127.0.0.1:6379> CONFIG SET loglevel "notice"
OK
redis 127.0.0.1:6379> CONFIG GET loglevel

1) "loglevel"
2) "notice"
```

------

**参数说明**

redis.conf 配置项说明如下：

| 序号 | 配置项                                                       | 说明                                                         |
| :--- | :----------------------------------------------------------- | :----------------------------------------------------------- |
| 1    | `daemonize no`                                               | Redis 默认不是以守护进程的方式运行，可以通过该配置项修改，使用 yes 启用守护进程（Windows 不支持守护线程的配置为 no ） |
| 2    | `pidfile /var/run/redis.pid`                                 | 当 Redis 以守护进程方式运行时，Redis 默认会把 pid 写入 /var/run/redis.pid 文件，可以通过 pidfile 指定 |
| 3    | `port 6379`                                                  | 指定 Redis 监听端口，默认端口为 6379，作者在自己的一篇博文中解释了为什么选用 6379 作为默认端口，因为 6379 在手机按键上 MERZ 对应的号码，而 MERZ 取自意大利歌女 Alessia Merz 的名字 |
| 4    | `bind 127.0.0.1`                                             | 绑定的主机地址                                               |
| 5    | `timeout 300`                                                | 当客户端闲置多长秒后关闭连接，如果指定为 0 ，表示关闭该功能  |
| 6    | `loglevel notice`                                            | 指定日志记录级别，Redis 总共支持四个级别：debug、verbose、notice、warning，默认为 notice |
| 7    | `logfile stdout`                                             | 日志记录方式，默认为标准输出，如果配置 Redis 为守护进程方式运行，而这里又配置为日志记录方式为标准输出，则日志将会发送给 /dev/null |
| 8    | `databases 16`                                               | 设置数据库的数量，默认数据库为0，可以使用SELECT 命令在连接上指定数据库id |
| 9    | `save <seconds> <changes>`Redis 默认配置文件中提供了三个条件：**save 900 1****save 300 10****save 60 10000**分别表示 900 秒（15 分钟）内有 1 个更改，300 秒（5 分钟）内有 10 个更改以及 60 秒内有 10000 个更改。 | 指定在多长时间内，有多少次更新操作，就将数据同步到数据文件，可以多个条件配合 |
| 10   | `rdbcompression yes`                                         | 指定存储至本地数据库时是否压缩数据，默认为 yes，Redis 采用 LZF 压缩，如果为了节省 CPU 时间，可以关闭该选项，但会导致数据库文件变的巨大 |
| 11   | `dbfilename dump.rdb`                                        | 指定本地数据库文件名，默认值为 dump.rdb                      |
| 12   | `dir ./`                                                     | 指定本地数据库存放目录                                       |
| 13   | `slaveof <masterip> <masterport>`                            | 设置当本机为 slave 服务时，设置 master 服务的 IP 地址及端口，在 Redis 启动时，它会自动从 master 进行数据同步 |
| 14   | `masterauth <master-password>`                               | 当 master 服务设置了密码保护时，slave 服务连接 master 的密码 |
| 15   | `requirepass foobared`                                       | 设置 Redis 连接密码，如果配置了连接密码，客户端在连接 Redis 时需要通过 AUTH <password> 命令提供密码，默认关闭 |
| 16   | ` maxclients 128`                                            | 设置同一时间最大客户端连接数，默认无限制，Redis 可以同时打开的客户端连接数为 Redis 进程可以打开的最大文件描述符数，如果设置 maxclients 0，表示不作限制。当客户端连接数到达限制时，Redis 会关闭新的连接并向客户端返回 max number of clients reached 错误信息 |
| 17   | `maxmemory <bytes>`                                          | 指定 Redis 最大内存限制，Redis 在启动时会把数据加载到内存中，达到最大内存后，Redis 会先尝试清除已到期或即将到期的 Key，当此方法处理 后，仍然到达最大内存设置，将无法再进行写入操作，但仍然可以进行读取操作。Redis 新的 vm 机制，会把 Key 存放内存，Value 会存放在 swap 区 |
| 18   | `appendonly no`                                              | 指定是否在每次更新操作后进行日志记录，Redis 在默认情况下是异步的把数据写入磁盘，如果不开启，可能会在断电时导致一段时间内的数据丢失。因为 redis 本身同步数据文件是按上面 save 条件来同步的，所以有的数据会在一段时间内只存在于内存中。默认为 no |
| 19   | `appendfilename appendonly.aof`                              | 指定更新日志文件名，默认为 appendonly.aof                    |
| 20   | `appendfsync everysec`                                       | 指定更新日志条件，共有 3 个可选值：**no**：表示等操作系统进行数据缓存同步到磁盘（快）**always**：表示每次更新操作后手动调用 fsync() 将数据写到磁盘（慢，安全）**everysec**：表示每秒同步一次（折中，默认值） |
| 21   | `vm-enabled no`                                              | 指定是否启用虚拟内存机制，默认值为 no，简单的介绍一下，VM 机制将数据分页存放，由 Redis 将访问量较少的页即冷数据 swap 到磁盘上，访问多的页面由磁盘自动换出到内存中（在后面的文章我会仔细分析 Redis 的 VM 机制） |
| 22   | `vm-swap-file /tmp/redis.swap`                               | 虚拟内存文件路径，默认值为 /tmp/redis.swap，不可多个 Redis 实例共享 |
| 23   | `vm-max-memory 0`                                            | 将所有大于 vm-max-memory 的数据存入虚拟内存，无论 vm-max-memory 设置多小，所有索引数据都是内存存储的(Redis 的索引数据 就是 keys)，也就是说，当 vm-max-memory 设置为 0 的时候，其实是所有 value 都存在于磁盘。默认值为 0 |
| 24   | `vm-page-size 32`                                            | Redis swap 文件分成了很多的 page，一个对象可以保存在多个 page 上面，但一个 page 上不能被多个对象共享，vm-page-size 是要根据存储的 数据大小来设定的，作者建议如果存储很多小对象，page 大小最好设置为 32 或者 64bytes；如果存储很大大对象，则可以使用更大的 page，如果不确定，就使用默认值 |
| 25   | `vm-pages 134217728`                                         | 设置 swap 文件中的 page 数量，由于页表（一种表示页面空闲或使用的 bitmap）是在放在内存中的，，在磁盘上每 8 个 pages 将消耗 1byte 的内存。 |
| 26   | `vm-max-threads 4`                                           | 设置访问swap文件的线程数,最好不要超过机器的核数,如果设置为0,那么所有对swap文件的操作都是串行的，可能会造成比较长时间的延迟。默认值为4 |
| 27   | `glueoutputbuf yes`                                          | 设置在向客户端应答时，是否把较小的包合并为一个包发送，默认为开启 |
| 28   | `hash-max-zipmap-entries 64 hash-max-zipmap-value 512`       | 指定在超过一定的数量或者最大的元素超过某一临界值时，采用一种特殊的哈希算法 |
| 29   | `activerehashing yes`                                        | 指定是否激活重置哈希，默认为开启（后面在介绍 Redis 的哈希算法时具体介绍） |
| 30   | `include /path/to/local.conf`                                | 指定包含其它的配置文件，可以在同一主机上多个Redis实例之间使用同一份配置文件，而同时各个实例又拥有自 |

**什么是守护进程？**

守护进程（Daemon Process），也就是通常说的 Daemon 进程（精灵进程），是 Linux 中的后台服务进程。它是一个生存期较长的进程，通常独立于控制终端并且周期性地执行某种任务或等待处理某些发生的事件。

守护进程是个特殊的孤儿进程，这种进程脱离终端，为什么要脱离终端呢？之所以脱离于终端是为了避免进程被任何终端所产生的信息所打断，其在执行过程中的信息也不在任何终端上显示。由于在 linux 中，每一个系统与用户进行交流的界面称为终端，每一个从此终端开始运行的进程都会依附于这个终端，这个终端就称为这些进程的控制终端，当控制终端被关闭时，相应的进程都会自动关闭。

**配置 redis 外网可访问**

由于 redis 采用的安全策略，默认会只准许本地访问。需要通过简单配置，完成允许外网访问。

修改 redis 的配置文件，将所有 bind 信息全部屏蔽。

```
# bind 192.168.1.100 10.0.0.1 
# bind 192.168.1.8 
# bind 127.0.0.1
```

修改完成后，需要重新启动 redis 服务。

修改 Linux 的防火墙(iptables)，开启你的 redis 服务端口，默认是 6379。

```
-A INPUT -m state –state NEW -m tcp -p tcp –dport 6379 -j ACCEPT 
…… 
-A INPUT -j REJECT –reject-with icmp-host-prohibited
```

请注意，一定要将 redis 的防火墙配置放在 REJECT 的前面。然后执行 **service iptables restart**。

至此，就能够链接到 redis 服务，并且能够正确显示了。

详情请见: [配置 redis 外网可访问](https://www.runoob.com/w3cnote/redis-external-network.html)

****

强烈建议不要在公网访问 redis,因为 redis 的处理速度非常快，所以如果你的密码比较简单，很容易就会通过暴力破解破解出密码：

```
# Warning: since Redis is pretty fast an outside user can try up to
# 150k passwords per second against a good box. This means that you should
# use a very strong password otherwise it will be very easy to break.
```

### 二、Redis数据类型

五种基本数据类型：字符串(String)、列表(List)、集合(Set)、有序集合(SortedSet)、哈希(Hash)

五种高级数据类型：消息队列(Stream)、地理空间(Geospatial)、HyperLogLog、位图(Bitmap)、位域(Bitfield)

**简介：**

- 性能极高
- 数据类型丰富，单键值对最大支持512M大小的数据
- 简单易用，支持多有主流编程语言
- 支持数据持久化、主从复制、哨兵模式等高可用特性

**各个数据类型应用场景：**

| 类型                 | 简介                                                   | 特性                                                         | 场景                                                         |
| :------------------- | :----------------------------------------------------- | :----------------------------------------------------------- | :----------------------------------------------------------- |
| String(字符串)       | 二进制安全                                             | 可以包含任何数据,比如jpg图片或者序列化的对象,一个键最大能存储512M | ---                                                          |
| Hash(字典)           | 键值对集合,即编程语言中的Map类型                       | 适合存储对象,并且可以像数据库中update一个属性一样只修改某一项属性值(Memcached中需要取出整个字符串反序列化成对象修改完再序列化存回去) | 存储、读取、修改用户属性                                     |
| List(列表)           | 链表(双向链表)                                         | 增删快,提供了操作某一段元素的API                             | 1,最新消息排行等功能(比如朋友圈的时间线) 2,消息队列          |
| Set(集合)            | 哈希表实现,元素不重复                                  | 1、添加、删除,查找的复杂度都是O(1) 2、为集合提供了求交集、并集、差集等操作 | 1、共同好友 2、利用唯一性,统计访问网站的所有独立ip 3、好友推荐时,根据tag求交集,大于某个阈值就可以推荐 |
| Sorted Set(有序集合) | 将Set中的元素增加一个权重参数score,元素按score有序排列 | 数据插入集合时,已经进行天然排序                              | 1、排行榜 2、带权重的消息队列                                |

### 三、redis为什么快

**基于内存**

内存的读写时比磁盘读写快很多的。redis是基于内存存储实现的数据库，相对于数据存在磁盘的数据库，就省去磁盘I/O的消耗。MySQL等磁盘数据库，需要建立索引来加快查询效率，而redis数据存放在内存，直接操作内存，所以很就很快。

![image-20240428135631269](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240428135631269.png)

**I/O多路复用**

什么是I/O多路复用？

- I/O：网络I/O
- 多路：多个网络连接
- 复用：复用同一个线程
- IO多路复用其实就是一种同步IO模型，它实现了一个线程可以监视多个文件句柄；一旦某个文件句柄就绪，就能够通知应用程序进行相应的读写操作；而没有文件句柄就绪时，就会阻塞应用程序，交出cpu

![image-20240428143816442](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240428143816442.png)

> 多路I/O复用技术可以让单个线程高效的处理多个连接请求，而redis使用epoll作为I/O多路复用技术的实现。并且redis自身的时间处理模型将epoll中的连接、读写、关闭都转换为事件，不在网络I/O上浪费过多的时间。

**单线程（合理的线程模型）**

单线程模型：避免了上下文切换

redis时单线程的，其实是**指redis的网络IO和键值对读写**时有 一个线程来完成的。单redis的其他功能，比如持久化、异步删除、集群数据同步等等，实际时由额外的线程执行的

### 四、线程io模型

redis是不是单线程？

6.0了以后发生了哪些变化？

针对问题1，一方面需要业务人员去规避，一方面Redis在4.0推出了**lazy-free机制，把bigkey释放内存的耗时操作放在了异步线程中执行，降低对主线程的影响**。

针对问题2，Redis在6.0推出了多线程，可以在高并发场景下利用**CPU多核多线程**读写客户端数据，进一步提升server性能，当然，**只是针对客户端的读写是并行的**，**每个命令的真正操作依旧是单线程的**。

### 五、持久化机制

- **为什么需要持久化？**

Redis是个基于内存的数据库。那服务器一旦宕机，内存中的数据将全部丢失。通常的解决方案是从后端数据库恢复这些数据，但后端数据库由性能瓶颈，如果是大数据量的恢复，1、会对数据库带来巨大的压力，2、数据库的性能不如Redis。导致程序响应慢。所以对Redis来说，实现数据的持久化，避免从后端数据库中恢复数据，是至关重要的。

- **Redis持久化有哪些方式呢？为什么我们需要重点学习RDB和AOF？**

从严格意义上说，Redis服务提供了四种持久化存储方案：`RDB`和`AOF`、`虚拟内存（VM）`和`DISKTORE`。虚拟内存（VM）方式，从Redis Version2.4开始就被官方明确表示不会再建议使用，Version 3.2版本中更找不到关于虚拟内存（VM）的任何配置范例，Redis的主要作者Salvatore Sanfilippo还专门写了一篇论文，来反思Redis对虚拟内存（VM）存储技术的支持问题。

至于DISKSTORE方式，是从Redis Version 2.8版本提出的一个存储设想，到目前为止Redis官方也没有任何stable版本中明确建议使用这方式，最关键的是目前官方文档上能够看到的Redis对持久化存储的支持明确就只有两种方案：**RDB和AOF**

#### AOF：增量备份

![image-20240501122117556](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240501122117556.png)

这种保存写操作命令到的持久化方式，就是redis里的AOF(Append Only File)持久化功能，注意只会记录写操作命令，读操作命令是不会被记录的，因为没意义。

在redis中AOF持久化功能默认是不开启的，需要我们修改redis.conf配置中的以下参数：

![image-20240501122853812](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240501122853812.png)

AOF日志文件其实就是普通的文本，我们可以通过`cat`命令查看里面的内容，不过里面的内容如果不知道一定的规则的话，可能会看不懂。

我这里以[set name xiaolin]命令作为例子，redis执行了这条命令后，记录在AOF日志里的内容如下图：

![image-20240501123817070](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240501123817070.png)

试想一下，如果redis每执行一条操作命令，就把该命令以追加的方式加入到每一个文件里，然后重启redis的时候，先去读取这个文件里的命令，并且执行它，这不就相当于恢复了缓存数据了嘛？

参数说明：

「`*3`」表示当前命令有三个部分，每部分都是以「`$+数字`」开头，后面紧跟着具体的命令、键或值。然后，这里的「`数字`」表示这部分中的命令、键或值一共有多少字节。例如，「`$3 set`」表示这部分有 3 个字节，也就是「`set`」命令这个字符串的长度。

不知道大家注意到了没有，Redis是先执行写操作命令后，才将该命令记录到AOF日志里的，这么做其实有两个好处。

1. 第一个好处，**避免额外的检查开销。**
   因为如果先将写操作命令记录到AOF日志里，再执行该命令的话，如果当前的命令语法有问题，那么如果不进行命令语法检查，该错误的命令记录到AOF日志里后 ，Redis在使用日志恢复数据时，就可能会出错。
   而如果先执行写操作命令在记录日志的话，只有在该命令执行成功后，才将命令记录到AOF日志里，这样就不用额外的检查开销，保证记录在AOF日志里的命令都是可执行并且正确的。
2. 第二好处，**不会阻塞当前写操作命令的执行**，因为当写操作命令执行成功后，才会将命令记录到AOF日志。
   当然，AOF持久化功能也不是没有潜在风险。
   第一个风险，执行写操作，命令和记录日志是两个过程，那当redis在还没得及将命令写入到硬盘时，服务发生了宕机，这个数据就会有**丢失风险。**
   第二个风险，前面说道，由于写操作命令执行成功后才记录到AOF日志，所以不会阻塞当前写操作命令的执行，但是**可能会给下一个命令带来阻塞风险。**
   因为将命令写入到日志的这个操作也是在主进程完成的（执行命令也是在主进程），也就是说这两个操作是同步的。

![image-20240501154052241](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240501154052241.png)

#### **RDB：全量备份**

> RDB就是Redis DataBase的缩写，中文名为快照/内存快照，RDB持久化是把当前进程数据生成快照保存到磁盘的过程，由于是某一时刻的快照，那么快照中的值要早于或者等于内存中的值。

**触发方式**

> 触发RDB持久化的方式有两种，分别是手动触发和自动触发

**手动触发**

> 手动触发分别对应的save和bgsave命令

- **save命令：**阻塞当前Redis服务器，直到RDB过程完成为止，对于内存比较大的实例会造成长时间的阻塞，线上环境不建议使用
- **bgsave命令**：Redis进程执行fork操作创建子进程，RDB持久化过程由子进程负责，完成后自动结束。阻塞只发生再fork阶段，一般时间很短。

![image-20240428163556567](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240428163556567.png)

具体流程如下：

- redis客户端执行bgsave命令或者自动触发bgsave命令；
- 主进程判断当前是否存在正在执行的子进程，如果存在，那么主进程直接返回；
- 如果不存在正在执行的子进程，那么就fork一个新的子进程进行持久化数据，fork过程是阻塞的，fork操作完后主进程即可执行其他操作；
- 子进程先将数据写入到临时的rdb文件中，待快照数据写入完成后再原子替换旧的rdb文件；
- 同时发送信号给主进程，通知主进程rdb持久化完成，主进程更新相关的统计信息（info Presitence下的rdb_*相关选项）

**自动触发**

> 在以下4种情况下自动触发

- redis.conf中配置`save m n`,即在m秒内有n次修改时，自动触发bgsave生成rdb文件；
- 主从复制时，从节点要从主节点进行全量复制时也会触发bgsave操作，生成时的快照发送到从节点；
- 执行debug reload命令，如果没有开启ao持久化，那么也会触发gbsave操作；

**redis.conf中配置RDB**

快照周期：内存快照虽然可以通过技术人员手动执行SAVE命令来进行，但生产环境下多数情况都会设置其周期执行条件。

- **Redis中默认的周期新设置**

![image-20240428214659541](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240428214659541.png)

以上三项默认信息设置代表的意义是：

- 如果900秒内有1Key信息发生变化，则进行快照；
- 如果300秒内有10Key信息发生变化，则进行快照；
- 如果60秒内有10000条Key信息发生变化，则进行快照。读者可以按照这个规则，根据自己的实际请求压力进行设置调整。

**其他相关配置**

![image-20240428222440063](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240428222440063.png)

**dbfilename：**RDB文件在磁盘上的名称。

**dir：**RDB文件的存储路径。默认设置为 "./"，也就是redis服务的住目录。

**stop-writes-on-bgsave-error：**上文提到的在快照中进行过程中，主进程照样可以接受客户端的任何写操作的特性，是指在快照操作正常的情况下。如果快照操作出现异常（例如操作系统用户权限不够，磁盘空间写满等等）时，redis就会禁止写操作。这个特性的主要目的是使运维人员在第一时间就发现redis的运行错误，并进行解决。一些特定的场景下。可以需要对这个特性进行配置，这时就可以调整这个参数项。该参数默认情况下值为yes，如果要关闭这个特性，制定即使出现快照错误redis一样运行写操作，则可以将该值更改为no。

**rdbcompressssion：**该属性将在字符串类型的数据被快照到磁盘文件时，启用LZF压缩算法。Redis官方的建议是保持该选项设置为yes，因为"it's almost always a win".

**rdbchecksun：**从RDB快照功能的version 5版本开始，一个64位CRC冗余校验编码会被放置在RDB文件的末尾，以便对整个RDB文件的完整性进行验证。这个功能大概会损失10左右的性能，但获得了更高的数据可靠性。所以如果你的Redis服务需要追求极致的性能，就可以将这个选项设置为no。

**RDB更深入理解**

> 我们通过几个实战问题来深入理解RDB·

- **由于生产环境中我们为redis开辟的内存区域都比较大（例如6GB），那么将内存中的数据同步到硬盘的过程可能就会持续比较长的时间，而实际情况是这段时间Redis服务一般都会收到数据写操作请求。那么如何保证数据的一致性呢？**

RDB中的核心思路是Copy-on-Write，来保证在进行快照操作的这段时间，需要压缩写入磁盘上的数据在内存中不会发生变化。

### 六、过期删除机制

- **定时删除**：定时删除策略的做法是，**在设置key的过期时间时，同时常见一个定时事件，由事件处理器自动执行key的删除操作。**



  **优点**：可以保证过期key会被尽快删除，也就是内存可以被尽快释放。因此定时删除是对内存友好的。
  **缺点：**在过期key比较多的情况下，删除过期key可能会占用相当于一部分CPU时间，在内存不紧张但CPU时间紧张的情况下，将CPU时间用于删除和当前任务无关的过期键上，无疑会对服务的响应时间和吞吐量造成影响。所以，定时删除策略对CPU不友好。


- **惰性删除**：不主动删除过期键，**每次从数据库访问key时，都检测key是否过期，如果过期则 删除该key**。

  **优点：**因为每次访问时，才会检查key是否过期，所以此策略只会使用很少的系统资源，因此，惰性删除策略对CPU时间最友好。
  **缺点：**如果一个key以及过期，而这个key又仍然保留在数据库中，那么只要这个过期key一直没被访问，它所占用的内存就不会释放，造成了一定的内存空间浪费，所以 ，惰性策略对内存不友好

- **定期删除：每隔一段时间[随机]从数据库中取出一定数量的key进行检查，并删除其中过期的key。**



redis是 惰性删除+定期删除

### 七、内存淘汰策略

> 注意内存淘汰策略和过期删除机制不是一回事

### 八、事务机制

### 九、集群策略

### 十、缓存一致性如何解决？

### 十一、缓存问题

用来缓存会产生什么问题，我们应该怎么去避免这些问题

缓存击穿

缓存穿透

缓存雪崩