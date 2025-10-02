#! /bin/bash

echo "Creation script for ./cli01/ at '$(pwd)'."

# clean.
rm -rf ./cli01

# create directories and files.
mkdir ./cli01
mkdir ./cli01/cfg
mkdir ./cli01/targetDir0
mkdir ./cli01/targetDir0/firefox
mkdir ./cli01/targetDir0/downloads
mkdir ./cli01/targetDir0/downloads/done
mkdir ./cli01/targetDir0/downloads/incomplete
echo 'ABC content' >>./cli01/targetDir0/abc.txt
echo 'ABC content' >>./cli01/targetDir0/firefox/abc.txt
echo 'ABC content' >>./cli01/targetDir0/downloads/abc.txt
echo 'ABC content' >>./cli01/targetDir0/downloads/done/abc.txt
echo 'ABC content' >>./cli01/targetDir0/downloads/incomplete/abc.txt

# copy directories.
cp -r ./cli01/targetDir0 ./cli01/targetDir1

# insert files.
cp ./cli01_assets/config.toml cli01/cfg/config.toml
cp ./cli01_assets/watchcat.txt cli01/targetDir1/.Watchcat
