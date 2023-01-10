package setting

import (
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func ReadFile(path, file, extension string) *Setting {
	setting := Setting{}

	setting.vp = viper.New()
	setting.vp.AddConfigPath(path)      // 设置读取的目录
	setting.vp.SetConfigName(file)      // 设置读取的文件名
	setting.vp.SetConfigType(extension) // 设置文件拓展名
	err := setting.vp.ReadInConfig()    //文件内容读取到内存
	if err != nil {
		panic(err)
	}

	return &setting
}
