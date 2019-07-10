#!/bin/sh

git checkout --orphan tmp
git add -A
git commit -am "update"
git branch -D master
git branch -m master
git push -f origin master
