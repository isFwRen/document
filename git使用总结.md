#### 一、git如何解除与绑定远程仓库关联

**解除**

有时候我们需要将一个项目上传到另一个远程仓库时，那么就需要解除原来的仓库关联。

- 首先切换到项目的根目录，查看项目原有的remote

```
git remote -v

#结果
origin https://github.com/isFwRen/goframe-shop-v2.git
```

- 接下来就是解除与原来远程仓库的关联

```  
git remote rm "remote名称"

//例如
git remote rm origin
```

- 取消git初始化

```
rm -rf .git
```

除了手动修改配置文件外, 也可以使用git命令, 效果和手动修改没区别

```csharp
csharp
复制代码git config --global init.defaultBranch main
```

以上方法只是让以后创建的项目默认分支为main, 但对于已经创建的项目则无能为力, 所以我们还需要对已存在的项目逐个进行修改.

**绑定远程仓库**

```
git remote add origin(远程仓库别名) 远程仓库地址

例子：
git remote add arigin git@github.com:isFwRen/goframe-shop-v2.git
```

修改远程从仓库名称

```
git remote rename arigin(旧远程仓库别名) origin(新远程仓库别名)
```

```
git push -u origin main

在git push -u origin master 命令中，-u 参数的含义是将本地分支与远程分支进行关联（或称为设置上游分支）
当没有用-u参数时，在使用git pull时会报错

git pull origin main --allow-unrelated-histories
这样可以解决不是历史根目录问题
```

有个小技巧：如果是本地绑定远程仓库的话，创建远程仓库时，最好都不添加文件，创建空仓库就行，这样推送方便很多

#### 二、修改已创建项目的主分支为main

1. 切换到主分支master
2. 使用`git branch -M main`命令, 把当前master分支改名为main, 其中`-M`的意思是移动或者重命名当前分支

除了手动修改配置文件外, 也可以使用git命令, 效果和手动修改没区别

```csharp
git config --global init.defaultBranch main
```

以上方法只是让以后创建的项目默认分支为main, 但对于已经创建的项目则无能为力, 所以我们还需要对已存在的项目逐个进行修改



