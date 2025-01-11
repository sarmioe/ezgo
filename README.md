#  EasyGo is a program that makes Git and Mo CloudSpace easy to use.

> The main reason is that I was driven crazy by Git when I used it to submit MoMit, so I decided to build this, which supports automatic synchronization.
>
> But now, Mo Cloud Space is still a concept product, of course you can use other cloud spaces.
>
> Let's see how to use it. You need to download Git and Go first and then configure the environment variables

[简体中文版本](读我.md) [Español Versión](Introducir.md) 

| command                                                      | Function                                                     |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| ezgo -v                                                       | Display version                                              |
| ezgo -h                                                       | Print Help default English                                   |
| ezgo -hzc                                                     | Print Help of Chinese                                        |
| ezgo -hzt                                                     | Print Help of Chinese Traditional                            |
| ezgo -hes                                                     | Print Help of Spanish                                        |
| ezgo -update [version]                                        | Update EasyGo                                                |
| ezgo -clone [URL] [Localpath] -branch--[branchname] -depth--[number] | Clone repo from cloud                                        |
| ezgo -sync [localpath] [URL]                                  | Run sync                                                     |
| ezgo -sync auto [time defualt is second]                      | Auto sync                                                    |
| ezgo -sync incremental                                        | Synchronize only difference files                            |
| ezgo -config                                                  | Configure EasyGo                                             |
| ezgo -env | Automatic environment check |
| ezgo -logs [level] | Output ezgo logs |
| ezgo -logs git | Output git logs |
| ezgo -logs go | Output Go logs |
| ezgo -push [commit] | Commit to remote repository |
| ezgo -pull [branch] | Pull a branch |
| ezgo -checkout [branchname] | Switch branch name |
| ezgo -conflict [way] | Resolve cloud and local conflicts |
| ezgo go build [ARCH] [System] [Output file name] | Compile Go files Only compile Go |
| ezgo go build all [Output file name] | Compile all architecture versions of all operating systems Only compile Go |
| ezgo go init [name] | Initialize Go project |
| ezgo cs sync [name] | Synchronize to Mo cloud space |
| ezgo cs config | Configure Mo cloud account and password |
| ezgo cs create [name] | Create a Mo cloud space |
| ezgo cs build [Language environment] [Output file name] | Specify a language environment and then compile. The environment can be viewed in cs.clauded.modiznodz.llc |
| ezgo cs download [name] [path] | Download files from the cloud to a local directory. Encrypted download by default |
| ezgo cs delete [name] | Delete cloud projects. Automatically erase and overwrite 7 times by default |
| ezgo cs archive [name] | Archive Mo cloud space to Mo cloud storage or download to local |
| ezgo cs fork [URL] [name] | Fork a project from another repository to Mo cloud |
| ezgo cs copy [name] [name2] | Copy a CloudSpace to another empty one |
| ezgo cs depoly [projectname] [environment] | Deploy a project to Mo cloud serverless hosting |

