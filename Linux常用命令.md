#### Linux常用命令

- **sudo 超级用户的意思**
  假如登录的用户使用的是user_fwr，没有创建文件夹权限，使用sudo mkdir就可以创建了，意思是超级用户使用。
  也可以直接切换root用户：sudo su（Switch user）默认是root用户，拥有服务器的最高权限。
- **man  帮助命令**
  - ​	eg：man ls
- **ls  	显示文件**
  - ls [-a]	显示全部文件（包含隐藏文件）
  - ls [-al]	以长列表的方式显示全部文件
  - ls [-alt]	按时间排序
  - ls /home/directory	显示某个路径下的文件列表
- **pwd 	显示当前工作目录**
- **cd	 切换工作目录（类似dos下用法）**
  - cd .. 			返回上一层目录
  - cd ../other  返回上一层目录并进入其他目录
  - cd ~             返回家（home）目录
  - cd /              返回到根目录
  - cd -              两个目录之前切换（相当于电视机的"回看"键）
  - tab              自动补全，如果存在多个相同目录  ，可按两下会显示所用相同目录，再进行选择
- **cat     显示文件内容（全部），还可以用来连接两个或多个文件，形成新的文件**
  - cat doc.txt				 显示doc.txt的内容
  - cat -n doc.txt            显示行号
  - tac doc.txt                 倒序显示文件内容
  - cat doc.txt|more     分屏显示（按空格显示下一屏，按回车显示下一行） 
- **touch** 创建文件命令 
- **echo** 打印回显命令，也可以用来创建命令
  例如：echo "hello word" > dome.txt表示将回显的内容重定向到dome.txt这个文件中，不存在就创建。
- **du**         查看文件大小（如果目录里面有特别大的文件的情况速度会非常慢） (**-h** 表示更人性化的显示)

**vi/vim	文本编辑器**

`vim .vimrc`文件，用于配置vim编辑器

- **一般模式：默认模式（vim一进来的模式），其他模式可通过esc键回到一般模式，可进行选择、复制、粘贴、撤销**

  - i                      在光标前插入文本
  - o                     在当前行的下面插入新行
  - dd                  删除行
  - yy                   复制当前行
  - nyy               复制n行（例如：2yy）
  - p                     粘贴（光标的下一行）
  - np                   粘贴行（例如：2p）
  - u                     撤销上一个操作（类似ctl+z）
  - ctl + r             撤销上一个操作（类似ctl+y）
  - r                      替换当前字符
  - /                      查找关键字
  - ^                     光标移动到行头部
  - $                     光标移动到行尾部
  - ctl + f （forward）           向前翻页查看
  - ctl + b  （backward）         向后翻页查看
  - ctl + u  （up）                   向下翻半页
  - ctl + d   （down）                  向上翻半页
  - gg               跳转到第一行
  - shift + g     跳转到最后一行
  - :50              跳转到对应行
  - 向下查找    /要查找内容（例如：/hello）后面更反斜线C不区分大小写(如：/hello\c)
  - 向上查找    ?要查找内容（例如：?hello）

- **编辑模式：在一般模式中按i，o，r，可编辑文件（进入编制模式），按esc可回到一般模式**

- **命令模式：在一般模式中按":"（进入命令命令模式），可保存修改或退出vi**

  - :w                      保存当前修改
  - :q                       退出
  - :q!                      强制退出，不保存修改
  - :x                       保存并退出，相当于 :wq
  - :set number    显示行号（输入 **set nu**缩写命令也可以）
  - :set nonumber 不显示行号（输入 **set nonu**缩写命令也可以）
  - :!系统命令         执行一个系统命令并显示结果
  - :sh                     切换到命令行，使用ctrl+d切换加vi
  
  ### 基本目录的作用
  
  1. **/bin：**该目录包含可执行的二进制文件（命令），这些命令通常是系统启动和运行所必需的。例如，ls、cp、mv等命令就位于/bin目录下。
  2. /**boot：**该目录包含启动Linux系统所需要的文件，如内核文件和引导加载程序（bootloader）配置文件。
  3. **/dev：**该目录包含设备文件，用于与系统的硬件设备进行交互。例如，硬盘、键盘、鼠标等设备在/dev目录下都用对应的设备文件。
  4. **/etc：**该目录包含系统的配置文件。大多数软件和服务的配置文件都存放在这个目录下，例如网络配置文件、用户账户配置文件等。
  5. **/home：**该目录啥用户的主目录（home directory）所在的位置。每个用户都有一个对应的子目录，用于存放用户的个人文件和配置。
  6. **/lib：**该目录包含系统所需的共享库文件（libraries），这些库文件为二进制文件提供了一些基本的功能支持。
  7. **/lib32和/libx32：**这些目录是在一些特定的Linux系统中使用的，用于存放32位和x32为系统所需的共享库文件。
  8. **/lib64：**该目录是64位系统的共享库文件所在的位置。
  9. **/lost+found：**该目录用于存放文件系统恢复过程中找到的文件碎片和丢失的文件。
  10. **/media：**该目录是可移动媒体设备（如光盘、USB驱动器）挂载的默认位置。
  11. **/mnt：**该目录用于临时改在其他文件系统或网络共享。
  12. **/opt：**该目录用于安装可选软件包（optional packages）。通常，第三方软件包会被安装到这个目录下。
  13. **/proc：**该目录是一个虚拟文件系统，提供了有关系系统内核和运行进程的信息。通过读取/proc目录下的文件，可以获取系统的各种状态和参数。
  14. **/root：**该目录是超级用户（root）的主目录。
  15. **/run：**该目录是在系统引导过程中创建的一个临时文件系统，用于存放运行时数据，如进程ID文件和套接字文件。
  16. **/sbin：**该目录包含系统管理员使用的系统命令，这些命令通常用于系统维护和管理，只有超级用户才能执行这些命令。
  17. **/srv：**该目录用于存放系统提供的服务的数据文件，例如网站的内容、ftp服务器的文件等。
  18. **/snap：**该目录是Snap软件包管理提供的服务的数据文件，例如网站的内容、FTP服务器的文件等。
  19. **/sys：**该目录是一个虚拟文件系统，提供了与系统内核相关的信息和控制接口。
  20. **/tmp：**该目录用于存放临时文件，这些文件在系统重新启动时会被清除。
  21. **/var：**该目录包含可变的数据文件，如日志文件、邮件文件数据文件等。

#### 文件权限

![image-20240806000252421](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240806000252421.png)

**chmod（change mode）修改权限**

- 添加权限：chmod +x hello.sh 表示将hello.sh变成一个可执行文件，
         	    chmod +rw hello.txt，表示将hello.txt变成一个有可读写文件
- 删除权限：chmod -x hello.sh 表示将hello.sh变成一个可执行文件，
         	    chmod -rw hello.txt，表示将hello.txt变成一个有可读写文件

- 给文件同组、其他添加权限：chmod u+x hello.sh（文件所有者添加可执行权限）
                                                    chmod g+x hello.sh（同组用户添加可执行权限）
                                                    chmod o+x hello.sh（其他用户添加可执行权限）
     同时添加：chmod go+x hello.sh（同组用户和其他用户同时添加可执行权限）
- 给文件同组、其他删除权限：chmod u-x hello.sh（文件所有者添加可执行权限）
                                                    chmod g-x hello.sh（同组用户添加可执行权限）
                                                    chmod o-x hello.sh（其他用户添加可执行权限）

**chmod使用数字修改权限**

![image-20240806001527816](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240806001527816.png)

- **chmod 777 hello.sh** ：表示将hello.sh文件所有者、同组用户和其他用户的权限修改为都可读写和执行。
