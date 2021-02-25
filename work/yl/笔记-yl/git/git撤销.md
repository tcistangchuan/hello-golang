##### git中文件提交错误后的回滚(https://www.cnblogs.com/bianruoyu/p/9566085.html)

**情况1:文件已作出修改 但是没进行 add操作 想要还原文件**

git checkout 指定的文件

git checkout .  (还原全部文件)

 

**情况2:文件作出修改 已进行add操作 但是没有 commit  想要删除add**

git reset HEAD   撤销全部已提交修改

git reset HEAD filename  撤销对指定文件的修改

 

**情况三:文件作出修改 已进行过 commit 操作但是没有push  想要删除commit**

git log 查看节点 
commit xxxxxxxxxxxxxxxxxxxxxxxxxx 
Merge: 
Author: 
Date:

然后 
git reset commit_id

 

**情况四:文件作出修改已push到仓库**  

此次操作之前和之后的commit和history都会保留，并且把这次撤销作为一次最新的提交 
git revert HEAD 撤销前一次 commit 
git revert HEAD^ 撤销前前一次 commit 
git revert commit-id (撤销指定的版本，撤销也会作为一次提交进行保存） 
git revert是提交一个新的版本，将需要revert的版本的内容再反向修改回去，版本会递增，不影响之前提交的内容。

也可以使用reset 

git reset --hard HEAD^   reset是指将HEAD指针指到指定提交，历史记录中不会出现放弃的提交记录。