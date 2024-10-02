# `sqlc` multi-tenant POC

## How to run

```
make mysql-up
make gen
make run
```

If `make` is not installed, inspect `Makefile` and run the commands listed there.

## Adding `proto` submodule

```
git submodule add https://github.com/sabitm/proto-shared-example ./proto
```

## Updating `proto` submodule

```
git submodule update --remote
```
