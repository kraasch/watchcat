#! /bin/python3

#> synopsis: TODO: write down.
#> type:     TODO: write down.
#> comment:  TODO: write down.

import os
import pathlib
import argparse

watchcat_modes = [
    'info',  # TODO: options: --level --> show the info with certain details (repo, root_folders, rule, report)
    'check', # TODO: options: --force --> do not preview files, just create specified state.
]

watchconf_name = 'Watchconf'
targets_file = '~/.config/watchcat/targets.txt'
targets_file = pathlib.Path('~/.config/watchcat/targets.txt').expanduser()
watchdir_reports = []

def read_watchconf(watchconf_path, target_dir):

    report = []

    if watchconf_path.exists():

        rules = watchconf_path.read_text().splitlines()
        parsed = []
        allowed_paths = []

        # parse configuration file.
        for rule in rules:
            path, codes = rule.split('|')
            path = path.strip().removesuffix("'").removeprefix("'")
            path = pathlib.Path(path) # make string a pathlib Path object.
            codes = [*codes] # split single code string into array of characters.
            parsed.append([path, codes])
            allowed_paths.append(str(path))

        # check directory.
        for rule_path, code in parsed:
            full_rule_path = target_dir / rule_path
            warnings = []
            errors = []

            # check if directory exists, otherwise create directory.
            if 'e' in codes:
                if not rule_path == '.': # do not create a directory named dot for the repository root.
                    if not full_rule_path.exists(): # create directory if it does not exist.
                        full_rule_path.mkdir(parents=True, exist_ok=True)
                        warnings.append(f'Warning: was created')

            # check if directory is zero (empty), otherwise alert.
            if 'z' in codes:
                for file in list(full_rule_path.iterdir()): # list all files in directory.
                    file = file.relative_to(target_dir)
                    if not str(file) == watchconf_name: # ignore configuration file itself.
                        has_illegal_file = not str(file) in allowed_paths
                        if has_illegal_file:
                            errors.append(f'Error: illegal file "{file}".')

            # log reports.
            report.append([full_rule_path, errors, warnings])
    else:
        print(f'Error: no watchconf file ("{watchconf_path}") for watchcat target directory "{target_dir}"')
    watchdir_reports.append([target_dir, report])

def open_watchcat_directories():
    for target_dir in targets_file.read_text().splitlines():
        target_dir = pathlib.Path(target_dir).expanduser()
        watchconf_path = target_dir / watchconf_name
        read_watchconf(watchconf_path, target_dir)

def print_reports(watchdir_reports, level=3):
    NL = os.linesep
    result = ''
    for target_dir, reports in watchdir_reports:
        for rule_path, errors, warnings in reports:
            if level == 2: # log per rule.
                result+=f'{rule_path}: {len(errors)}, {len(warnings)}' + NL
            elif level == 3: # log per report
                for warning in warnings:
                    result+=f'{rule_path}: {warning}' + NL
                for error in errors:
                    result+=f'{rule_path}: {error}' + NL
    if result == '':
        result = '✓'
    print(result)

if __name__ == '__main__':
    # parse arguments.
    parser = argparse.ArgumentParser(prog='WATCHCAT', description='Report and enforce rules in directory tree.')
    parser.add_argument('--level', type=int, help='Log level [2-rule or 3-report]')
    args = parser.parse_args()

    open_watchcat_directories()
    print_reports(watchdir_reports, level=args.level)
