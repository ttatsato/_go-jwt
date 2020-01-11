package auth

// 環境変数にSIGNINGKEYをセットしておく
// in cli: export SIGNINGKEY=thisissecretkey
import (
	"net/http"
	"os"
	"time"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
)

var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	// headerのセット
	// ヘッダ情報には署名アルゴリズムの種類やメタ情報が組み込まれる。
	// JWTとHS256を指定している。
	// alg ... 署名アルゴリズム HS256
	// typ ... トークンのタイプ JWT
	token := jwt.New(jwt.SigningMethodHS256)

	/*
	 * claimsのセット
	 * claimsにはユーザー属性情報を埋め込んでいます。
	 * sub	ユーザーの一意識別子
	 * name	フルネーム
	 * profile	プロフィールページの URL
	 * picture	プロフィール画像の URL
	 * website	ウェブサイトもしくはブログの URL
	 * email	メールアドレス
	 * gender	性別
	 * birthdate 誕生日
	 * zoneinfo	地域情報
	 * phone_number	電話番号
	 * address	住所
	 * updated_at	情報最終更新時刻
	 */
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["sub"] = "54545454545"
	claims["name"] = "taro"
	claims["iat"] = time.Now()
	claims["epx"] = time.Now().Add(time.Hour * 24).Unix()

	// 電子署名
	tokenString, _ := token.SignedString([]byte(os.Getenv("SIGNINGKEY")))

	// JWTを返却
	w.Write([]byte(tokenString))
})

// JwtMiddleware check token
var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNINGKEY")), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})