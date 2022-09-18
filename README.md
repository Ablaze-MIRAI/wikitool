<div align="center">

![Last commit](https://img.shields.io/github/last-commit/Ablaze-MIRAI/wikitool?style=flat-square)
![Repository Stars](https://img.shields.io/github/stars/Ablaze-MIRAI/wikitool?style=flat-square)
![Issues](https://img.shields.io/github/issues/Ablaze-MIRAI/wikitool?style=flat-square)
![Open Issues](https://img.shields.io/github/issues-raw/Ablaze-MIRAI/wikitool?style=flat-square)
![Bug Issues](https://img.shields.io/github/issues/Ablaze-MIRAI/wikitool/bug?style=flat-square)

# Wikitool 🔥

[Ablaze-MIRAI/Wiki](https://github.com/Ablaze-MIRAI/wiki)を管理するためのCLIツール。
Wikiの構造に則っていれば誰でも使用できます。
</div>

## Wikiについて

Wikiとは、Githubでドキュメント管理するためのディレクトリ構造とその実装です。

```sh
.
├── README.md(現在地)
└── カテゴリ
     ├── README.md
     └── 記事.md
```
上のように、カテゴリ名でディレクトリ名を作り、その中に`README.md`と記事を追加します。`README.md`は必須ではありませんが、wikitoolで検索した際にプレビューが表示されるので追加することを推奨します。

## 🚀 使い方

### 新規記事作成

`wikitool new`

### 記事の編集

`wikitool edit`

現在、カテゴリの追加(ディレクトリの新規作成)はwikitoolから出来ない為、新しいカテゴリが必要な場合は自身で作成して下さい。
また、Gitについても連携機能がまだ未実装なためご自身で実行してください。

## ⬇️  Install

`go install https://github.com/Ablaze-MIRAI/wikitool@latest`

## ⛏️   開発

```sh
go mod tidy

go run main.go

go build
```
## 📝 Todo

- [ ] カテゴリーの追加機能
- [ ] 全文検索
- [ ] Git連携

## 📜 ライセンス

MIT

### 🧩 Modules

- promptui
- go-fuzzzyfinder
- cli

## 👏 影響を受けたプロジェクト

GaaTS(Github as a Text Storage)

## 💕 スペシャルサンクス

[Modules](#🧩-modules)
