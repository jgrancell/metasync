# Metasync

Metasync provides a convenient way to quickly sync commonly used and repeated files between multiple repositories, and keep changes in sync.

[![metasync](https://img.shields.io/travis/com/jgrancell/metasync?style=for-the-badge&logo=travis)](https://travis-ci.com/github/jgrancell/metasync)
[![coverage](https://img.shields.io/codecov/c/github/jgrancell/metasync?color=65187a&style=for-the-badge&token=p8NQJsRPDX)](https://codecov.io/gh/jgrancell/metasync/)

[![releases](https://img.shields.io/github/v/release/jgrancell/metasync?style=for-the-badge)](https://github.com/jgrancell/metasync/releases)
[![GitHub license](https://img.shields.io/github/license/jgrancell/metasync?color=333333&style=for-the-badge)](https://github.com/jgrancell/metasync/blob/main/LICENSE)

## Objectives

It gets extremely time consuming maintaining significant numbers of repositories, especially when you use similar CI/CD tools throughout and have a number of meta or helper files in each repository.

Rather than having to go into each repository individually to make file changes, metasync provides a straightforward way to sync your metafiles with a master template repository.

Metasync allows you to keep your metafiles in sync, and even version them against specific template repository tags, branches, or individual refs.

## Examples of Use

| Command | Explanation | Minimum Version |
| :-----: | :---------: | :-------------: |
| `metasync sync` | Syncs the local repository with a remote source defined in a local `.metasync.yml` files | ------- |
| `metasync sync -dryrun` | Shows required changes between the local repo and source templates, without making changes. | ------- |
| `metasync sync -diff` | Shows inline diffs of changes between the remote source template and the local files. | ------- |