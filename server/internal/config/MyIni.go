package config

import (
	"fmt"
)

// Config custom config struct
type Config struct {
	CfgBase     `yaml:"base"`
	MySQLInfo   MysqlDbInfo `yaml:"mysql_info"`
	WxInfo      WxInfo      `yaml:"wx_info"`
	FileHost    string      `yaml:"file_host"`
	Oauth2Url   string      `yaml:"oauth2_url"`
	RegisterURL string      `yaml:"register_url"` //注册或修改密码URL
	TokenType   string      `yaml:"token_type"`
	AppID       string      `yaml:"app_id"`
	AppSecret   string      `yaml:"app_secret"`
	Port        string      `yaml:"port"`   // 端口号
	Kdniao      Kdniao      `yaml:"kdniao"` // 快递鸟
	Email       Email       `yaml:"email"`
}

// MysqlDbInfo mysql database information. mysql 数据库信息
type MysqlDbInfo struct {
	Host     string `validate:"required"` // Host. 地址
	Port     int    `validate:"required"` // Port 端口号
	Username string `validate:"required"` // Username 用户名
	Password string // Password 密码
	Database string `validate:"required"` // Database 数据库名
}

// Kdniao 快递鸟
type Kdniao struct {
	EBusinessID string `yaml:"business_id"`
	AppKey      string `yaml:"app_key"`
}

// Email 邮件配置
type Email struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
}

// GetPort 获取端口号
func GetPort() string {
	return _map.Port
}

// SetMysqlDbInfo Update MySQL configuration information
func SetMysqlDbInfo(info *MysqlDbInfo) {
	_map.MySQLInfo = *info
}

// GetMysqlDbInfo Get MySQL configuration information .获取mysql配置信息
func GetMysqlDbInfo() MysqlDbInfo {
	return _map.MySQLInfo
}

// GetMysqlConStr Get MySQL connection string.获取mysql 连接字符串
func GetMysqlConStr() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&interpolateParams=True",
		_map.MySQLInfo.Username,
		_map.MySQLInfo.Password,
		_map.MySQLInfo.Host,
		_map.MySQLInfo.Port,
		_map.MySQLInfo.Database,
	)
}

///////////////////////////////
//// 微信   ///////////////////

// WxInfo 微信相关配置
type WxInfo struct {
	AppID     string `yaml:"app_id"`     // 微信公众平台应用ID
	AppSecret string `yaml:"app_secret"` // 微信支付商户平台商户号
	APIKey    string `yaml:"api_key"`    // 微信支付商户平台API密钥
	MchID     string `yaml:"mch_id"`     // 商户号
	NotifyURL string `yaml:"notify_url"` // 通知地址
	ShearURL  string `yaml:"shear_url"`  // 微信分享url
}

// GetWxInfo 获取微信配置
func GetWxInfo() WxInfo {
	return _map.WxInfo
}

///////////////////////////////

// GetFileHost 获取文件host
func GetFileHost() string {
	return _map.FileHost
}

// GetCheckTokenURL checktoken
func GetCheckTokenURL() string {
	return _map.Oauth2Url + "/check_token"
}

// GetLoginURL 登陆用的url
func GetLoginURL() string {
	return _map.Oauth2Url + "/authorize"
}

// GetLoginNoPwdURL token 授权模式登陆
func GetLoginNoPwdURL() string {
	return _map.Oauth2Url + "/authorize_nopwd"
}

// GetRefreshTokenURL 获取并且刷新 refreshToken
func GetRefreshTokenURL() string {
	return _map.Oauth2Url + "/refresh_token"
}

// GetTokenType 获取token类型
func GetTokenType() string {
	return _map.TokenType
}

// GetAppInfo 获取授权信息码
func GetAppInfo() (appid, appsecret string) {
	return _map.AppID, _map.AppSecret
}

// GetRegistURL ...
func GetRegistURL(url string) string {
	return _map.RegisterURL + url
}

// GetKdniao 获取快递鸟
func GetKdniao() Kdniao {
	return _map.Kdniao
}

// GetEmail 获取快递鸟
func GetEmail() Email {
	return _map.Email
}
