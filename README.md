# gorillang

ウホゥホウホッッホゥ ウホゥホウホ！！ゥゥ ウホゥホウホ！！オホ ウホ？ゥゥオホ！！ ウホうほホッウウウッホ ウホウォホホウウウホ ウホホホウォ ウホゥホウホォ！ウウ ウホゥホウホウッホッ！ ウホゥホウホッッホッ ウホホッウォウォ？

# 元ネタ様

グランブルーファンタジー ゴリラ語翻訳ジェネレーター『うほうほ〜っ！』

https://xn--bck3aza1a2if6kra4ee0hf.gamewith.jp/article/show/107716

__このコマンドラインツールのゴリラ語は上記元ネタ様と完全互換です。__

# インストール

## MacOS

```
$ brew tap momotaro98/gorillang
$ brew install momotaro98/gorillang/gorillang
```

## Go言語と$GOPATHが通った$PATHを持つ奇特な方

```
$ go get -u github.com/momotaro98/gorillang/cmd
```

# 遊び方

## 日本語 → ゴリラ語

```
$ gorillang encode こんにちは
ウホゥホウホホッゥホ ウホゥホウホオホゥホ ウホゥホウホウォッッ ウホゥホウホウォウウ ウホゥホウホウォウッホ
```

## 日本語 → ゴリラ語

```
$ gorillang decode ウホゥホウホホッゥホ ウホゥホウホオホゥホ ウホゥホウホウォッッ ウホゥホウホウォウウ ウホゥホウホウォウッホ
こんにちは
```

## Tips for MacOS users

コピペした文を利用する場合、`pbpaste`と`xargs`を使えばコマンドを固定できます。

```
$ pbpaste | xargs gorillang encode
$ pbpaste | xargs gorillang decode
```
