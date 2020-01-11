# JWT (ジョット, Json Web Token)

クライアント&サーバー間で認証できる。



## JWTの特徴

特徴としては、

- 発行者が鍵を使ってJSONを署名し、トークンとして使う。
- JSONなので、任意の情報を含めることができる。
- 発行者は鍵を使ってトークンの検証を行う為、改竄の検知できる。
- ステートレスであり、セッションストアを持たなくても検証ができる。

## JWTのフォーマット

```
{base64エンコードしたheader}.{base64エンコードしたclaims}.{署名}
```

(例)

```
エンコードされている状態
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c

デコードされた状態
Header
{
  "alg": "HS256",
  "typ": "JWT"
}

claims(payload)
{
  "sub": "1234567890",
  "name": "John Doe",
  "iat": 1516239022
}

署名
HMACSHA256(
  base64UrlEncode(header) + "." +
  base64UrlEncode(payload),
  your-256-bit-secret
) secret base64 encoded

```

## ref

https://jwt.io/