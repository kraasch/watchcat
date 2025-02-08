
# watchcat

## to-dos

### random thoughts and ideas

Main purpose.

  - keep overview over files by white-listing everything
    - [ ] what state is a directory or project in.
    - [ ] have everything known white-listed, such that new things appear as new.

A collection of constraints to the file system, with the following features:

  - Throws errors (or alerts, eg system popups, messages, etc) when something breaks the constraints.
  - Direct remote calendar support:
  	- [ ] Eg: Automatically pop-up the errors as to-do items in CalDav, WebDav, iCal, etc (?).
  	- [ ] Let user customize which event types become to-do lists.
  - Constraints can be added, edited and removed.
  - Solutions for broken constraints can be run automatically.
  - Recursively checks for `.git` directories and uncommited changes (or other unsafe git repo states).
  - Check if directories have backup.
  	- [ ] Git repos can be compared with their remote.
  	- [ ] Other directories can use **rsync**.
  - Distinguish remote and local subdirectories. 
  	- [ ] Have different policies for both types or repos.
  	- [ ] For example remotes can have the policy of resulting strict TODOs or non-strict TODOs, independent of the policy of local repos.
  - Enforce certain leave types on directory tree.
  	- [ ] It is not allowed for a subdirectory just end, without having a declared type.
  	- [ ] And if the type of a subdirectory does not match its declared type, this needs to result in a TODO (depending on a policy).
  - Allow to create listings of lists of subtrees.
  	- [ ] eg run a command that lists all TODOs in a subdirectory.
  - Show statistics.
  	- [ ] For example how many percent of directory leaves are updated
  - Mainly have everything under one directory (the DGE root), but allow to link in other local directories.
  	- [ ] eg by following links between files as created by the LN command.
  - Other optional features (per directory?) ...
    - [ ] Automatically run shellcheck on shell files.
    - [ ] Have a daemon or cron job that checks for integrity on remote servers (?).

Tool: directory graph enforcer (DGE)

  - Should become a tool to easily check if a directory structure conforms to certain rules specified in a rule file.
    - [ ] These rules can be expressed in config files, similar to `.gitignore`, throughout a project tree.
  - A DGE project could have an entire home direcotry, containing several actual projects.
    - [ ] Where git is used to organize projects, the DGE is used to manage the overview over several projects and glues them together.
  - It should also be able to warn if hidden side projects exist in unexpected corners of the directory graph that are not checked-in into git.
    - [ ] Like this DGE can prevent data loss.
  - Another use case is to enforce an easy-to-understand directory graph, with flat hierachy and evenly formed branches.
    - [ ] prevent directory-depth to go over a certain depth level.
  - It should have several features.
    - [ ] read rules for project layout and alert if there are diviations.
    - [ ] enforce evenly formed folders.
    - [ ] enforce upper bound to leaf depth (eg no folders deeper than 3).
    - [ ] enforce leafs are all projects (fully-functional git projects or otherwise marked projects).

Example of a bad directory graph.

```text
  ______/______
./ ./  / \    \
      /\  \.  /\.
     /  \   ./
   ./\.  \.
```

Example of a good directory graph.

```text
         /\
  ______/  \____________
./ ./ ./    \. \. \. \. \.
   .
```

### tasks

First.

 - [ ] write in GO language.
   + [ ] checkout "github.com/go-git/go-git/v5"
   + [ ] open file, read line by line, each line is a path and a rule set for that path.
   + [ ] path can contain home: replace home directory.
   + [ ] search find files in path.
   + [ ] see tree behind that path matches rule set.
   + [ ] print git status from within go, list uncommited changes.
   + [ ] print git status from within go, list unpushed changes.
 - [ ] merge in GITWALKER project.

 - [ ] MAYBE: merge in HUB project.
 - [ ] MAYBE: merge in BDIRBS project.

Later.

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
 - [ ] # WATCHCONF CODES (default linting: alert about broken rules).
   - [ ] i             = ignore directory (ensure no rules are specified for directory subtree).
   - [ ] o             = ensure directory does not contain any files (linting: delete fast, shred -zu , view+select+delete).
   - [ ] O             = ensure directory does not contain any directories (linting: delete fast, shred -zu , view+select+delete).
   - [ ] z             = ensure directory or file is empty (if it exists) (linting: delete fast, shred -zu , view+select+delete).
   - [ ] e             = ensure directory or file exists (linting: create).
   - [ ] s             = ensure all files have the same directory depth.
   - [ ] T(e1[, en])   = ensure only the listed file types exist (as output by the linux command 'file').
   - [ ] t(e1[, en])   = ensure only the listed file endings exist (eg. .pdf, .txt, .md, .srt, .sub, .mp4, .mp3, etc).
   - [ ] d(N)          = ensure minimum directory depth of length N.
   - [ ] D(N)          = ensure maximum directory depth of length N.
   - [ ] f(N)          = ensure minimum file depth of length N.
   - [ ] F(N)          = ensure maximum file depth of length N.
   - [ ] w             = allow existence of Watchconf sub trees.
   - [ ] W             = allow existence of Watchconf sub trees (recursively include them into the report).
   - [ ] h             = allow existence of hidden directories and files.
   - [ ] g             = allow existence of git direcotries.
   - [ ] G             = treat git directories as final nodes.
   - [ ] p[a-Z...]     = allow file names to only contain the specified characters.
   - [ ] P[a-Z...]     = allow directory names to only contain the specified characters.
   - [ ] n[a-Z...]     = like n, but disallows specified characters.
   - [ ] N[a-Z...]     = like N, but disallows specified characters.
   - [ ] x(N)          = ensure file size max.
   - [ ] X(N)          = ensure file size min.
   - [ ] y(N)          = ensure directory size max.
   - [ ] Y(N)          = ensure directory size min.
   - [ ] a             = forbid duplicated files.
   - [ ] A             = forbid duplicated directories.
   - [ ] b             = forbid duplicated file names.
   - [ ] B             = forbid duplicated directory names.

### watchconf file example

```text
.                   ezN[ ]
./folderA/          z
./folderA/subA/     ez
./folderA/subB/     ez
./folderB/          zs
./folderB/fileA.txt e
./folderB/fileB.txt z
./folderC/          ez
./folderD/          t(txt,srt,sub,mp4)
./folderE/          t(mp3)
```
