usecase

- 手順。作成・更新・削除・一覧を「どの順で呼ぶか」を書く。
- entity を組み立て、下の口（port）に渡す。
- Echo / SQL は書かない。

queryservice

- 参照専用の口。
- List / FindByID の型と interface だけ。
- SQL は infra。
- 今の queryservice.Todo に created_at があるのは、読む形だから。
- CQRS は、CURD のうち CUD と R を単純に分けるのではなく、R を参照専用の窓口として定義する
- だから entity は通さずに dto を使う。

authservice

- 認証の口。
- 同じ interface のまま HS256 → Cognito と実装だけ替える。
- JWT ライブラリは infra。
- Day 14 では触らない。
