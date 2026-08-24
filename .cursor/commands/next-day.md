---
name: next-day
description: >-
  Advances 現在地 and 進捗 in docs/LEARNING_ROADMAP.md by one Day.
  Use when the current Day is done and the user invokes this command.
  Trigger: "next-day", "Day を進める", "現在地を次へ"
user-invocable: true
---

# Next Day — 現在地を 1 日進める

Use this command when the user marks the current Day complete and wants `docs/LEARNING_ROADMAP.md` updated.

---

## Iron Law(s)

> **1. `docs/LEARNING_ROADMAP.md` の「現在地」と「進捗」だけを書き換える。Day 本文・実装コードは触らない。**
> **2. 一度に 1 Day だけ進める。飛ばさない。**

---

## Workflow / Steps

1. `docs/LEARNING_ROADMAP.md` の「現在地」を読む。`今日やる日` を N、`最後に完了した日` を M とする。N が「なし」なら終了済みとして止める。
2. 「進捗」の `- [ ] Day N` を `- [x] Day N` にする。すでに `[x]` ならチェックは触らない。
3. `最後に完了した日` を `Day N` にする。
4. N が 36 なら `今日やる日` を `なし` にする。それ以外は `Day N+1` にする。
5. 変更後の 2 行と進捗の該当行だけ確認する。

---

## Decision Table

| Situation | Choice | Reason |
| --- | --- | --- |
| `今日やる日` が `なし` | ファイルを変えず、完了済みと伝える | これ以上ない |
| N が 36 | 完了を Day 36、今日を `なし` | 全体完了 |
| N が 1–35 | 完了を Day N、今日を Day N+1 | 1 日だけ進める |
| 進捗の Day N が既に `[x]` | 現在地だけ直す | 二重チェックしない |

---

## Output Format

日本語。次の形だけ。

```markdown
## 更新
今日やる日: <旧> → <新>
最後に完了した日: <旧> → <新>
進捗: Day <N> を [x]
```
