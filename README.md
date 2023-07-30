# 🍸gin-todo-api

## このリポジトリについて
Gin+GORM+PosgreSQLで作成したTodoAPIです。

APIの実行イメージは[powershell_curl_test.md](powershell_curl_test.md)をご確認下さい

## 動作確認方法
前提:<br>
Go 1.20とPostgreSQL 15.3のインストール<br>
PostgreSQLはポート5432に設定してください。<br>
1. ローカルのPostgreSQLに、databaseを作成します。
   database作成権限のあるuserでログインし、以下を実行します。
   ```
   create database todo;
   \c todo
   create user todo_user with password 'todo_user';
   create schema todo_user;
   grant all privileges on schema todo_user to todo_user;
   ```
1. このリポジトリをローカルにクローンしてください。
1. クローンしたディレクトリで以下を実行し、サーバを起動します。
   ```
   go run main.go
   ```
1. curlコマンド等を用いて、HTTPリクエストを行います。
   私がテストした際のコマンドやログは、powershell_curl_test.mdで確認できます。ただし、コマンドはPowershell向けであることに注意してください。
   GET以外のリクエストで期待されるJSONは、以下の形式です。
   ```
   { "title" : "test1" }
   ```
