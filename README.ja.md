<div align="center">

![Last commit](https://img.shields.io/github/last-commit/Comamoca/igonore?style=flat-square)
![Repository Stars](https://img.shields.io/github/stars/Comamoca/igonore?style=flat-square)
![Issues](https://img.shields.io/github/issues/Comamoca/igonore?style=flat-square)
![Open Issues](https://img.shields.io/github/issues-raw/Comamoca/igonore?style=flat-square)
![Bug Issues](https://img.shields.io/github/issues/Comamoca/igonore/bug?style=flat-square)

# 📄 igonore

Goで書かれたgitignoreファイルを生成するcliツール

</div>

<table>
  <thead>
    <tr>
      <th style="text-align:center">🍡日本語</th>
      <th style="text-align:center"><a href="README.md">🍔English</a></th>
    </tr>
  </thead>
</table>

<div align="center">

</div>

## 🚀 使い方

```sh
# ヘルプを見る
igonore help

  igonore - .gitignore generator written in Go

  Usage: igonore
         Generate .gitignore interactively.

         igonore langage langage2...
         Generate a .gitignore file for the specified language.

         Ex.) igonore go node

# インタラクティブに生成する
igonore

# 言語を指定する
igonore go node
```
## ⬇️  Install

`go install github.com/Comamoca/igonore@latest`

## ⛏️   開発

```sh
git clone https://github.com/Comamoca/igonore
cd igonore
go build
./igonore
```
## 📝 Todo

Notthing... :zzz:

## 📜 ライセンス

MIT

### 🧩 Modules

|リポジトリ                         |ライセンス                             | 
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

## 👏 影響を受けたプロジェクト

[gitignore.io](https://www.toptal.com/developers/gitignore/)

## 💕 スペシャルサンクス

- mh-cbon
[gigo](https://github.com/mh-cbon/gigo)
