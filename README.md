<div align="center">

![Last commit](https://img.shields.io/github/last-commit/Comamoca/igonore?style=flat-square)
![Repository Stars](https://img.shields.io/github/stars/Comamoca/igonore?style=flat-square)
![Issues](https://img.shields.io/github/issues/Comamoca/igonore?style=flat-square)
![Open Issues](https://img.shields.io/github/issues-raw/Comamoca/igonore?style=flat-square)
![Bug Issues](https://img.shields.io/github/issues/Comamoca/igonore/bug?style=flat-square)

# 📄 igonore

cli tool to generate gitignore files written in Go

</div>

<table>
  <thead>
    <tr>
      <th style="text-align:center">🍔English</th>
      <th style="text-align:center"><a href="README.ja.md">🍡日本語</a></th>
    </tr>
  </thead>
</table>

<div align="center">

</div>

## 🚀 How to use

```sh
# see help
igonore help

  igonore - .gitignore generator written in Go

  Usage: igonore
         Generate .gitignore interactively.

         igonore langage langage2...
         Generate a .gitignore file for the specified language.

         Ex.) igonore go node

# Generate interactively
igonore

# specify language
igonore go node
```

## ⬇️  Install

`go install github.com/Comamoca/igonore@latest`

## ⛏️   Development

```sh
git clone https://github.com/Comamoca/igonore
cd igonore
go build
./igonore
```

## 📝 Todo

Notthing... :zzz:

## 📜 License

MIT

### 🧩 Modules

|Repositories                       |Licenses                               |
|-----------------------------------|---------------------------------------|
|github.com/nsf/termbox-go          |MIT License                            | 
|github.com/lucasb-eyer/go-colorful |MIT License                            | 
|github.com/cuonglm/gogi            |BSD 3-Clause "New" or "Revised" License|
|github.com/gdamore/tcell           |Apache License 2.0                     |
|golang.org/x/term                  |BSD 3-Clause "New" or "Revised" License|
|golang.org/x/text                  |BSD 3-Clause "New" or "Revised" License|
|github.com/manifoldco/promptui     |BSD 3-Clause "New" or "Revised" License|
|github.com/mattn/go-runewidth      |MIT License                            |   
|github.com/gdamore/encoding        |Apache License 2.0                     | 
|github.com/pkg/errors              |BSD 2-Clause "Simplified" License      | 
|github.com/rivo/uniseg             |MIT License                            | 
|github.com/chzyer/readline         |MIT License                            | 
|golang.org/x/sys                   |BSD 3-Clause "New" or "Revised" License|
|github.com/ktr0731/go-fuzzyfinder  |MIT License                            | 

## 👏 Affected projects

[gitignore.io](https://www.toptal.com/developers/gitignore/)

## 💕 Special Thanks

- mh-cbon
[gigo](https://github.com/mh-cbon/gigo)
