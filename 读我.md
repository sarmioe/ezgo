#  EasyGo 让Git和Mo云空间便于使用的小程序

> 主要还是我在使用Git提交MoMit的时候被Git给逼疯了,决定造个这,支持自动同步哦
>
> 但是现在,Mo云空间还是个赛博概念,当然你可以用别的云空间
>
> 来看看怎么用吧 需要先下载Git和Go然后配置环境变量

| 命令                                                         | 功能                                                         |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| ezgo v                                                       | 显示版本                                                     |
| ezgo h                                                       | 输出帮助 默认英文                                            |
| ezgo h zh_CN                                                 | 输出简体中文的帮助                                           |
| ezgo h zh_TW                                                 | 输出繁体中文的帮助                                           |
| ezgo h es                                                    | 输出西班牙文的帮助                                           |
| ezgo update [version]                                        | 更新EasyGo                                                   |
| ezgo clone [URL] [Localpath] -branch--[branchname] -depth--[number] | 从远端直接克隆仓库                                           |
| ezgo sync [localpath] [URL]                                  | 执行同步 URL是远程仓库的 如果将ezgo和项目置于一个文件夹 就把Localpath位置设置为./即可 |
| ezgo sync auto [time defualt is second]                      | 指定时间后自动检测更改和同步                                 |
| ezgo sync incremental                                        | 只同步差异文件 用git add . + git status实现                  |
| ezgo config                                                  | 配置EasyGo 包括云端ssh密钥                                   |
| ezgo env                                                     | 自动环境检查                                                 |
| ezgo logs [level]                                            | 输出ezgo日志                                                 |
| ezgo logs git                                                | 输出git日志                                                  |
| ezgo logs go                                                 | 输出Go日志                                                   |
| ezgo push [commit]                                           | 提交到远端存储库                                             |
| ezgo pull [branch]                                           | 拉取一个分支                                                 |
| ezgo checkout [branchname]                                   | 切换分支名称                                                 |
| ezgo conflict [way]                                          | 解决云端和本地冲突                                           |
| ezgo go build [ARCH] [System] [Output file name]             | 编译Go文件 只能编译Go                                        |
| ezgo go build all [Output file name]                         | 编译所有操作系统的所有架构版本 只能编译Go                    |
| ezgo go init [name]                                          | 初始化Go项目                                                 |
| ezgo cs sync [name]                                          | 同步到Mo云空间                                               |
| ezgo cs config                                               | 配置Mo云的账号密码                                           |
| ezgo cs create [name]                                        | 创建Mo云空间                                                 |
| ezgo cs build [Language environment] [Output file name]      | 指定一个语言环境 然后编译 环境可以前往cs.clauded.modiznodz.llc查看 |
| ezgo cs download [name] [path]                               | 从云端下载文件到本地目录 默认加密下载                        |
| ezgo cs delete [name]                                        | 删除云端项目 默认自动7次擦除覆盖                             |
| ezgo cs archive [name]                                       | 归档Mo云空间到Mo云存储或下载到本地                           |
| ezgo cs fork [URL] [name]                                    | 从其他存储库里fork项目到Mo云                                 |
| ezgo cs copy [name] [name2]                                  | 复制一个CloudSpace到另一个空的                               |
| ezgo cs depoly [projectname] [environment]                   | 部署一个项目到Mo云无服务器托管                               |

