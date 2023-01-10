package setting

import "time"

type ServerSetting struct {
	Url            string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	HandlerTimeout time.Duration
}

func (s *Setting) SetupServer() *ServerSetting {
	var setting = ServerSetting{}
	err := s.vp.UnmarshalKey("Server", &setting) // 将读取到的内容反射到数据结构中
	if err != nil {
		panic(err)
	}

	// 读取进来的默认单位是ns，需要进行转换
	setting.ReadTimeout *= time.Second
	setting.WriteTimeout *= time.Second
	setting.HandlerTimeout *= time.Second
	return &setting
}

type MongoSetting struct {
	Url       string
	User      string
	Password  string
	Db        string
	Article   string
	AccessLog string
	Timeout   time.Duration
}

func (s *Setting) SetupMongo() *MongoSetting {
	var setting = MongoSetting{}
	err := s.vp.UnmarshalKey("Mongo", &setting)
	if err != nil {
		panic(err)
	}

	setting.Timeout *= time.Second
	return &setting
}

////////////////////////////

type MySQLSetting struct {
	User     string
	Password string
	Url      string
	DBName   string
	Timeout  time.Duration
}

func (s *Setting) SetupMySQL() *MySQLSetting {
	var setting = MySQLSetting{}
	err := s.vp.UnmarshalKey("MySQL", &setting)
	if err != nil {
		panic(err)
	}
	setting.Timeout *= time.Second
	return &setting
}

///////////////////////////////////

type JWTSetting struct {
	Secret string        // 密匙
	Issuer string        // 发行人
	Expire time.Duration // 有效时间段
}

func (s *Setting) SetupJWT() *JWTSetting {
	var setting = JWTSetting{}
	err := s.vp.UnmarshalKey("JWT", &setting)
	if err != nil {
		panic(err)
	}
	setting.Expire *= time.Second
	return &setting
}
