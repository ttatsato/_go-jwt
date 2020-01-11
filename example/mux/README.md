# 手順
JWTを取得
```
curl localhost:8080/auth
```

JWTをセットしてcurlを実行
```
curl localhost:8080/private -H "Authorization:Bearer <先ほど返ってきたJWT>"
```