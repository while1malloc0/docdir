# docdir

A small tool for summarizing directory structures

## Introduction

`docdir` prints a Unix `tree`-esque representation of a directory structure, but with each folder annotated with a description.

## Quickstart

Running `docdir` on an empty directory with no summary files will simply print the name of the directory.
[embedmd]:# (../e2e/testdata/simple.ct shell)

Any subdirectories will be printed a-la `tree`.
[embedmd]:# (../e2e/testdata/nested.ct shell)