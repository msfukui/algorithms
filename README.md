# Algorithms, explained and illustrated

[![Go build & test](https://github.com/msfukui/algorithms/actions/workflows/go.yml/badge.svg)](https://github.com/msfukui/algorithms/actions/workflows/go.yml)

「アルゴリズム図鑑 絵で見てわかる26のアルゴリズム」(ISBN978-4-7981-4977-6) の読書メモのリポジトリです。

## Setup

for macOS:

```
$ sudo port install go
...
$ go mod init example.com/algorithms
```

## Procedure

* 対象のディレクトリを作って移動します

* 入力と出力を決めて雛形のテストコードを書きます

* テストが red であることを確認します

* 本体の雛形を書いてテストを green にします

* テストをちゃんと書いて red であることを確認します

* 本体をちゃんと書いて green にします

* 以下、コードに納得できるまで繰り返します

* アルゴリズムの動きの途中経過がわかるように、適度に本体内で fmt.Printf します

## Test

example:

```
$ go test ./...
ok  	example.com/algorithms/2_sort/2_bubble_sort	0.314s
ok  	example.com/algorithms/2_sort/3_selection_sort	0.523s
ok  	example.com/algorithms/2_sort/4_insertion_sort	0.927s
ok  	example.com/algorithms/2_sort/5_heap_sort	1.140s
```

## Ref.

* golangのテストはじめ - Qiita

    https://qiita.com/tmzkysk/items/8bb37795ac223664d682

* Go の testing パッケージの基本を理解する - Qiita

    https://qiita.com/taisa831/items/85fea8d970bcadd796b9

* gosec - Golang Security Checker

    https://github.com/securego/gosec

* Goの静的解析ツールをgolintからstaticcheckに移行した話

    https://blog.cybozu.io/entry/2021/02/26/081013

* Reviewdog を飼ってコードレビューや開発を改善しませんか

    http://haya14busa.com/reviewdog/

* セキュアにGoを書くための「ガードレール」を置こう - 安全なGoプロダクト開発に向けた持続可能なアプローチ

    https://flattsecurity.hatenablog.com/entry/go_guardrail
