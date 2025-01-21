# ezgoit A small tool that makes Git and Go easy to use

> This tool makes Go and Git easy to use. Just do it like this: `ezgo -fmt` to directly format the Go language code
>
> Compared with Git and Go, which require remembering complex commands and a lot of interactions, ezgoit makes them simpler.
>
> ezgoit comes from: Easy+Go+Git. When it is actually running, it is simplified to ezgo for easy operation.
>
> Open source tool, conditionally open source based on GPLv3 license.

[简体中文版本](读我.md)

## Let's see how to use it

### Basic command list:

| Command | Explanation |
| --------- | ------------------------------------------------------------ |
| ezgo -v | Output version |
| ezgo -u | Update version online |
| ezgo -un | Update version offline, need to specify the source package location |
| ezgo -h | Output help (simplified version) |
| ezgo -c | Configure ezgoit so that it does not need to specify parameters each time it runs. Of course, the function of using bat or shell script is the same |
| ezgo -env | Check Git and Go language environment and version |
| ezgo -l | Output log |

### Go command list:

| Command | Explanation |
| -------- | -------------------------------- |
| ezgo -g | Check Go environment |
| ezgo -gi | Initialize Go project |
| ezgo -b | Compile Go project |
| ezgo -ba | Compile all versions of the environment supported by Go language |
| ezgo -f | Format Go project |
| ezgo -ge | Download dependencies |
| ezgo -t | Test code |
| ezgo -m | Generate go.mod file |

### Git command list:

| Command | Explanation |
| --------- | ----------------------------- |
| ezgo -gic | Configure Git username and password |
| ezgo -gi | Initialize a Git repository |
| ezgo -gc | Clone the repository to local |
| ezgo -cm | Commit changes with one click |
| ezgo -cmt | Cycle git add . and then commit changes with one click |
| ezgo -ch | Switch branches |
| ezgo -br | Create,list and delete branches and more. |
| ezgo -mr | Merge branches |