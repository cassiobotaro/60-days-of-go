# Installing delve

Follow instructions [here](https://github.com/derekparker/delve/tree/master/Documentation/installation)

# Delve via terminal

## Step 1

start debugger through command `dlv debug`

## Step 2

define breakpoint `b day46/debug.go:6`

## Step 3

let's start debugging `c`

## Step 4

navigate through lines using `s` until line 11

## Step 5

print variable `p greetingArg`

## Step 6

learn more comands typing help

## Step 7

type `q` to exit


# Delve via vim-godebug

https://github.com/jodosha/vim-godebug

## Step 1

open file using neovim `nvim debug.go`

## Step 2

`:GoToggleBreakpoint` to set breakpoint (line 6 is recommended)

## Step 3

`:GoDebug` to start debugging


## Step 4

navigate through lines using `s` until line 11

## Step 5

print variable `p greetingArg`

## Step 6

learn more comands typing help

## Step 7

type `q` to exit
