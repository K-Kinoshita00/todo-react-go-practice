# todo-react-go-practice

個人開発用の ToDo Web アプリです。機能は最小の CRUD に抑え、認証・認可・ローカル環境・CI/CD・安価なクラウド公開まで、本番 SaaS で使われる型に沿って学びます。

## 方針

- 実装は手書きのみ
- AI はレビューのみ使う（「レビュー」と頼む）
- 後から層を足さない。空でもディレクトリと責務は最初から置く
- 1 日約 30 分。飛ばさず、終わらなければ翌日に持ち越す
- Cursor は `.cursor/rules/` と `.cursor/skills/review-day/` がこの方針を強制する

進め方・日次計画・完了条件は [docs/LEARNING_ROADMAP.md](docs/LEARNING_ROADMAP.md) だけを読めば再開できます。

## 構成（目標）

モノレポです。

- `api/` — Go。domain / application / infra / interface / registry
- `web/` — Vite + React + MUI
- `openapi/` — API 契約の単一ソース
- `database/migrations/` — SQL マイグレーション
- ルートの `compose.yaml` と `Makefile` — ローカル一式

## 全体の完了条件

- ログインしたユーザーが自分の ToDo を CRUD できる
- ローカルは Compose 一発で起動する
- PR で lint / test が回る
- 無料枠寄りの AWS に載せられる

## 非公開メモ

社内の参考実装との対応表はリポジトリに含めません。手元に `docs/private/` がある場合だけ参照してください。
