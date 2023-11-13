#! /bin/python3

#> synopsis: TODO: write down.
#> type:     TODO: write down.
#> comment:  TODO: write down.

import os
import pathlib

watchconf_name = 'Watchconf'

targets_file = '~/.config/watchcat/targets.txt'
targets_file = os.path.expanduser(targets_file) # expand tilde to home directory.

def read_watchconf(watchconf_path, target_dir):
    if watchconf_path.exists():
        with watchconf_path.open() as f:
            f.readline()
    else:
        print(f'Error: no watchconf file ("{watchconf_path}")for watchcat target directory "{target_dir}"')

def main():
    with open(targets_file) as f:
        for target_dir in f:
            target_dir = os.path.expanduser(target_dir) # expand tilde to home directory.
            target_dir = target_dir.strip()
            watchconf_path = pathlib.Path(target_dir) / watchconf_name
            read_watchconf(watchconf_path, target_dir)

if __name__ == '__main__':
    main()
