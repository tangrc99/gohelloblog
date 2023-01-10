package app

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/tangrc99/gohelloblog/global"
	"time"
)

//// Claims 是一些实体（通常指的用户）的状态和额外的元数据
//type Claims struct {
//	Username string `json:"username"`
//	Password string `json:"password"`
//	jwt.StandardClaims
//}

type JWT struct {
	Secret string
	Issuer string
	Expire int
}

// Claims 是一些实体（通常指的用户）的状态和额外的元数据
type Claims struct {
	IsAdmin bool `json:"is_admin"`
	jwt.StandardClaims
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

func GenerateToken(isAdmin bool) (string, error) {
	//设置token有效时间
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)

	claims := Claims{
		IsAdmin: isAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),        // 过期时间
			Issuer:    global.JWTSetting.Issuer, // 指定token发行人
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

func ParseToken(token string) (*Claims, error) {

	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})

	if tokenClaims != nil {
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		// 要传入指针，项目中结构体都是用指针传递，节省空间。
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err

}
