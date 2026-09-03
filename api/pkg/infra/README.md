repository
- DB との会話。
- 今ある Insert / Update / List の SQL がこれ。
- 口は usecase.TodoRepository と queryservice.TodoQueryService。

auth
- 認証基盤との会話。
- JWT 検証、JWKS、Cognito SDK。
- 口は authservice。
- 中身は後で HS256 → Cognito と差し替える。