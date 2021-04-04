# Algorithms, explained and illustrated

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
$ cd 2_sort/2_bubble_sort
$ go test
input  = []
output = []
input  = [5 9 3 1 2 8 4 7 6]
  list(1) = [5 9 3 1 2 8 4 7 6]
  list(2) = [1 5 9 3 2 4 8 6 7]
  list(3) = [1 2 5 9 3 4 6 8 7]
  list(4) = [1 2 3 5 9 4 6 7 8]
  list(5) = [1 2 3 4 5 9 6 7 8]
  list(6) = [1 2 3 4 5 6 9 7 8]
  list(7) = [1 2 3 4 5 6 7 9 8]
  list(8) = [1 2 3 4 5 6 7 8 9]
output = [1 2 3 4 5 6 7 8 9]
PASS
ok  	example.com/algorithms/2_sort/2_bubble_sort	0.227s
```

## Ref.

* Go の testing パッケージの基本を理解する - Qiita

    https://qiita.com/taisa831/items/85fea8d970bcadd796b9
