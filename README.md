# sort-awesome-go-by-stars
[![CircleCI](https://circleci.com/gh/k-kurikuri/sort-awesome-go-by-stars.svg?style=svg)](https://circleci.com/gh/k-kurikuri/sort-awesome-go-by-stars)
[![codecov](https://codecov.io/gh/k-kurikuri/sort-awesome-go-by-stars/branch/master/graph/badge.svg)](https://codecov.io/gh/k-kurikuri/sort-awesome-go-by-stars)

This is a tool to sort the software listed in "awesome go" by the number of stars.  
Output top 5 in order of star number.


## Install
```
$ go get -u github.com/k-kurikuri/sort-awesome-go-by-stars
 ```

## Usage
Example
 ```
$ sort-awesome-go-by-stars "Audio and Music"
 ```

 ```
$ sort-awesome-go-by-stars "Command Line"
 ```

### Pass multiple arguments
 ```
$ sort-awesome-go-by-stars "Third-party APIs" "Web Frameworks"
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
