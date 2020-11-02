# Fluentd out_httpの検証

## 使い方

まず環境を起動します。

```
docker-compose up
```

起動後、Nginxにリクエストを送ると、NginxのGETリクエストのログがGoのサーバーに送られます。

```
http localhost
```
