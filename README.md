# Google Login Backend with Go

このプロジェクトは、Goで構築されたシンプルなバックエンドAPIで、Googleログインを処理します。フロントエンドからGoogleトークンを受け取り、それを検証して、有効であればJWTトークンを発行してフロントエンドに返します。

## 機能

- フロントエンドからGoogleトークンを受け取る
- Googleトークンを検証し、信頼できるソースかを確認
- 検証が成功した場合、新しいトークンを発行し、フロントエンドに返す

## 関連リポジトリ

- [フロントエンドリポジトリ](https://github.com/Tsubasa-2005/google-auth) - このバックエンドAPIと接続するフロントエンドアプリケーション。

## セットアップ

1. **リポジトリをクローン**:

   ```
   git clone https://github.com/Tsubasa-2005/GoogleAuthAPI.git
   cd GoogleAuthAPI
   ```

2. **依存関係のインストール**:

   ```
   go mod tidy
   ```

3. **環境変数の設定**:

   プロジェクトルートに `.env` ファイルを作成し、以下の変数を設定します。

   ```
    DB_PASS=postgres
    INSTANCE_UNIX_SOCKET=localhost
    DB_NAME=google_auth
    SECRET_KEY=secret
    CORS=http://localhost:3000
    PASSWORD=password
    GOOGLE_CLIENT_ID=your_google_client_id
   ```

4. **サーバーの実行**:

   ```
   make run
   ```