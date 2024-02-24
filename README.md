
# watchcat

## to-dos

first.

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

later.

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
