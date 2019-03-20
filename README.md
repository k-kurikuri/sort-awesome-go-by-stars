# sort-awesome-go-by-stars
This is a tool to sort the software listed in "awesome go" by the number of stars.  
Output top 5 in order of star number.


## Install
execute the following command.  
This package uses [dep](https://github.com/golang/dep), please install it

```
$ git clone git@github.com:k-kurikuri/sort-awesome-go-by-stars.git && cd sort-awesome-go-by-stars
$ dep ensure
$ go build
 ```

## Usage
Example
 ```
$ ./sort-awesome-go-by-stars "Audio and Music"
 ```

 ```
$ ./sort-awesome-go-by-stars "Command Line"
 ```

### Pass multiple arguments
 ```
$ ./sort-awesome-go-by-stars "Third-party APIs" "Web Frameworks"
 ```

### Output
 ```
.................................
Command Line
+-------+---------------------------------------+--------------------------------+
| STAR  |              PACKAGE URL              |          DESCRIPTION           |
+-------+---------------------------------------+--------------------------------+
| 10980 | https://github.com/spf13/cobra        | cobra - Commander for modern   |
|       |                                       | Go CLI interactions.           |
| 10224 | https://github.com/urfave/cli         | urfave/cli - Simple, fast,     |
|       |                                       | and fun package for building   |
|       |                                       | command line apps in Go        |
|       |                                       | (formerly codegangsta/cli).    |
|  2323 | https://github.com/alecthomas/kingpin | kingpin - Command line and     |
|       |                                       | flag parser supporting sub     |
|       |                                       | commands.                      |
|  1311 | https://github.com/chzyer/readline    | readline - Pure golang         |
|       |                                       | implementation that provides   |
|       |                                       | most features in GNU-Readline  |
|       |                                       | under MIT license.             |
|  1311 | https://github.com/jessevdk/go-flags  | go-flags - go command line     |
|       |                                       | option parser.                 |
+-------+---------------------------------------+--------------------------------+
```
