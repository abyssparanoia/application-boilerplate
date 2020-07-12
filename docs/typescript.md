# typescript 環境について

- lerna と yarn の workspace で各 pakcage を管理します。

## 初期化

- firebase の client 用の認証用情報を `packages/firebase-client/firebase.client.key.json` として配置する (firebase の管理画面から取得できます。)
- firebase の権限がついた、gcp の serviceAccount を serviceAccount.json として配置する。

```bash

# direnvの準備
# パスとかは適宜書き換えてください
> cp .envrc.tepl .envrc

> docker-compose build

# dotenvのコピー
> cp packages/web/.env.tepl packages/web/.env

> yarn

> yarn build
```

## サーバースタート

```bash

# サーバー側の起動
> docker-compose up -d

# next serverの起動
> yarn workspace @boilerplate/web start:dev

```
