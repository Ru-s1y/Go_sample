# Golang Web Application Sample

参考: 「Goプログラミング実践入門 標準ライブラリでゼロからWebアプリを作る」 3章〜

## PostgreSQL(Homebrew)
### ログイン
```
$ psql -U USERNAME
```
### 起動
```
$ brew service start postgresql
```

### 終了
```
$ brew service stop postgresql
```

### ユーザー作成
```
# create role ROLENAME with login password ‘PASSWORD’;
```

### データベース作成
```
# create database DATABASENAME;
```

### その他
#### データベース一覧
```
# \l
```

#### テーブル一覧
```
# \dt
```

今のところSQL文とかはMySQLと一緒みたい...