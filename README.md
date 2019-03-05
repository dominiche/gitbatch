git批量操作工具（目前只支持gitlab）。支持以下命令：
```bash
#根据gitlab上的group，把整个group的项目clone下来
gitb clone group masget4.0-scfs

#根据repo xml文件，把xml中的项目clone下来
gitb clone repo masget4.0-manifest:scfs.xml:master
#如果本地有repo xml文件的话，可以直接根据xml文件来clone：
gitb clone repo scfs.xml

#批量创建新分支（进入到之前批量clone后生成的文件夹里）
gitb checkout -b develop origin/develop
#再批量切回master分支
gitb checkout master

#批量fetch
gitb fetch

#批量pull
gitb pull
```



说一下config文件的配置：

```yaml
git:
    type: gitlab
    token:              #你的gitlab账号的private-token，在Profile Settings-account中可以看到
    path: https://gitlab.masget.com
```



（最后说一下为什么都有批量操作的工具了，如：repo，为什么还要再搞一个？因为repo clone下来的项目不是原生的git项目，git bash不支持，导致ide的git图形化比较工具无法工作。所以就花了点时间造了一个。）

（设置了环境变量后，就可以随处gitb操作了。其他的都是使用原生git命令，完美衔接）