# 学習ロードマップ

このファイルだけ読めば、翌日同じ方針で再開できます。コードは手書きし、AI にはレビューだけ依頼します。

## 現在地

再開したらここを見て、当日セクションだけ開く。

- 今日やる日: Day 4
- 最後に完了した日: Day 3

## 毎日の手順

1. このファイルの「現在地」と当日セクションだけ読む
2. 手元に `docs/private/` がある場合だけ、当日の対応表を見る
3. 手書きで実装する（AI にコードを書かせない）
4. 「レビュー」と依頼する（`.cursor/skills/review-day` が当日の完了条件で見る。実装は返さない）
5. 完了条件を満たしたら「現在地」を更新する

1 日は約 1 時間。終わらなければ同じ Day を翌日も続ける。先の Day は読まない。

## 設計方針

後から層を足さない。空でもディレクトリと責務は最初から置く。

### モノレポ

- `api/` — Go。`pkg/domain` / `application` / `infra` / `interface` / `registry`
- `web/` — Vite + React + MUI。`src/features` / `providers` / `lib`
- `openapi/` — 契約の単一ソース
- `database/migrations/` — SQL マイグレーション
- ルート `compose.yaml` と `Makefile` — ローカル一式

### 各層がしないこと

- `domain` — HTTP、SQL、AWS SDK を知らない
- `application` — Echo / React / 具体的な DB ドライバを知らない
- `infra` — ユースケース判断を持たない。I/O の実装だけ
- `interface` — ビジネスルールを持たない。HTTP の入出力とミドルウェアだけ
- `registry` — 生成と配線だけ。ロジックを書かない
- `web/src/features` — 他 feature の内部を直接 import しない
- `web/src/lib` — 画面固有の文言やレイアウトを持たない

### 認証（差し替え可能な AuthService）

同じ interface のまま実装だけ替える。

1. HS256 の自前 JWT（Bearer / claims / middleware を理解する）
2. LocalStack の Cognito User Pool + JWKS 検証
3. 実 AWS Cognito（同じ interface）

LocalStack コミュニティ版で Cognito が不足する場合は、無料枠の実 Cognito をローカルから使う。Day 23 で判断する。

### AWS（最後に載せる）

ローカルで完成してからクラウドへ行く。Day 31 の請求アラームを飛ばさない。

無料枠に寄せる最終形:

- Web: S3 + CloudFront
- 認証: Cognito（MAU 無料枠）
- API + DB: EC2 `t3.micro` / `t4g.micro` 1 台に Docker（API + Postgres）

ECS + ALB + RDS は理解用の短時間ラボにし、使い終わったら破棄する。常時起動しない。

## 全体の完了条件

- ログインしたユーザーが自分の ToDo を CRUD できる
- ローカルは Compose 一発で起動する
- PR で lint / test が回る
- 無料枠寄りの AWS に載せられる
- 公開面に、非公開の参考実装の名前・URL・アカウント情報がない

---

## Day 1 — ディレクトリ骨格

目標: 空でも責務が分かるモノレポの骨格を置く。

やること:

- 次のディレクトリを作る。中身は `.gitkeep` か、層の役割を 1 行書いた README でよい
  - `api/cmd/server/`
  - `api/pkg/domain/entity/`
  - `api/pkg/domain/domainservice/`
  - `api/pkg/application/usecase/`
  - `api/pkg/application/queryservice/`
  - `api/pkg/application/authservice/`
  - `api/pkg/infra/repository/`
  - `api/pkg/infra/auth/`
  - `api/pkg/interface/handler/`
  - `api/pkg/interface/middleware/`
  - `api/pkg/interface/gen/openapi/`
  - `api/pkg/registry/`
  - `web/src/features/todos/`
  - `web/src/providers/auth/`
  - `web/src/lib/axios/`
  - `web/src/lib/cognito/`
  - `web/src/lib/react-query/`
  - `openapi/`
  - `database/migrations/`
- 公開ファイルに、社内参考実装の名前や URL を書かない

完了条件:

- 上記ディレクトリが存在する
- 各層が「何をしないか」を自分の言葉で説明できる

翌日の入口: `compose.yaml`（未作成）

---

## Day 2 — PostgreSQL だけ起動する

目標: Compose で DB だけを安定起動する。

やること:

- ルートに `compose.yaml` を置く
- PostgreSQL 17 系のサービスを 1 つ定義する
- ポート、ユーザー、パスワード、DB 名は環境変数にする
- `.env` は git に入れない。`.env.example` にキーだけ書く
- `docker compose up -d` で起動し、`psql` または GUI で接続できることを確認する

完了条件:

- `localhost:5432` に接続できる
- コンテナ再起動後もボリュームでデータが残る

翌日の入口: `api/cmd/server/` と `compose.yaml`

---

## Day 3 — API の health

目標: API コンテナが `GET /health` を返す。

やること:

- `api/` に Go モジュールを初期化する
- `cmd/server` で HTTP サーバを立て、`GET /health` が `{"status":"ok"}` を返す
- API 用 Dockerfile を用意し、Compose から起動する
- ホットリロードは必須にしない。動けばよい
- domain / usecase にはまだロジックを書かない

完了条件:

- `curl http://localhost:8080/health` が 200 を返す
- API と DB が同じ Compose ネットワークにいる

翌日の入口: `web/`（未初期化）

---

## Day 4 — Vite + MUI の空画面

目標: Web コンテナで MUI の空画面が出る。

やること:

- `web/` を Vite + React + TypeScript で初期化する
- MUI を入れ、画面に見出し 1 つを出す
- Compose に web サービスを追加する（開発サーバ）
- API はまだ呼ばない

完了条件:

- ブラウザで `http://localhost:3000`（または決めたポート）が開く
- MUI のコンポーネントが 1 つ以上描画されている

翌日の入口: `Makefile`（未作成）

---

## Day 5 — Makefile と環境変数

目標: 日常操作を Makefile に寄せる。

やること:

- `make up` / `down` / `logs` / `ps` を定義する
- lint 用の空ターゲットを予約する（中身は後で足す）
- `.env.example` を完成させる（DB、API ポート、Web ポート）
- README に「ローカル起動は `make up`」と 1 行足してよい。手順の本体はこのファイルに残す

完了条件:

- `make up` で web / api / db が起動する
- シークレットが git に含まれない

翌日の入口: `openapi/openapi.yaml`

---

## Day 6 — OpenAPI に health

目標: 契約をコードより先に置く。

やること:

- `openapi/openapi.yaml` を OpenAPI 3 で作る
- `GET /health` だけ定義する
- 実装は当日は触らない。すでに動いている health とパスを揃える

完了条件:

- 仕様ファイル単体で health のリクエストとレスポンスが読める
- 実装のパスと仕様のパスが一致する

翌日の入口: `openapi/openapi.yaml`

---

## Day 7 — OpenAPI に todos CRUD

目標: ToDo の契約を仕様だけで固める。

やること:

- `Todo` スキーマ（id, title, completed, createdAt など最小）
- `GET /todos` / `POST /todos` / `GET /todos/{id}` / `PATCH /todos/{id}` / `DELETE /todos/{id}`
- エラーレスポンスの形を 1 つ決める
- 認証ヘッダはまだ必須にしない（Day 21 で足す）
- 実装は書かない

完了条件:

- CRUD 5 操作が仕様にある
- 後から実装する人が仕様だけ見てハンドラを書ける

翌日の入口: `api/pkg/interface/gen/openapi/`

---

## Day 8 — API コード生成

目標: 仕様からサーバの型とインタフェースを生成する。

やること:

- oapi-codegen（または同等）の設定を `api/` に置く
- 生成物は `api/pkg/interface/gen/openapi/` に出す
- 生成コマンドを Makefile に足す（例: `make generate-api`）
- 生成物は手で編集しない

完了条件:

- 生成された ServerInterface に todos のメソッドがある
- 生成をやり直しても差分が安定する

翌日の入口: `web/` の生成設定

---

## Day 9 — Web コード生成

目標: 同じ仕様からクライアント型を生成する。

やること:

- orval / openapi-typescript など、1 つ選んで固定する
- 生成物の置き場を `web/src/` 配下に決める
- `make generate-web` を足す
- 画面からはまだ呼ばない

完了条件:

- todos の型またはクライアント関数が生成されている
- API 側とパス・型名がずれていない

翌日の入口: `compose.yaml` の swagger サービス

---

## Day 10 — Swagger UI

目標: ブラウザで契約を確認できる。

やること:

- Compose に Swagger UI を追加する
- バンドル済み、または単一の `openapi.yaml` をマウントする
- `make up` で UI が開くようにする

完了条件:

- ブラウザで health と todos の定義が見える
- 「仕様が単一ソース」になっている

翌日の入口: `database/migrations/`

---

## Day 11 — マイグレーション

目標: `todos` テーブルをマイグレーションで作る。

やること:

- ツールは Flyway か golang-migrate のどちらかに固定する
- `database/migrations/` に CREATE TABLE を書く
- id, title, completed, created_at, updated_at
- Compose から migrate を実行できるようにする（Makefile 推奨）
- アプリの CRUD はまだ書かない

完了条件:

- 空の DB に対して migrate が通る
- 同じ migrate を再実行しても壊れない（バージョン管理されている）

翌日の入口: `api/pkg/infra/repository/`

---

## Day 12 — repository

目標: SQL を infra に閉じる。

やること:

- `todos` の insert / list / get / update / delete を repository に書く
- インタフェースは application 側に置くか、domain のポートにする
- handler から SQL を直接呼ばない
- まだ HTTP に繋がない。テスト用の main や一時スクリプトでもよい

完了条件:

- repository 単体（または一時コード）で 1 件 insert し、list で取れる
- SQL が handler / usecase に漏れていない

翌日の入口: `api/pkg/domain/entity/` と `api/pkg/application/usecase/`

---

## Day 13 — domain と usecase

目標: ユースケースが「何をするか」を持つ。

やること:

- `Todo` エンティティを domain に置く
- 作成・更新・削除・一覧の usecase を書く
- usecase は repository インタフェースに依存する
- バリデーションは最低限（title 空禁止など）
- HTTP にはまだ繋がない

完了条件:

- usecase のテストまたは手元実行で、作成した ToDo が一覧に出る
- usecase が `net/http` や Echo を import していない

翌日の入口: `api/pkg/interface/handler/` と `api/pkg/registry/`

---

## Day 14 — handler と registry

目標: 生成インタフェースを実装し、DI で配線する。

やること:

- handler が OpenAPI 生成の ServerInterface を満たす
- registry で db / repository / usecase / handler を組み立てる
- `cmd/server` は registry を呼ぶだけにする
- 認証は入れない

完了条件:

- サーバが生成インタフェース経由で todos を受け付ける
- 新規依存の追加場所が registry だと説明できる

翌日の入口: 起動中の API と Swagger UI

---

## Day 15 — curl で CRUD

目標: 未認証で ToDo の CRUD が通る。

やること:

- Swagger または curl で 5 操作を試す
- 存在しない id は 404
- 壊れた JSON は 400
- 動作確認で見つけたバグだけ直す。先の機能は足さない

完了条件:

- 作成・取得・更新・削除・一覧がすべて期待どおり
- DB を直接見ても行と一致する

翌日の入口: `web/src/features/todos/`

---

## Day 16 — 一覧 UI

目標: MUI で ToDo 一覧が出る。

やること:

- 一覧画面を 1 枚作る
- API の list を呼ぶ（fetch 直書きでよい。Query は Day 19）
- ローディング中の表示を 1 つ入れる
- 作成フォームはまだ作らない

完了条件:

- ブラウザを開くと、Day 15 で作ったデータが見える
- API 停止時に画面が白くならない（エラー表示でよい）

翌日の入口: 同じ一覧画面

---

## Day 17 — 作成 UI

目標: 画面から ToDo を追加できる。

やること:

- タイトル入力と送信ボタン
- 成功後に一覧を更新する
- 空タイトルは送れない

完了条件:

- 画面から追加した行が一覧に出る
- リロードしても残る

翌日の入口: 同じ一覧画面

---

## Day 18 — 更新と削除

目標: 完了トグルと削除ができる。

やること:

- チェックで `completed` を PATCH する
- 削除ボタンで DELETE する
- 確認ダイアログは任意。誤削除防止は推奨

完了条件:

- 完了状態がリロード後も保たれる
- 削除した行が一覧と DB から消える

翌日の入口: `web/src/lib/react-query/` と `web/src/features/todos/`

---

## Day 19 — Query と features 整理

目標: データ取得を hooks に寄せ、feature 境界を整える。

やること:

- TanStack Query を入れる
- list / create / update / delete を hooks にする
- コンポーネントから URL 文字列を消す
- `features/todos` 以外に ToDo 専用ロジックを置かない

完了条件:

- 画面コンポーネントが API パスを知らない
- 作成・更新後に一覧が自動で更新される

翌日の入口: 同じ画面の空・エラー

---

## Day 20 — 空とエラー

目標: 正常系以外でも使える。

やること:

- 0 件のときの表示
- API エラーの表示
- 通信中の二重送信防止（ボタン disable など）

完了条件:

- DB を空にすると空状態が見える
- API を止めるとエラー状態が見える
- 連打しても重複作成されない

翌日の入口: `api/pkg/application/authservice/` と `api/pkg/infra/auth/`

---

## Day 21 — HS256 JWT

目標: Bearer と claims を自前 JWT で理解する。

やること:

- 開発用に JWT を発行する小さな手段を用意する（スクリプトまたは一時エンドポイント）
- API middleware で `Authorization: Bearer` を検証する
- `/health` は検証しない。todos は検証する
- claims に `sub` を入れる
- OpenAPI に Bearer を足してよい

完了条件:

- トークンなしの todos は 401
- 正しいトークンで Day 15 と同じ CRUD ができる
- usecase が JWT ライブラリを直接 import していない（infra / middleware 側）

翌日の入口: `web/src/lib/axios/` と `web/src/providers/auth/`

---

## Day 22 — Web が Bearer を付ける

目標: 画面のリクエストにトークンが乗る。

やること:

- axios（または fetch ラッパ）の interceptor で Authorization を付ける
- 開発中は発行した JWT を手元でセットしてよい（ログイン画面は Day 25）
- 401 のときの表示を 1 つ決める

完了条件:

- ブラウザのネットワークタブで Bearer が見える
- トークンを外すと一覧が 401 になる

翌日の入口: `compose.yaml` の LocalStack

---

## Day 23 — LocalStack と User Pool

目標: ローカルで Cognito 相当の User Pool を立てる。

やること:

- Compose に LocalStack を追加する
- 初期化スクリプトで User Pool と App Client を作る
- テストユーザーを 1 人作る
- 発行した IdToken / AccessToken を控える
- コミュニティ版で User Pool が作れない場合は、無料枠の実 Cognito をローカルから使う（フォールバック）。選んだ理由を自分用メモに残す（公開ファイルにアカウント ID を書かない）

完了条件:

- トークンを 1 つ取得できる
- jwt の issuer / audience が説明できる

翌日の入口: `api/pkg/infra/auth/`

---

## Day 24 — JWKS 検証へ差し替え

目標: AuthService の実装を JWKS 検証に替える。

やること:

- JWKS から署名を検証する
- iss / aud / exp を見る
- HS256 実装は残してもよいが、起動設定で切り替える
- handler / usecase は触らない（触るなら registry と infra だけ）

完了条件:

- Day 23 のトークンで todos が通る
- 改ざんトークンは 401
- 「検証実装だけ替えた」と説明できる

翌日の入口: `web/src/lib/cognito/` と `web/src/providers/auth/`

---

## Day 25 — ログイン画面

目標: 画面からサインインし、トークンを保持する。

やること:

- メール（またはユーザー名）とパスワードのフォーム
- Cognito クライアント（amazon-cognito-identity-js など）でサインインする
- トークンをメモリまたはセッションストレージに持つ。永続方針を 1 つ決める
- 未ログインならログインへ寄せる

完了条件:

- ログイン後に一覧が見える
- ログアウトすると一覧 API が 401 になる

翌日の入口: `todos` テーブルと usecase

---

## Day 26 — 自分の ToDo だけ

目標: `sub` で所有者を分ける。

やること:

- `todos` に owner（`sub`）列を足すマイグレーション
- 作成時に `sub` を保存する
- list / get / update / delete は自分の行だけ
- 他人の id を指定したら 404（存在を漏らさない）

完了条件:

- ユーザー A の ToDo がユーザー B に見えない
- 既存の自分の操作はこれまでどおり動く

翌日の入口: `api/pkg/interface/middleware/`

---

## Day 27 — バリデーション、CORS、ヘッダ

目標: 公開前の最低限の防御を入れる。

やること:

- OpenAPI のバリデーションをサーバで効かせる（または同等の入力検証）
- CORS を Web の origin だけに絞る
- セキュリティヘッダ（少なくとも `X-Content-Type-Options` とフレーム抑制）
- エラー本文に内部情報（SQL、スタック）を出さない

完了条件:

- 不正な body が 400 になる
- 許可していない origin からブラウザは呼べない
- エラー JSON に内部実装が混ざらない

翌日の入口: `api/pkg/application/usecase/` のテスト

---

## Day 28 — usecase テスト

目標: ビジネスルールを HTTP なしで固定する。

やること:

- repository をモックまたはフェイクにする
- 作成（空タイトル拒否）と「他人の ToDo は更新できない」をテストする
- テストは `api/` で `go test` できるようにする

完了条件:

- `go test` が usecase を通す
- 失敗メッセージで何が壊れたか分かる

翌日の入口: `Makefile` の format / lint / 型チェック

---

## Day 29 — format / lint / 型チェック

目標: CI に載せるコマンドをローカルで固定する。

やること:

- API: `gofmt` と `go vet`（golangci-lint は任意）
- Web: format、lint、`tsc --noEmit`
- Makefile の `format` / `lint` / `test` を実装する
- 生成物は lint 対象から外す

完了条件:

- `make format && make lint && make test` が通る
- 意図的に壊すとどれかが失敗する

翌日の入口: `.github/workflows/`

---

## Day 30 — GitHub Actions

目標: PR で品質チェックが回る。

やること:

- pull_request で lint / test を回す
- Node / Go のバージョンを固定する
- シークレットを workflow に直書きしない
- デプロイはまだしない

完了条件:

- 空の PR または同じブランチで workflow が成功する
- フォーマットを崩すと CI が落ちる

翌日の入口: AWS アカウントの請求設定（コンソール）

---

## Day 31 — 請求アラームと最小 IAM

目標: 課金を先に止める仕組みを作る。

やること:

- 請求アラーム（少額、例: 1 USD と 5 USD）を入れる
- ルートユーザーで日常作業しない。作業用 IAM ユーザーまたは Identity Center
- 使うサービスだけに権限を絞る方針をメモする（公開ファイルにアカウント ID を書かない）
- この Day が終わるまで他の AWS リソースを作らない

完了条件:

- アラームが「OK」で存在する
- 作業用の人（またはロール）でコンソールに入れる
- ルートのアクセスキーを作っていない

翌日の入口: Cognito コンソール

---

## Day 32 — 実 Cognito

目標: ローカルアプリを実 User Pool に向ける。

やること:

- User Pool と公開クライアントを 1 つ作る
- ローカル Web / API の環境変数だけを差し替える
- テストユーザーでログインする
- プール ID を README に書かない。`.env` のみ

完了条件:

- 実 Cognito のトークンでローカル todos が通る
- LocalStack 用設定と実 Cognito 用設定が切り替えできる

翌日の入口: `web/` の本番ビルド

---

## Day 33 — S3 と CloudFront

目標: 静的 Web を安く公開する。

やること:

- `web` を本番ビルドする
- S3 に置き、CloudFront で配信する
- バケットは公開設定にせず、CloudFront 経由だけにする
- API の origin をビルド時環境変数で渡す

完了条件:

- CloudFront の URL でログイン画面が出る
- S3 を直接匿名 GET できない

翌日の入口: EC2（または同等の 1 台）

---

## Day 34 — API を安価ホスト

目標: API と Postgres を 1 台に載せる。

やること:

- `t3.micro` または `t4g.micro` を 1 台（無料枠対象を確認する）
- Docker で api + db を動かす
- セキュリティグループは自分の IP と CloudFront / Web から必要なポートだけ
- HTTPS は、手元のドメインがある場合のみ。なければ HTTP + 制限 IP で学習し、公開範囲を狭くする
- ECS + ALB + RDS は、この Day の本線にしない。やるなら別ラボとして作り、当日中に破棄する

完了条件:

- 公開 Web から API の health が通る
- インスタンスの hourly 料金が何か言える

翌日の入口: `.github/workflows/`

---

## Day 35 — GitHub OIDC デプロイ

目標: 長期アクセスキーなしでデプロイする。

やること:

- GitHub Actions から AWS へ OIDC で入る
- Web の S3 同期（と CloudFront 無効化）を workflow にする
- API 側はイメージ転送または `scp` / SSM など、キーをリポジトリに置かない方法を 1 つ選ぶ
- 環境ごとのシークレットは GitHub Secrets または OIDC ロールに閉じる

完了条件:

- 手元に AWS アクセスキーを置かずに main（または決めた枝）へ push して Web が更新される
- workflow ログにシークレットが出ない

翌日の入口: 請求コンソールと作ったリソース一覧

---

## Day 36 — 費用確認と破棄手順

目標: 残すものと消すものを固定する。

やること:

- コストエクスプローラーまたは請求を見る
- 残すもの: Cognito / S3 / CloudFront / 必要な EC2 1 台
- 消すもの: 実験用の ALB、NAT、RDS、余った EBS、未使用 Elastic IP
- 破棄手順を `docs/LEARNING_ROADMAP.md` のこの節に、サービス名だけで追記する（ID は書かない）
- 止め方（EC2 stop）と消し方（destroy）を区別する

完了条件:

- 不要リソースが残っていない
- 翌月も残す場合の概算が言える
- 「現在地」を Day 36 完了にする

翌日の入口: なし（全体完了）。必要なら下の「全体完了後に足してよいもの」へ。

---

## 全体完了後に足してよいもの

Day 36 が終わるまで読まない。本線に載せず、層を壊さない範囲で 1 つずつ足す。

### 公開の穴

- nginx または Caddy を EC2 に置く（TLS 終端、API へのリバプロ）。ローカルの CORS 学習は残す
- CloudFront に独自ドメインと ACM の HTTPS
- API のレート制限と、エラー本文を増やさないままの構造化ログ
- Postgres のバックアップとリストア手順（ID は公開ファイルに書かない）

### 認証の次

- サインアップ、メール確認、パスワードリセット
- リフレッシュトークンと、期限切れ時の再取得
- ログアウトでサーバ側セッションまたはトークンを無効化するなら、その方針を 1 つに決める

### プロダクト

- 一覧のページネーション、フィルタ、ソート（先に OpenAPI を足す）
- 期限やメモなど、ToDo の最小列以外を 1 つだけ足す
- 空・エラーのあとに、オフラインや再試行の表示

### 品質と運用

- Playwright など E2E（ログインして CRUD まで）
- staging を prod と分ける（プール、バケット、EC2 を混ぜない）
- ECS + ALB + RDS は短時間ラボ。作り終わったら破棄する。常時起動しない

やらない方がよいもの: 本線の層をまたぐ巨大な管理画面、課金を無視した常時 ALB / NAT / RDS。

---

## 進捗

- [x] Day 1 ディレクトリ骨格
- [x] Day 2 PostgreSQL
- [x] Day 3 API health
- [ ] Day 4 Vite + MUI
- [ ] Day 5 Makefile
- [ ] Day 6 OpenAPI health
- [ ] Day 7 OpenAPI todos
- [ ] Day 8 API コード生成
- [ ] Day 9 Web コード生成
- [ ] Day 10 Swagger UI
- [ ] Day 11 マイグレーション
- [ ] Day 12 repository
- [ ] Day 13 domain / usecase
- [ ] Day 14 handler / registry
- [ ] Day 15 curl CRUD
- [ ] Day 16 一覧 UI
- [ ] Day 17 作成 UI
- [ ] Day 18 更新・削除
- [ ] Day 19 Query / features
- [ ] Day 20 空とエラー
- [ ] Day 21 HS256 JWT
- [ ] Day 22 Web Bearer
- [ ] Day 23 LocalStack User Pool
- [ ] Day 24 JWKS 検証
- [ ] Day 25 ログイン画面
- [ ] Day 26 所有者分離
- [ ] Day 27 バリデーション / CORS / ヘッダ
- [ ] Day 28 usecase テスト
- [ ] Day 29 format / lint / 型チェック
- [ ] Day 30 GitHub Actions
- [ ] Day 31 請求アラーム / IAM
- [ ] Day 32 実 Cognito
- [ ] Day 33 S3 / CloudFront
- [ ] Day 34 安価ホスト
- [ ] Day 35 OIDC デプロイ
- [ ] Day 36 費用確認と破棄
