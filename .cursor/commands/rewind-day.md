---
name: rewind-day
description: >-
  Reverts 現在地 and 進捗 in docs/LEARNING_ROADMAP.md by one Day.
  Use when next-day was invoked by mistake or the last completed Day should be redone.
  Trigger: "rewind-day", "Day を戻す", "1日戻す"
user-invocable: true
---

# Rewind Day — 現在地を 1 日戻す

Use this command when the user wants to undo the last Day advance in `docs/LEARNING_ROADMAP.md`.

---

## Iron Law(s)

> **1. `docs/LEARNING_ROADMAP.md` の「現在地」と「進捗」だけを書き換える。Day 本文・実装コードは触らない。**
> **2. 一度に 1 Day だけ戻す。**

---

## Workflow / Steps

1. `docs/LEARNING_ROADMAP.md` の「現在地」を読む。`最後に完了した日` を M とする。M が「なし」なら初日なので止める。
2. 「進捗」の `- [x] Day M` を `- [ ] Day M` にする。
3. `今日やる日` を `Day M` にする。
4. M が 1 なら `最後に完了した日` を `なし` にする。それ以外は `Day M-1` にする。
5. 変更後の 2 行と進捗の該当行だけ確認する。

---

## Decision Table

| Situation | Choice | Reason |
| --- | --- | --- |
| `最後に完了した日` が `なし` | ファイルを変えず、初日だと伝える | これ以上戻らない |
| M が 1 | 今日を Day 1、完了を `なし`、Day 1 を `[ ]` | 開始状態 |
| M が 2–41 | 今日を Day M、完了を Day M-1、Day M を `[ ]` | 1 日だけ戻す |
| `今日やる日` が `なし`（全体完了後） | 上と同じ（M は 41） | 最終日に戻す |

---

## Output Format

日本語。次の形だけ。

```markdown
## 更新
今日やる日: <旧> → <新>
最後に完了した日: <旧> → <新>
進捗: Day <M> を [ ]
```
