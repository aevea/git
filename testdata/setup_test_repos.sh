#!/bin/bash

cd testdata

echo 'testdata/commits_on_branch_test/ directory does not exist at the root; creating...'
rm -rf commits_on_branch
git clone commits_on_branch.bundle
echo 'done'

echo 'testdata/git_tags/ directory does not exist at the root; creating...'
rm -rf git_tags
git clone git_tags.bundle
echo 'done'

echo 'testdata/annotated_git_tags_mix/ directory does not exist at the root; creating...'
rm -rf annotated_git_tags_mix
git clone annotated_git_tags_mix.bundle
echo 'done'

echo 'testdata/detached_head/ directory does not exist at the root; creating...'
rm -rf detached_head
git clone detached_head.bundle
echo 'done'
