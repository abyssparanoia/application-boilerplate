# development

開発フローについて、記載します。

## 概要

本プロジェクトは基本的に swagger 駆動でフロントとバックは IF のやりとりをします。

## フロー

### swagger

https://stoplight.io/

上記のアプリをダウンロードし、swagger 配下で開いてください。
swagger.yaml を GUI 経由で編集することができます。

### コード生成

yaml をいじった場合、次の作業をする必要があります。

- golang 側
  model となるコードを生成します

```bash
> make swaggergen
```

- typescript
  httpclient を openapi-generator によって自動生成します。

```bash

> brew install openapi-generator

> yarn swagger:generate
```

### ERD

データベースを参照して、ERD を自動生成します。

```bash
> docker-compose up default-schema
```

docs/er/default 配下に生成された index.html をブラウザで開いてください
