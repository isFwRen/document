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

#### 三、git版本回退

1. 使用**`git reflog`**查看提交操作命令
2. **`git reset --hard commit`**版本ID
3. 永久的丢弃所有的本地改动：

   ```
   git reset --hard
   ```

#### 四、拉取代码操作说明

拉取回来的代码默认是远程的main分支

在本地创建分支并跟远程绑定的方式：

- **`git checkout -b dev origin/dev`**命令说明：在本地创建一个dev分支，并与远程分支的dev分支做绑定。
- 首先**`git checkout -b dev`**，然后通过**`git push --set-upstream origin dev`**进行远程绑定。
  git push --set-upstream <远程主机名(远程仓库名，一般设为origin)> <本地分支名>:<远程分支名>

> 注：代码创建一个dome.go文件的时候，本地的分支都会有显示，当这个文件在dev提交到远程之后，其他才不会显示。

所以在自己分支创建了文件，并测试成功，要先推送到自己的远程分支，然后再拉取dev分支回来，然后再dev分支进行git merge fwr，再解决冲突，然后推送远程dev分支到测试环境进行测试，测试没问题，在本地的main分支进行git pulll拉取最新的main分支代码，然后切换到fwr分支，在自己分支进行merge main分支，合并完成后，将自己的分支推送到远程，然后进行代码上线合并。

实操案例：

1. 拉取远程代码：git clone git@github.com:dex-pert/dex-pert-backend.git
2. 本地默认分支为main分支，设置与远程main分支绑定，**`git push --set-upstream origin main`**
3. cd进入dex-pert-backend目录，默认是main分支，使用**`git checkout -b dev origin/dev`**创建本地dev分支并与远程dev分支进行绑定。
4. 然后切换回main分支，使用**`git checkout -b fwr`**创建自己的分支，代码与main的代码相同，然后使用git **`push --set-upstream origin fwr`**将自己的分支推送到远程上并在远程上创建自己的远程分支。
5. 在自己的分支进行代码开发，本地测试没问题了，将自己的代码推送到自己的fwr远程分钟，本地切换到dev分支，使用git pull，拉取最新的dev分支代码，然后git merger fwr将自己的分支合并到测试分支，然后推送到远程dev分支进行测试。
6. 测试环境通过，本地切换到main分支，使用git pull拉取最新的main分支代码，获取最新代码后，切换到自己的分支，然后使用 git merge main将主分支合并到自己的分支，如果有冲突，解决冲突再推送到远程自己分支。

大概这样一个开发流程。

> 注意：
>
> 在main分支进行git checkout -b fwr进行分支创建，fwr分支的代码是main分支的代码，如果在dev分支进行git checkout -b fwr，那么fwr分支的代码就是dev分支的代码



