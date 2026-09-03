---
name: start-day
description: >-
  Starts the current Day from docs/LEARNING_ROADMAP.md 現在地.
  Shows 目標 / やること / 完了条件 and how to proceed. Does not write app code.
  Use when the user invokes this command to begin today's work.
  Trigger: "start-day", "今日を始める", "Day を始める"
user-invocable: true
---

# Start Day — 記載の日を始める

Use this command when the user wants to begin the Day recorded as `今日やる日` in `docs/LEARNING_ROADMAP.md`. Follow `.cursor/rules/learning-mode.mdc`.

---

## Iron Law(s)

> **1. アプリの完成実装（ファイル全体・完成関数・適用可能な差分パッチ）は出さない。ヒントの一部コード（シグネチャ・骨格・断片）は可。**
> **2. 「現在地」の当日だけを扱う。先の Day を実装・提案しない。Day 本文・現在地・進捗は書き換えない。**

---

## Workflow / Steps

1. `docs/LEARNING_ROADMAP.md` の「現在地」を読む。`今日やる日` を N とする。N が「なし」なら全体完了として止める。
2. Day N セクションだけ読む（目標 / やること / 完了条件）。振り返り質問は出さない（`/next-day` のあと）。先の Day は開かない。
3. 直前 Day の「翌日の入口」がある場合だけ、そこに書かれたパスを見る。
4. `docs/private/` がある場合だけ、当日の対応表を読む。固有名詞・URL・アカウント情報を出力に出さない。
5. 「やること」と「翌日の入口」に出たパスの有無だけ確認する。中身の完成コードは書かない。
6. 層ルールが当日に関係するなら `api-layers` / `web-features` を読む。破る書き方は勧めない。
7. Output Format で当日を始める。実装はユーザーが手書きする。求められても完成実装は出さない。

---

## Decision Table

| Situation | Choice | Reason |
| --- | --- | --- |
| `今日やる日` が `なし` | ファイルを変えず、全体完了と伝える | 始める日がない |
| 「書いて」「実装して」 | 拒否し、やることと方針だけ示す | 完成実装は手書き |
| 「ヒント」「なぜ」 | 方針と理由。シグネチャ・骨格・断片は可 | 学習を奪わない |
| 先の Day の話 | 当日に戻す | 飛ばさない |
| `docs/private/` あり | 当日対応表だけ読む。固有名詞は出さない | 公開面に残さない |

---

## Output Format

日本語。次の形だけ。完成実装・適用可能な差分パッチは載せない。

```markdown
## 今日
Day N — <タイトル>

## 目標
<当日の目標を短く>

## やること
- <ロードマップの項目。手書きする場所が分かるようにパスを残す>

## 完了条件
- <当日の完了条件>

## 方針
<層の置き場所と依存の向き。一部コードならシグネチャ・骨格・断片まで。完成関数は出さない>

## 次
手書きしたら `/review-day`。Pass なら `/next-day`。
```

---

## Examples

- Good: 「`Todo` は `api/pkg/domain/entity/`。usecase は repository の interface に依存する」
- Good: ヒントとして `type Todo struct` のフィールド名や usecase の関数シグネチャだけ出す
- Bad: entity / usecase の完成ファイルや適用可能な差分パッチを貼る
- Bad: Day 15 の handler 実装を Day 14 の開始で要求する
