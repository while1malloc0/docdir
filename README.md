# docdir

A small tool for summarizing directory structures

## Introduction

`docdir` prints a Unix `tree`-esque representation of a directory structure, but with each folder annotated with a description.

## Quickstart

Running `docdir` on an empty directory with no summary files will simply print the name of the directory.

```shell
$ cd examples/simple
$ docdir no-summary
no-summary
```

Any subdirectories will be printed a-la `tree`.

```shell
$ cd examples/nested
$ docdir no-summary
no-summary
├── has-no-sub-dirs
└── has-sub-dirs
    └── a-sub-dir
```
