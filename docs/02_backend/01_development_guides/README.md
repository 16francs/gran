# Backend

## 採用技術

* 言語等
  * 言語: Go
  * フレームワーク: gin
  * ライブラリ: firebase, colog, validator, xerrors

* アーキテクチャ
  * オニオンアーキテクチャ
    * [Qiita: ドメイン駆動+オニオンアーキテクチャ概略](https://qiita.com/little_hand_s/items/2040fba15d90b93fc124)
    * [Qiita: goプロジェクトにオニオンアーキテクチャを導入した](https://qiita.com/nanamen/items/f37d1047368929e377fd)

## 開発ルール

### 開発手順

1. Swaggerの作成
2. Domain層の作成
3. Infrastructure層の作成
4. Application層の作成
4. Interface層の作成
5. テストを作成
  * application
  * application/validation
  * domain/service
  * infrastructure/validation

### 実装詳細

* [ディレクトリ関連図](https://github.com/16francs/gran/tree/master/doc/02_backend/01_development_rules/directory.md)

### Application/Validation と Domain/Validation について

* Application/Validation
  * 入力フォームに関する制約についてのバリデーション
  * バリデーション情報は、Application/Requestの構造体に記載

* Domain/Validation
  * データベースの設計に関連する制約についてのバリデーション
  * バリデーション情報は、Domainの構造体に記載

* 文字列の長さ等のApplication,Domainどちらのバリデーションにも含められるもの
  * 原則、Domain/Valiadtion にのみ記載する形とする

## その他

### マイクロサービス化の粒度

* [User Service](https://github.com/16francs/gran/tree/master/doc/02_backend/03_user_api)

### ディレクトリ構成

```sh
sample_service
├── cmd
├── config
├── internal
│   ├── application
│   │   ├── request
│   │   ├── response
│   │   └── validation
│   ├── domain
│   │   ├── repository
│   │   ├── service
│   │   └── validation
│   ├── infrastructure
│   │   ├── persistence
│   │   └── validation
│   └── interface
│       └── handler
│           └── v1
├── lib
│   └── firebase
│       ├── authentication
│       └── firestore
└── registry
```