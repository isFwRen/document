### Linux常用命令

- **man  帮助命令**
  - ​	eg：man ls

- **ls  	显示文件**
  - ls [-a]	显示全部文件
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

**vi/vim	文本编辑器**

- **一般模式：默认模式，其他模式可通过esc键回到一般模式，可进行选择、复制、粘贴、撤销**

  - i                      在光标前插入文本
  - o                     在当前行的下面插入新行
  - dd                  删除行
  - yy                   复制当前行
  - n+yy               复制n行
  - p                     粘贴
  - u                     撤销上一个操作
  - r                      替换当前字符
  - /                      查找关键字

- **编辑模式：在一般模式中按i，o，r，可编辑文件，按esc可回到一般模式**

- **命令模式：在一般模式中按":"，可保存修改或退出vi**

  - :w                      保存当前修改
  - :q                       退出
  - :q!                      强制退出，不保存修改
  - :x                       保存并退出，相当于 :wq
  - :set number    显示行号
  - :!系统命令         执行一个系统命令并显示结果
  - :sh                     切换到命令行，使用ctrl+d切换加vi

  

  