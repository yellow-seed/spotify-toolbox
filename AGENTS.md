# AI Agents Documentation

このドキュメントは、AIエージェントがこのプロジェクトを理解し、適切に支援するための情報を提供します。

## プロジェクト概要

Spotify Toolbox は、Spotifyのポッドキャストを一括管理するためのCLIツールです。Go言語とCobraフレームワークを使用して構築されています。

### 主な機能

- ライブラリ内のポッドキャスト一覧表示
- 条件指定によるポッドキャストの一括削除
- Dry-runモードによる安全な操作確認

## 技術スタック

- **言語**: Go 1.21以上
- **CLIフレームワーク**: [Cobra](https://github.com/spf13/cobra) v1.10.2
- **設定管理**: [Viper](https://github.com/spf13/viper) v1.21.0
- **ビルドツール**: Make

## ディレクトリ構造

```text
spotify-toolbox/
├── cmd/
│   └── spotify-toolbox/    # アプリケーションのエントリーポイント
│       └── main.go         # main関数
├── internal/
│   ├── cmd/                # Cobraコマンド定義
│   │   ├── root.go        # ルートコマンドと設定初期化
│   │   ├── list.go        # リストコマンド
│   │   └── delete.go      # 削除コマンド
│   ├── spotify/            # Spotify APIクライアント（今後実装予定）
│   └── config/             # 設定管理（今後実装予定）
├── pkg/
│   └── logger/             # ロギングユーティリティ（今後実装予定）
├── bin/                    # ビルド成果物（.gitignore対象）
├── .env.example            # 環境変数のサンプル
├── go.mod                  # Goモジュール定義
├── go.sum                  # 依存関係のチェックサム
├── Makefile                # ビルドタスク定義
└── README.md               # プロジェクトドキュメント
```

### ディレクトリの役割

- `cmd/`: アプリケーションのエントリーポイント。mainパッケージを含む
- `internal/`: プロジェクト固有の非公開コード。外部からインポート不可
- `pkg/`: 外部からもインポート可能な汎用的なコード（今後必要に応じて実装）

## 開発環境のセットアップ

### 必要な環境

- Go 1.21以上
- Make（オプションだが推奨）
- Spotify Developer アカウント

### セットアップ手順

1. リポジトリのクローン

   ```bash
   git clone https://github.com/yellow-seed/spotify-toolbox.git
   cd spotify-toolbox
   ```

1. 依存関係のインストール

   ```bash
   make install
   # または
   go mod download
   ```

1. 環境変数の設定

   ```bash
   cp .env.example .env
   # .envファイルを編集してSpotify API認証情報を設定
   ```

1. ビルド

   ```bash
   make build
   # または
   go build -o bin/spotify-toolbox ./cmd/spotify-toolbox
   ```

1. 実行

   ```bash
   ./bin/spotify-toolbox --help
   ```

## コーディング規約

### Go標準規約

- `gofmt`による自動フォーマット（`make fmt`で実行）
- `go vet`による静的解析（`make vet`で実行）
- [Effective Go](https://go.dev/doc/effective_go)のベストプラクティスに従う

### プロジェクト固有規約

- パッケージ名は小文字の単一単語
- エクスポートされる関数・型は明確なドキュメントコメントを記載
- エラーハンドリングは適切に行い、エラーを無視しない
- コンテキストを使用する関数は第一引数に`ctx context.Context`を受け取る

## コミットメッセージ規約

このプロジェクトでは、[Conventional Commits](https://www.conventionalcommits.org/)形式を使用しています。

### 形式

```
<type>: <subject>

<body>

<footer>
```

### Type（必須）

- `feat`: 新機能の追加
- `fix`: バグ修正
- `docs`: ドキュメントのみの変更
- `style`: コードの動作に影響しない変更（フォーマット、セミコロンなど）
- `refactor`: バグ修正や機能追加ではないコード変更
- `perf`: パフォーマンス改善
- `test`: テストの追加や修正
- `chore`: ビルドプロセスやツールの変更

### Subject（必須）

- 50文字以内
- 命令形で記述（例: "Add" ではなく "Adds" でもなく "Add"）
- 最初の文字は大文字にしない
- 末尾にピリオドを付けない

### Body（オプション）

- 変更の理由や方法を説明
- 72文字で折り返す
- 命令形で記述

### Footer（オプション）

- 破壊的変更がある場合は `BREAKING CHANGE:` を記載
- Issue番号を参照する場合は `Closes #123` など

### 例

```
feat: add user authentication

Implement JWT-based authentication system with login and logout endpoints.

Closes #123
```

```
fix: resolve memory leak in data processing

The issue was caused by not properly releasing resources after processing.
This fix ensures all resources are cleaned up correctly.
```

## テスト戦略

### テストの実行

```bash
# 全テスト実行
make test

# カバレッジ付きテスト
make test-coverage
```

### テスト方針

- ユニットテスト: `*_test.go`ファイルに記述
- テーブル駆動テストを推奨
- モックが必要な場合はインターフェースを定義
- カバレッジ目標: 80%以上（今後設定予定）

## デプロイメント

### ローカルインストール

```bash
make build
# バイナリは bin/spotify-toolbox に生成される
```

### リリース手順（今後実装予定）

1. バージョンタグの作成
2. GitHub Actionsによる自動ビルド
3. マルチプラットフォームバイナリの生成
4. GitHub Releasesへの公開

## その他の重要な情報

### Cobraフレームワークの構造

このプロジェクトではCobraを使用したCLI実装を採用しています。

- `internal/cmd/root.go`: ルートコマンドと共通設定
- `internal/cmd/*. go`: 各サブコマンドの実装
- 設定ファイルとフラグの両方をサポート（Viper使用）

### 設定の優先順位

1. コマンドラインフラグ
2. 環境変数（`SPOTIFY_`プレフィックス）
3. 設定ファイル（`~/.spotify-toolbox.yaml` または `--config`で指定）

### 今後の実装予定

- Spotify Web API統合
- OAuth 2.0認証フロー
- ポッドキャスト一覧取得・削除機能の実装
- ユニットテストの追加
- CI/CDパイプラインの構築

