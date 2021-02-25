

#### #git cherry-pick xxx(commit_id)

cherry-pick是Git里对commit操作很好的一个指令，比如想把test分支中的其中多个commit merge到master，那么你需要挑你所需要的commit合到master，这时候就用cherry-pick来捡。

```
git cherry-pick -x <commit_id>
增加 -x 参数，表示保留原提交的作者信息进行提交。

在 Git 1.7.2 版本开始，新增了支持批量 cherry-pick ，就是可以一次将一个连续的时间序列内的 commit ，设定一个开始和结束的 commit ，进行 cherry-pick 操作。
git cherry_pick <start-commit-id>…<end-commit-id>

可以看到，它的范围就是 start-commit-id 到 end-commit-id 之间所有的 commit，但是它这是一个 (左开，右闭] 的区间，也就是说，它将不会包含 start-commit-id 的 commit。
而如果想要包含 start-commit-id 的话，就需要使用 ^ 标记一下，就会变成一个 [左闭，右闭] 的区间，具体命令如下。
git cherry-pick <start-commit-id>^...<end-commit-id>
```





#### #git rebase （变基 衍合）

```
参考：http://jartto.wang/2018/12/11/git-rebase/
(1)合并多次提交纪录:
git rebase -i parentCommitID
-i弹出的文本编辑器，将要修改的commit的前方的pick改为s，保存。
  
  文本编辑器里面的提示有：
  pick：保留该commit（缩写:p）
  reword：保留该commit，但我需要修改该commit的注释（缩写:r）
  edit：保留该commit, 但我要停下来修改该提交(不仅仅修改注释)（缩写:e）
  squash：将该commit和前一个commit合并（缩写:s）
  fixup：将该commit和前一个commit合并，但我不要保留该提交的注释信息（缩写:f）
  exec：执行shell命令（缩写:x）
  drop：我要丢弃该commit（缩写:d）
  
如果你异常退出了 vi 窗口，不要紧张：
git rebase --edit-todo
这时候会一直处在这个编辑的模式里，我们可以回去继续编辑，修改完保存一下：
git rebase --continue
完成，查看结果:
git log --graph

(2) 同步/合并分支
例如在master切个开发分支feature1,开发中其他人往master上提交了代码，你要同步master的代码
若用 (feature1) git merge master  会污染分支，
则使用（feature1） git rebase master 。

rebase 做了什么操作呢？
首先，git 会把 feature1 分支里面的每个 commit 取消掉；
其次，把上面的操作临时保存成 patch 文件，存在 .git/rebase 目录下；
然后，把 feature1 分支更新到最新的 master 分支；
最后，把上面保存的 patch 文件应用到 feature1 分支上

rebase起了冲突怎么解决？
git rebase过程中也许会出现冲突conflict。在这种情况，git会停止 rebase 并会让你去解决冲突。在解决完冲突后，用 git add 命令去更新这些内容，无需执行git commit。只需执行git rebase --continue, 这样 git 会继续应用余下的 patch 补丁文件。最后执行git push -f

如何终止git rebase?并且分支会回到 rebase 开始前的状态。
git rebase --abort
```



#### # git常用操作

```
还未add的操作:
# 查看本地修改但未add的修改的文件
git status -s 
# 查看本地修改但未add的修改的文件的详情
git diff 文件路径
# 撤销未add的修改
git checkout 文件路径
# add文件
git add 文件路径

已经add的文件未commit操作：
# 查看本地add但未commit的修改的文件:
git status -s 或者 git status
# 查看本地add但未commit的修改的文件的详情:
git diff --cached 文件路径
# 撤销未commit的修改:
git checkout head 文件路径
# commit文件:
git commit -m '备注' 文件路径

已经commit未push的操作：
# 查看已经commit的代码详情和某个分支作对比
git diff origin/xx xx（HEAD）
# git reset，回退到某个commit_id
git reset --hard head 回退到当前commit的版本 清除工作区和暂缓区代码
git reset head^ 回退到上一个版本
git reset commit_id 
其实就是--soft 、--mixed以及--hard是三个恢复等级。使用--soft就仅仅将头指针恢复，已经add的缓存以及工作空间的所有东西都不变。如果使用--mixed(git reset默认)，就将头恢复掉，已经add的缓存也会丢失掉，工作空间的代码什么的是不变的。如果使用--hard，那么一切就全都恢复了，头变，aad的缓存消失，代码什么的也恢复到以前状态。
# 撤销某次commt_id
git revert commit_id
# 查看某次提交所修改的文件
git show --stat commit_id
git show commit_id [文件路径] 查看某次提交某个文件的修改详情

已经push的操作
# 撤销/回退
git reset 或者 git revert , 再 git push --force
```



CREATE TABLE `academic_promotion` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `code` varchar(11) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '批号',

 delivery_materials_url varchar(11) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '交付材料url',

execution_details_url varchar(11) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '执行明细url',

  `buy_time` date NOT NULL COMMENT '购买日期',

  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

