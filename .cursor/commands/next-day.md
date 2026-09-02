---
name: next-day
description: >-
  Advances 現在地 and 進捗 in docs/LEARNING_ROADMAP.md by one Day,
  then asks that Day's 振り返り question in chat.
  Use when the current Day is done and the user invokes this command.
  Trigger: "next-day", "Day を進める", "現在地を次へ"
user-invocable: true
---

# Next Day — 現在地を 1 日進める

Use this command when the user marks the current Day complete and wants `docs/LEARNING_ROADMAP.md` updated.

---

## Iron Law(s)

> **1. `docs/LEARNING_ROADMAP.md` の「現在地」と「進捗」だけを書き換える。Day 本文・実装コードは触らない。振り返りの答えをファイルに書かない。**
> **2. 一度に 1 Day だけ進める。飛ばさない。**

---

## Workflow / Steps

1. `docs/LEARNING_ROADMAP.md` の「現在地」を読む。`今日やる日` を N、`最後に完了した日` を M とする。N が「なし」なら終了済みとして止める。
2. Day N セクションの「振り返り」にある質問を読む。
3. 「進捗」の `- [ ] Day N` を `- [x] Day N` にする。すでに `[x]` ならチェックは触らない。
4. `最後に完了した日` を `Day N` にする。
5. N が 41 なら `今日やる日` を `なし` にする。それ以外は `Day N+1` にする。
6. 変更後の 2 行と進捗の該当行だけ確認する。
7. 出力の「振り返り」に手順 2 の質問を載せ、このチャットで答えるよう促す。答えをファイルに書かない。ユーザーが答えなくても進めてよい。

---

## Decision Table

| Situation | Choice | Reason |
| --- | --- | --- |
| `今日やる日` が `なし` | ファイルを変えず、完了済みと伝える | これ以上ない |
| N が 41 | 完了を Day 41、今日を `なし` | 全体完了 |
| N が 1–40 | 完了を Day N、今日を Day N+1 | 1 日だけ進める |
| 進捗の Day N が既に `[x]` | 現在地だけ直す | 二重チェックしない |
| ユーザーがチャットで答えない | それでも進める。質問だけ出す | ブロックしない |
| ユーザーがチャットで答えた | 短いフィードバックのみ。ファイルに答えを残さない | 公開面に学習メモを増やさない |

---

## Output Format

日本語。次の形だけ。

```markdown
## 更新
今日やる日: <旧> → <新>
最後に完了した日: <旧> → <新>
進捗: Day <N> を [x]

## 振り返り（Day N）このチャットで答えてください
…
```
