<!-- Code generated by docsdata/README.tpl.md. DO NOT EDIT. -->
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

If a `DESCRIPTION` file is found, the first sentence of that file will be printed.

```shell
$ cd examples/nested
$ docdir with-description
with-description     # this example has descriptions in it
├── has-no-sub-dirs  # this directory has no subdirectories
└── has-sub-dirs     # this directory has subdirectories
    └── a-sub-dir    # this directory is a subdirectory
```
