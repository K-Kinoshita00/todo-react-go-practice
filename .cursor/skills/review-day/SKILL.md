---
name: review-day
description: >-
  当日の手書き実装を完了条件でレビューする。実装コードは書かない。
  Load when the user asks for a review of the current Day, completion check, or
  handwritten code.
  Trigger: "レビュー", "見て", "完了条件", "Day", "review".
user-invocable: false
---

# Review Day — 当日レビュー

Use this skill when the user asks to review handwritten work. Follow `.cursor/rules/learning-mode.mdc`.

---

## Iron Law(s)

> **1. 完成コード・差分パッチ・「こう書け」の実装は出さない。**
> **2. 判定は `docs/LEARNING_ROADMAP.md` の当日「完了条件」だけを正とする。**

---

## Workflow / Steps

1. `docs/LEARNING_ROADMAP.md` の「現在地」と当日セクション（目標 / やること / 完了条件）を読む。
2. 当日に触ったファイルだけを見る。先の Day の未着手は指摘しない（予約ディレクトリは除く）。
3. `docs/private/` がある場合は当日の対応表だけ読む。固有名詞をレビュー文面に出さない。
4. 層ルール（`api-layers` / `web-features`）に反していないか見る。
5. 完了条件を 1 項目ずつ Pass / Fail にする。
6. Fail は「どの条件が未達か」と「自分で直す場所（パス）」まで。書き方の全文は出さない。
7. すべて Pass なら「現在地」を次の Day にする更新を提案する。勝手に書き換えない。

---

## Decision Table

| Situation | Choice | Reason |
| --------- | ------ | ------- |
| 完了条件を満たす | Pass。現在地更新を提案 | 日次を進める |
| 動きはするが層が破れている | Fail。層名とファイルを示す | 最初から綺麗に |
| 先の Day の実装がある | 残すなとは言わない。当日範囲外と注記 | 飛ばしていないか確認 |
| 生成物を手編集している | Fail | 契約が単一ソース |
| 公開ファイルに非公開痕跡 | Fail（High） | 公開リポジトリ |
| ヒントを追加で求められた | 質問と方針。コードブロックの実装は出さない | 手書き |

---

## Output Format

日本語。次の形だけ。

```markdown
## 判定
Day N: Pass | Fail

## 完了条件
- [ ] 条件: Pass | Fail — 一言

## 指摘
- High | Medium | Low: ファイル — 何が方針に反するか（直す場所まで。実装は書かない）

## 次
Pass なら: 現在地を Day N+1 にするか確認する
Fail なら: 同じ Day を続ける
```

---

## Examples

- Good: 「`api/pkg/interface/handler` が SQL を持っている。SQL は `infra/repository` へ」
- Bad: repository の完成関数を貼る
- Bad: Day 21 の JWT 実装を Day 3 のレビューで要求する
