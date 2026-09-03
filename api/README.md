# バックエンドmemo

## クリーンアーキテクチャ

### 層

| 層             | このリポジトリ                                | 流れでの役割                 |
| -------------- | --------------------------------------------- | ---------------------------- |
| presentation   | interface/handler                             | HTTP を usecase の引数にする |
| application    | usecase + 口（TodoRepository / queryservice） | 何をするか。口を定義する     |
| domain         | entity（必要なら domainservice）              | 書きの規則。読みは通さない   |
| infrastructure | infra/repository                              | 口を SQL で実装する          |

読みが domain を飛ばすのは CQRS どおりで、層が欠けているのではない。

読み取り時にはバリデーション等が必要ないので、domain を飛ばす。

### 手順

書き:

handler（interface）→ usecase.Create（application）→ entity.NewTodo（domain）→ TodoRepository.Insert（口）→ todo_command（infra）

読み:

handler → usecase.List → queryservice.List（どちらも application）→ todo_query（infra）→ dto