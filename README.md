git批量操作工具（目前仅支持gitlab）。

支持以下命令：

(注：以下命令示例中，为方便起见，笔者把可执行文件重命名为‘gitb’，并设置了环境变量。)

- clone命令，包括group和repo两种

  - group：

    gitb clone group [group_name]

    ```bash
    #比如想把gitlab上common组下的所有项目clone下来
    gitb clone group common
    ```

  - repo

    - 本地有repo支持的xml文件：

      gitb clone repo  [file_path]

      ```bash
      #本地有repo支持的xml文件，如有个commo.xml，把xml中记录的项目全部clone下来
      gitb clone repo commo.xml
      ```

    - 直接根据gitlab上的repo式的xml文件来clone：

      gitb clone repo [xml文件所在的项目]:[xml的路径]:[项目分支]

      ```bash
      gitb clone repo repo-xml:common.xml:master
      ```

- checkout命令

  ```bash
  #批量创建新分支（进入到之前批量clone后生成的文件夹里）
  gitb checkout -b develop origin/develop
  #批量切换回master分支
  gitb checkout master
  ```

- fetch命令

  ```bash
  #批量fetch
  gitb fetch
  ```

- pull命令

  ```bash
  #批量pull
  gitb pull
  ```

  

注：其他不支持的git命令，请使用原生git命令。



config文件的配置：

```yaml
git:
    type: gitlab        #目前仅支持gitlab
    token:              #你的gitlab账号的private-token，在"Profile Settings"-account中可以看到
    path:               #你的gitlab服务器地址，如：https://gitlab.xxx.com
```
注意：这里使用的yaml第三方库好像不支持注释，所以请不要在config文件中加入注释。//TODO: 更换yaml库，使支持注释
