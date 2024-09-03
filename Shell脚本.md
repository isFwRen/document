#### Shell脚本

##### 1、**创建shell脚本**

使用vi/vim hello.sh创建shell脚本，shell脚本的后缀名以.sh结尾。

![image-20240806152120468](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240806152120468.png)

shell脚本首行以**`#!`**开头，后面跟着的是这个脚本文件使用哪个版本的解释器(路径可以换成想要的解释器版本)

**查看解释器命令**：**`cat /etc/shells`**

![image-20240806152450433](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240806152450433.png)

**查看当前使用的解释器：**echo $SHELL（打印当前shell版本的环境变量）

![image-20240806152643582](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240806152643582.png)

**2、创建好脚本后，进行简单的编写**

![image-20240806153007044](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240806153007044.png)

编写成功后，使用**`ls -ltr`**查看文件信息

![image-20240806153144455](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240806153144455.png)

会发现创建好的脚本只有读写权限，并没有执行的权限，这样脚本时不能运行的，接下来就要给脚本添加可执行命令:**`chmod a+x hello.sh`**（命令解释：a表示文件所有者u、同组用户g和其他用户o，表示这个脚本文件的三组用户都添加了可执行权限），添加权限后在查看文件信息：

![image-20240806153731477](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240806153731477.png)

会发现横杠换成了x，其中hello.sh的颜色也发生了改变。

> 小技巧：如果不给这个脚本添加可执行权限，也想执行这个脚本的话，使用`zsh hello.sh`,使用zsh可以执行没有执行权限的脚本

**3、编写一个判断素数案例**

![image-20240806154612412](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240806154612412.png)

```shell
#!/bin/bash

# 函数：判断一个数是否是素数
is_prime() {
    local num=$1	#local关键字表示局部变量，只在函数内部使用

    # 1 不是素数
    if [ "$num" -le 1 ]; then  #if判断的中括号[]的两边都要有一个空格，不然会出现语法错误
        echo "False"
        return 1
    fi

    # 2 是素数
    if [ "$num" -eq 2 ]; then
        echo "True"
        return 0
    fi

    # 如果是偶数且不等于2，不是素数
    if [ $((num % 2)) -eq 0 ]; then
        echo "False"
        return 1
    fi

    # 检查从3到sqrt(num)之间的奇数
    for ((i=3; i*i<=num; i+=2)); do
        if [ $((num % i)) -eq 0 ]; then
            echo "False"
            return 1
        fi
    done

    echo "True"
    return 0
}

# 读取输入
read -p "请输入一个数字: " number

# 调用函数并显示结果
if [ "$(is_prime $number)" == "True" ]; then
    echo "$number 是素数"
else
    echo "$number 不是素数"
fi

```

![image-20240806160131552](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240806160131552.png)

![image-20240806161059207](https://fwr-typora-img.oss-cn-guangzhou.aliyuncs.com/typroa-img/image-20240806161059207.png)