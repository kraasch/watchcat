
# Watchcat

Directory structure monitor and fix dispatcher (TUI and CLI). 

<p align="center">
  <img src="./resources/example.png" width="200"/>
</p>

The purpose of watchcat is define, restrict and fix the layout of a directory tree.

Watchcat lets you define restrictions, watchcat then ...

  - reports failure to keep the layout,
  - suggests commands to fix the layout,
  - executes fixes automatically,
  - can execute different fix strategies (`shred` vs `rm`),
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

List of features

  - [ ] define a layout.
  - [ ] report failure to keep the layout.
  - [ ] suggests commands to fix the layout.
  - [ ] executes fixes automatically.
  - [ ] can execute different fix strategies (`shred` vs `rm`).
  - [ ] give user choice which fixes to run and which ones to ignore.
  - [ ] can be scheduled to report or fix.

Done:

  - [ ] xxx

## Tasks

Next:

  - [ ] xxx

Later:

  - [ ] require paths to directories in rules to end in `/` (similar to
        gitignore files).
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

## Ideas // TODO: remove later.

Tasks and purpose:

  - [ ] keep overview over files by white-listing everything
    - [ ] what state is a directory or project in.
    - [ ] have everything known white-listed, such that new things appear as new.
  - [ ] A collection of constraints to the file system, with the following features:
    - [ ] Throws errors (or alerts, eg system popups, messages, etc) when something breaks the constraints.
    - [ ] Direct remote calendar support:
    	- Eg: Automatically pop-up the errors as to-do items in CalDav, WebDav, iCal, etc (?).
    	- Let user customize which event types become to-do lists.
    - [ ] Constraints can be added, edited and removed.
    - [ ] Solutions for broken constraints can be run automatically.
    - [ ] Recursively checks for `.git` directories and uncommited changes (or other unsafe git repo states).
    - [ ] Check if directories have backup.
    	- [ ] Git repos can be compared with their remote.
    	- [ ] Other directories can use **rsync**.
    - [ ] Distinguish remote and local subdirectories. 
    	- [ ] Have different policies for both types or repos.
    	- [ ] For example remotes can have the policy of resulting strict TODOs or non-strict TODOs, independent of the policy of local repos.
    - [ ] Enforce certain leave types on directory tree.
    	- [ ] It is not allowed for a subdirectory just end, without having a declared type.
    	- [ ] And if the type of a subdirectory does not match its declared type, this needs to result in a TODO (depending on a policy).
    - [ ] Allow to create listings of lists of subtrees.
    	- [ ] eg run a command that lists all TODOs in a subdirectory.
    - [ ] Show statistics.
    	- [ ] For example how many percent of directory leaves are updated
    - [ ] Mainly have everything under one directory (the DGE root), but allow to link in other local directories.
    	- [ ] eg by following links between files as created by the LN command.
    - [ ] Other optional features (per directory?) ...
      - [ ] Automatically run shellcheck on shell files.
      - [ ] Have a daemon or cron job that checks for integrity on remote servers (?).
  - [ ] Tool: recursive file tree manager.
    - [ ] creates `.DGE` subdirectory .
    - [ ] each tree is fully managed until it ends in leaves.
    - [ ] managed means a rule has to be declared for each directory or file (path).
    - [ ] each path has to match a declared pattern.
    - [ ] leaves can generated TODOs.
    - [ ] leaves can be other DGE projects.
    - [ ] have commands to put file under and out of control.
      - [ ] `DGE add <file-name>` -- add file to control.
      - [ ] `DGE remove <file-name>` -- remove file from control.
  - [ ] Tool: tagger.
    - [ ] watch over a directory of similar media files (eg my film/picture/music/meme collection).
    - [ ] tag individual files
    - [ ] run checks on the tags: tag groups cover all files.
    - [ ] run checks on the tags: tag groups do not tag a file twice.
    - [ ] group tags together to groups.
    - [ ] define and check constraints on groups.
      - [ ] eg mutually exclusive groups
  - [ ] Tool: directory graph enforcer (DGE)
    - [ ] Should become a tool to easily check if a directory structure conforms to certain rules specified in a rule file.
      - [ ] These rules can be expressed in config files, similar to `.gitignore`, throughout a project tree.
    - [ ] A DGE project could have an entire home direcotry, containing several actual projects.
      - [ ] Where git is used to organize projects, the DGE is used to manage the overview over several projects and glues them together.
    - [ ] It should also be able to warn if hidden side projects exist in unexpected corners of the directory graph that are not checked-in into git.
      - [ ] Like this DGE can prevent data loss.
    - [ ] Another use case is to enforce an easy-to-understand directory graph, with flat hierachy and evenly formed branches.
      - [ ] prevent directory-depth to go over a certain depth level.
    - [ ] It should have several features.
      - [ ] read rules for project layout and alert if there are diviations.
      - [ ] enforce evenly formed folders.
      - [ ] enforce upper bound to leaf depth (eg no folders deeper than 3).
      - [ ] enforce leafs are all projects (fully-functional git projects or otherwise marked projects).
  - [ ] MAYBE: merge-in GITWALKER project.
  - [ ] MAYBE: merge-in HUB project.
  - [ ] MAYBE: merge-in BDIRBS project.
  - [ ] feature:
    - [ ] open file, read line by line, each line is a path and a rule set for that path.
    - [ ] path can contain home: replace home directory.
    - [ ] search find files in path.
    - [ ] see tree behind that path matches rule set.
    - [ ] print git status from within go, list uncommited changes.
    - [ ] print git status from within go, list unpushed changes.
  - [ ] feature: allow watchcat directories to be git repos.
  - [ ] feature: give watchcat the ability to check for the existence of certain git remote addresses.
  - [ ] feature: restore computer setup from a list of directories and git remotes.
  - [ ] implement checking the rules and alert (positive and negative).
  - [ ] implement linting of rules.
  - [ ] implement profiles or Watchconf files (ie user-specified profiles, in addition to pre-defined profiles: eg media store for videos and music, categorized media store, projects of restricted complexity).
  - [ ] when everything else is coded: write some unit tests (eg by writing everything in python).
  - [ ] when everything else is coded: implement algorithm to check for conflicting rules.
  - [ ] decide what to do if a git repo includes a watchconf file.
  - [ ] maybe use other linter or C module to identify size of direcotries faster.
  - [ ] implement the following watchconf codes.
  - [ ] make tests for all watchconf codes (and permutations).

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

## Mini Documentation

Watchcat defines layout rules for directory trees.
Each directory tree is a target.
All targets are listed in a watchcat configuration file.
Rules can be defined in a watchcat configuration file or in a separate Watchconf
file within the target directory.

  - Watchcat config file: 
  - Watchconf file: 

```text
OPS    | PATH               | COMMENT or                   |
PREFIX |                    | COMMENT with OPS SUFFIX      |
-------+--------------------+------------------------------+--------
ezN{}   .                   : Comment.                     | RULE
        ./folderD/          :[t(txt,srt,sub,mp4)] Comment. | RULE
                                                           | ...
```

Planned operand codes (op codes, or ops).

  - [ ] `a(dyfz;min/max)` -- requires things (dir/file/symlink) to have names of minimum or maximum length or both.
  - [ ] `d` -- requires there be no directories in this directory.
  - [ ] `e` -- requires directory to be empty.
  - [ ] `f` -- requires there be no files in this directory.
  - [ ] `g(xur)` -- requires the directory to have a Git repo (distinguish: repo exists, uncommit changes exist, remote repo exists).
  - [ ] `h(dfyz)` -- requires there to be no hidden things (dir/file/symlink) in this directory.
  - [ ] `i(min/max)` -- requires directory to be of minimum or maximum size on disk or both.
  - [ ] `l(dfyz)` -- requires names to only use lowercase letters (distinguish: files, directories, symlinks).
  - [ ] `m(dfyz;min/max)` -- requires there to be things (dir/file/symlink) only from min to max tree depth.
  - [ ] `n(f;A-z,...)` -- requires thing names not to have the specified characters. See more info below.
  - [ ] `r(dfyz;regex)` -- requires thing names (dir/file/symlink) to be like regex.
  - [ ] `s(min/max)` -- requires files to be of minimum or maximum size or both.
  - [ ] `t(type,...)` -- requires to be only to be of the specified file types.
  - [ ] `u(dfyz)` -- requires names to only use uppercase letters (distinguish: files, directories, symlinks).
  - [ ] `w` = requires directory to have a Watchconf.
  - [ ] `x` -- requires for the directory to exist.
  - [ ] `y(hs)` -- requires there be no symlinks in this directory (distinguish: hard and soft).

Other:

  - [ ] `!` -- forbid everything except allowed things (permissions through uppercase letter specifications).
  - [ ] `*` -- inherit all operands from parent.
  - [ ] `.` -- require directory to exist.
  - [ ] `?` -- requires the directory tree under this directory not have any rules specifications, not in config, nor Watchconf files.
  - [ ] `b(dyfz)` -- requires there to be no duplicated things (dir/file/symlink).
  - [ ] `p(dyfz)` -- requires there to be no duplicate thing names (dir/file/symlink).

Combinations and shortcuts:

  - [ ] `z` == `dyf`, requires the directory to be empty (of files, subdirectories and symlinks).
  - [ ] `{from-to;categories}`, special syntax which only allows certain categories at certain depth of the directory tree.
    - [ ] for example: `{0-1d}` -- in this directory and below only directories, i.e. no files.
    - [ ] for example: `{0-1d}{2gw}` -- in the directory level below only git repos or further Watchconf files.
    - [ ] for example: `{0-1d}{2f}`  -- in the directory level below only files.
    - [ ] for example: `{0-1d}{2u}` -- in the directory level below git repos with uncommitted changes, but not without remotes.

Notes:

  - [ ] `N` and `n`:
    - [ ] distinguish: files, directories, symlinks.
    - [ ] specify groups: `A-Z`, `a-z`, `A-z`, `0-9`, etc.
    - [ ] specify special groups: `<whitespace>`, `<alphanum>`, etc.
    - [ ] specify single characters: `abcABC!?:;,.-_=`, etc.
    - [ ] specify special characters: `<space>`, `<tab>`, etc.
  - [ ] `T` and `t`:
    - [ ] specify pre-defined single types: `jpg` (for jpeg and jpg), `png`, etc.
    - [ ] specify pre-defined type groups: `<img>`, `<vid>`, `<aud>`, `<txt>`, `<bin>`, etc.
    - [ ] tested with `file` command on linux (not by file ending, which is done by the `r()` operand.

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

