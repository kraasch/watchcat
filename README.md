
# Watchcat

Directory structure monitor and fix dispatcher (TUI and CLI). 

<p align="center">
  <img src="./resources/example.png" width="200"/>
</p>

The purpose of watchcat is define, restrict and fix the layout of a directory tree.
Where git is used to organize single projects, the watchat can be used to manage the overview over several projects and to glue several projects or directories together.

Watchcat lets you define restrictions, watchcat then ...

  - reports failure to keep the layout,
  - suggests commands to fix the layout,
  - executes fixes automatically,
  - can execute different fix strategies (`shred` vs `rm` or move them away with `mv`),
  - lets you choose which fixes to run and which ones to ignore,
  - can be scheduled to report or fix.

## Demo

Demo picture:

  - Coming...

<!-- TODO: add demo.
<p align="center">
  <img src="./resources/example.png" width="300"/>
</p>
-->

Example of a bad directory graph.

```text
  ______/______
./ ./  / \    \
      /\  \.  /\.
     /  \   ./
   ./\.  \.
```

Watchconf example.

```text
ezN{}  .                   : Comment.
z      ./folderA/          : Comment.
ez     ./folderA/subA/     : Comment.
ez     ./folderA/subB/     : Comment.
zs     ./folderB/          : Comment.
e      ./folderB/fileA.txt : Comment.
z      ./folderB/fileB.txt : Comment.
ez     ./folderC/          : Comment.
       ./folderD/          :[t(txt,srt,sub,mp4)] Comment.
t(mp3) ./folderE/          : Comment.
```

Example of a watchcat comamnd:

`watchcat --execute`

Example of a good directory graph after running watchcat:

```text
         /\
  ______/  \____________
./ ./ ./    \. \. \. \. \.
   .
```

## Features

<!--
Goals:
  - [ ] Directory structure monitor and fix dispatcher (TUI and CLI). 
  - [ ] The purpose of watchcat is define, restrict and fix the layout of a directory tree.
  - [ ] Where git is used to organize single projects, the watchat can be used to manage the overview over several projects and to glue several projects or directories together.
  - [ ] report if a git repo is in wrong depth of a repo.
  - [ ] report if changes are uncommited in git.
  - [ ] enforce upper bound to leaf depth (eg no folders deeper than 3).
  - [ ] enforce flat directory tree hierarchy.
  - [ ] enforce evenly formed branches.
  - [ ] have reports.
  - [ ] have alert list of things which need fixing.
  - [ ] show fix suggestions as bash command list.
  - [ ] show fix suggestions which watchcat can run, let user choose.
  - [ ] give user a way to pre-defined which alert should be fixed how (ask user or auto fix).
-->

List of basic features:

  - [ ] define a layout.
  - [ ] report failure to keep the layout.
  - [ ] suggests commands to fix the layout.
  - [ ] executes fixes automatically.
  - [ ] can execute different fix strategies (`shred` vs `rm` or move them away with `mv`).
  - [ ] give user choice which fixes to run and which ones to ignore.
  - [ ] can be scheduled to report or fix.
  - [ ] have a report of alerts which have to be done in order to fix layout.
  - [ ] for each alert provide different ways of fixing.
    - [ ] unwanted files: delete, shred, move away.
    - [ ] uncommited git repos: commit, checkout locally, checkout remote.
    - [ ] unpushed git repos: push, checkout locally, checkout remote.
    - [ ] git without remote: add remote.
    - [ ] big files: shred, remove or move away.
    - [ ] big directories: shred, remove or move away.
    - [ ] wrong names: manual rename, auto rename, remove, shred.
    - [ ] required nodes: create watchconf, create git repo, create directory, create files.
  - [ ] have command to check, report and fix rules, ie. `watchcat`.
  - [ ] have command to add directories as target entries to a watchcat config file, ie. `wcat`.
    - [ ] for example `wcat add ./Downloads -name 'downloads'`.
    - [ ] for example `wcat remove downloads` or `wcat remove ./Downloads`.
    - [ ] ask location of watchcat config (global or local).
    - [ ] ask for name of target.

Features for later:

  - [ ] create pre-defined watchcat configs and rules for certain repo types.
    - [ ] eg media store for videos and music, categorized media store, projects of restricted complexity.
  - [ ] Recursively checks for `.git` directories and uncommited changes (or other unsafe git repo states like not having any remotes configured).
  - [ ] Check if directories have backup.
  	- [ ] Git repos can be compared with their remote.
  	- [ ] Other directories can use **rsync**.
  - [ ] Enforce certain leave types on directory tree.
  	- [ ] It is not allowed for a subdirectory just end, without having a declared type.
  	- [ ] And if the type of a subdirectory does not match its declared type, this needs to result in an alert (depending on a policy).
  - [ ] Require linters to be run without reports.
    - [ ] Fix: Automatically run linter's solution on files.
  - [ ] Replace https://github.com/kraasch/git-walker with watchcat.
  - [ ] Allow any sort of script to cause an alert for a watchat node.
  - [ ] Allow any sort of script to make a fix for a watchat node.

Maybe:

  - [ ] feature: check for any `TODO` tags in a software repository (if yes, use external package).
  - [ ] feature: give watchcat the ability to check for the existence of certain git remote addresses, ie if remote is https-based or ssh-based.

Done:

  - [ ] xxx

### Planned Operand Codes

Basic:

  - [ ] `a(dyfzh;min/max)` -- requires things (dir/file/symlink/hidden) to have names of minimum or maximum length or both.
  - [ ] `d` -- requires there be no directories in this directory.
  - [ ] `e` -- requires directory to be empty, ie nothing at all (dir/files/sym/hidden).
  - [ ] `f` -- requires there be no files in this directory.
  - [ ] `g(xcpr)` -- requires the directory to have a Git repo (distinguish: repo exists, uncommit changes exist, unpushed changes exist, remote repo exists).
  - [ ] `h(dyfzh)` -- requires there to be no hidden things (dir/file/symlink) in this directory.
  - [ ] `i(min/max)` -- requires directory to be of minimum or maximum size on disk or both.
  - [ ] `l(dyfzh)` -- requires names to only use lowercase letters (distinguish: files, directories, symlinks, hidden).
  - [ ] `m(dyfzh;min/max)` -- requires there to be things (dir/file/symlink/hidden) only from min to max tree depth.
  - [ ] `n(f;A-z,...)` -- requires thing names not to have the specified characters. See more info below.
  - [ ] `r(dyfzh;regex)` -- requires thing names (dir/file/symlink/hidden) to be like regex.
  - [ ] `s(min/max)` -- requires files to be of minimum or maximum size or both.
  - [ ] `t(type,...)` -- requires to be only to be of the specified file types.
  - [ ] `u(dyfzh)` -- requires names to only use uppercase letters (distinguish: files, directories, symlinks, hidden).
  - [ ] `w` = requires directory to have a Watchconf.
  - [ ] `x` -- requires for the directory to exist.
  - [ ] `y(hs)` -- requires there be no symlinks in this directory (distinguish: hard and soft).

Other:

  - [ ] `!` -- forbid everything except allowed things (permissions through uppercase letter specifications).
  - [ ] `*` -- inherit all operands from parent.
  - [ ] `.` -- require directory to exist.
  - [ ] `?` -- requires the directory tree under this directory not have any rules specifications, not in config, nor Watchconf files.
  - [ ] `b(dyfzh)` -- requires there to be no duplicated things (dir/file/symlink/hidden).
  - [ ] `p(dyfzh)` -- requires there to be no duplicate thing names (dir/file/symlink/hidden).

Combinations and shortcuts:

  - [ ] `z` == `dyf` (directories/symlinks/files, not hidden ones).
  - [ ] `h` == `dyf` (directories/symlinks/files, including hidden ones).
  - [ ] `{from-to;categories}`, special syntax which only allows certain categories at certain depth of the directory tree.
    - [ ] for example `{0-1d}` -- in this directory and below only directories, i.e. no files.
    - [ ] for example `{0-1d}{2gw}` -- in the directory level below only git repos or further Watchconf files.
    - [ ] for example `{0-1d}{2f}`  -- in the directory level below only files.
    - [ ] for example `{0-1d}{2c}` -- in the directory level below git repos with uncommitted changes, but not without remotes.
    - [ ] restrict categories to : `xcprfdyzwg` or something useful.

Notes:

  - [ ] `N` and `n`:
    - [ ] distinguish: files, directories, symlinks, hidden.
    - [ ] specify groups: `A-Z`, `a-z`, `A-z`, `0-9`, etc.
    - [ ] specify special groups: `<whitespace>`, `<alphanum>`, etc.
    - [ ] specify single characters: `abcABC!?:;,.-_=`, etc.
    - [ ] specify special characters: `<space>`, `<tab>`, etc.
  - [ ] `T` and `t`:
    - [ ] specify pre-defined single types: `jpg` (for jpeg and jpg), `png`, etc.
    - [ ] specify pre-defined type groups: `<img>`, `<vid>`, `<aud>`, `<txt>`, `<bin>`, etc.
    - [ ] tested with `file` command on linux (not by file ending, which is done by the `r()` operand.

Done:

  - [ ] xxx

## Tasks

Next:

  - [ ] xxx

Later:

  - [ ] allow definition of operands on the right side, by following the
        comment `:` with angular brackets, ie `some/dir :[xxxx] Some comment.`.
  - [ ] extract `./pkg/gocfg/` into its own repo at `github.com/kraasch/gocfg`.
  - [ ] sort through **Ideas** section and delete it.
  - [ ] throw an error if there is no Watchconf at the target's root (except if
        rules are defined in watchcat's global configuration).
    - [ ] maybe allow recursive search for Watchconf after setting some toggle
          in watchcat's global configuration, eg. `ruleLocation = "recursive"`.

Think about:

  - [ ] are rules to files allowed, ie. `op ~/downloads/abc.txt : some file` and
        what operands are allowed on files?

Done:

  - [ ] xxx

## Installation

Install the program:

```bash
go install github.com/kraasch/watchcat@latest
```

Get the package:

```bash
go get github.com/kraasch/watchcat
```

## Usage

Use the program:

```bash
git clone github.com/kraasch/watchcat
cd watchcat/
make build
./build/watchcat -help
```

Use the package:

```go
import (
  "github.com/kraasch/watchcat"
)

watchcat.DoSomething("Hello")
```

## Mini Glossary

Watchcat defines layout rules for directory trees.
Each directory tree is a target.
All targets are listed in a watchcat configuration file.
Rules can be defined in a watchcat configuration file or in a separate Watchconf
file within the target directory.

  - **watchcat** -- the program.
  - **target** -- a subdirectory under watchcat's control.
  - **watchcat** config file -- a configuration TOML specifying targets.
  - **Watchconf** file -- a local configuration text file referenced from within the watchcat config.
  - **rules** -- lines of text of the form operands-path-comment, they can be lines of a Watchcat file or within a subsection of the watchcat config file.
  - **operands** codes (op codes or ops) -- restrictions for certain paths or directory trees.
  - **alert** -- a rule being broken.
  - **fix** -- a proposed solution to create the states which fulfil all restrictions (bash script or watchcat action).

```text
OPS    | PATH               | COMMENT or                   |
PREFIX |                    | COMMENT with OPS SUFFIX      |
-------+--------------------+------------------------------+--------
ezN{}   .                   : Comment.                     | RULE
        ./folderD/          :[t(txt,srt,sub,mp4)] Comment. | RULE
                                                           | ...
```

## Related Projects

  - A Go implementation of Git: https://github.com/go-git/go-git

## Feedback

I can be reached via [alex@kraasch.eu](mailto:alex@kraasch.eu).

## Contributing

Feel free to help me.

## Acknowledgments

Uses the following software:

  - see [go.mod](./go.mod) and [go.sum](./go.sum).

Made by the following people:

  - see Github info.

## License

View the [license file](./LICENSE).

