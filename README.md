# sort-awesome-go-by-stars
[![codecov](https://codecov.io/gh/k-kurikuri/sort-awesome-go-by-stars/branch/master/graph/badge.svg)](https://codecov.io/gh/k-kurikuri/sort-awesome-go-by-stars)

This is a tool to sort the software listed in "awesome go" by the number of stars.  
Output top 5 in order of star number.


## Install
```
$ go get -u github.com/k-kurikuri/sort-awesome-go-by-stars
 ```

## Usage
You can find the name of the content in the [awesome-go](https://github.com/avelino/awesome-go) repository's readme. The command line argument must be set to the name of the content.
![content-name](https://user-images.githubusercontent.com/5502629/84234098-8d8e9b00-ab2e-11ea-80e0-d60a1cbba5ae.png)

### Command Example
if the content name is `"Audio and Music"`
 ```
$ sort-awesome-go-by-stars "Audio and Music"
 ```

if the content name is `"Command Line"`
 ```
$ sort-awesome-go-by-stars "Command Line"
 ```

### Output
 ```
.................................
╔═══════╤═══════════════════════════════════════╤═════════════════════════════════════════════════════════════════════════════════════════════════════════════╗
║ STAR  │              PACKAGE_URL              │                                                 DESCRIPTION                                                 ║
╟━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━╢
║ 10984 │ https://github.com/spf13/cobra        │ cobra - Commander for modern Go CLI interactions.                                                           ║
║ 10229 │ https://github.com/urfave/cli         │ urfave/cli - Simple, fast, and fun package for building command line apps in Go (formerly codegangsta/cli). ║
║ 2324  │ https://github.com/alecthomas/kingpin │ kingpin - Command line and flag parser supporting sub commands.                                             ║
║ 1311  │ https://github.com/chzyer/readline    │ readline - Pure golang implementation that provides most features in GNU-Readline under MIT license.        ║
║ 1311  │ https://github.com/jessevdk/go-flags  │ go-flags - go command line option parser.                                                                   ║
╟━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━╢
║       │                                       │                                                                                                Command Line ║
╚═══════╧═══════════════════════════════════════╧═════════════════════════════════════════════════════════════════════════════════════════════════════════════╝
```
