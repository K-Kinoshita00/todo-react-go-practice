domain

- 「このアプリの業務そのもの」の層です。
- ToDo とは何か、何が正しいか（title 必須、status の意味）を置き、HTTP・SQL・DB ドライバは知りません。

entity

- domainの中の「1 件の本体」。
- 外側は application（何をするか）→ infra（保存）→ interface（HTTP）です。
- domain を真ん中に置くのは、DB や Echo を替えても「ToDo の規則」を動かさないためです。

domainservice

- 1 件の entity に載せられない業務規則です。
- HTTP / SQL は知りません（層の README どおり）。
- Todo の「title 空禁止」は entity で足ります。
- 複数件にまたがる判断（例: 同じ title の未完了を許さない）だけがここに来ます。
- Day 14 の CRUD では まだ書かなくてよい。
