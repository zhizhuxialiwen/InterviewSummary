# Git提交代码的流程——新手适用

pull：是下拉代码，相等于将远程的代码下载到你本地，与你本地的代码合并
push：是推代码，将你的代码上传到远程的动作
完整的流程是：

1. 工作区：更新本地仓库

`$git fetch origin master(远程):master(本地)` 
`$git merge origin/master(远程):master(本地)`

或
 `git pull origin master`

 例如，
 `$git branch -a`
`$git fetch origin main` 
`$git merge origin/main`

2. 工作区：显示有变更的文件

git status命令用于显示工作目录和暂存区的状态。使用此命令能看到那些修改被暂存到了, 哪些没有, 哪些文件没有被Git tracked到。git status不显示已经commit到项目历史中去的信息。看项目历史的信息要使用git log.//原

`$ git status`

* 查看工作区修改文件

显示工作区和暂存区的差异，工作区修改的数据没有提交暂存区。
`$ git diff`

显示工作区与当前分支最新commit之间的差异
`$ git diff HEAD`


* 将工作区修改文件添加暂存区

`$ git add .`

* 显示暂存区和上一个commit的差异，工作区修改的数据已经提交暂存区

`$ git diff --cached [file]`

3. 暂存区添加到本地仓库

`$ git commit -m [message]`
或`$ git commit -e`

4. 本地仓库：查看提交日志

`$ git log`

5. 本地仓库上传远程仓库

`$ git push [remote] [branch]`
例如：`git push -u origin master`
-u表示第一次将本地仓库推送到远程仓库，下次默认连接远程仓库

例如`git push origin wen`
