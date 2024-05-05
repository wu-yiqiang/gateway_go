package config

type Configuration struct {
	App     App     `mapstructure:"app" json:"app" yaml:"app"`
	Log     Log     `mapstructure:"log" json:"log" yaml:"log"`
	Mysql   Mysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Jwt     Jwt     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	Storage Storage `mapstructure:"storage" json:"storage" yaml:"storage"`
	Swagger Swagger `mapstructure:"swagger" json:"swagger" yaml:"swagger"`
	Cluster Cluster `mapstructure:"cluster" json:"cluster" yaml:"cluster"`
	Kafka   Kafka   `mapstructure:"kafka" json:"kafka" yaml:"kafka"`
}

type App struct {
	Env        string `mapstructure:"env" json:"env" yaml:"env"`
	Port       string `mapstructure:"port" json:"port" yaml:"port"`
	AppName    string `mapstructure:"app_name" json:"app_name" yaml:"app_name"`
	AppUrl     string `mapstructure:"app_url" json:"app_url" yaml:"app_url"`
	StaticPort string `mapstructure:"static_port" json:"static_port" yaml:"static_port"`
}

type Log struct {
	Level      string `mapstructure:"level" json:"level" yaml:"level"`
	RootDir    string `mapstructure:"root_dir" json:"root_dir" yaml:"root_dir"`
	Filename   string `mapstructure:"filename" json:"filename" yaml:"filename"`
	Format     string `mapstructure:"format" json:"format" yaml:"format"`
	ShowLine   bool   `mapstructure:"show_line" json:"show_line" yaml:"show_line"`
	MaxBackups int    `mapstructure:"max_backups" json:"max_backups" yaml:"max_backups"`
	MaxSize    int    `mapstructure:"max_size" json:"max_size" yaml:"max_size"` // MB
	MaxAge     int    `mapstructure:"max_age" json:"max_age" yaml:"max_age"`    // day
	Compress   bool   `mapstructure:"compress" json:"compress" yaml:"compress"`
}

type Database struct {
	Driver              string `mapstructure:"driver" json:"driver" yaml:"driver"`
	Host                string `mapstructure:"host" json:"host" yaml:"host"`
	Port                int    `mapstructure:"port" json:"port" yaml:"port"`
	Database            string `mapstructure:"database" json:"database" yaml:"database"`
	UserName            string `mapstructure:"username" json:"username" yaml:"username"`
	Password            string `mapstructure:"password" json:"password" yaml:"password"`
	Charset             string `mapstructure:"charset" json:"charset" yaml:"charset"`
	MaxIdleConns        int    `mapstructure:"max_idle_conns" json:"max_idle_conns" yaml:"max_idle_conns"`
	MaxOpenConns        int    `mapstructure:"max_open_conns" json:"max_open_conns" yaml:"max_open_conns"`
	LogMode             string `mapstructure:"log_mode" json:"log_mode" yaml:"log_mode"`
	EnableFileLogWriter bool   `mapstructure:"enable_file_log_writer" json:"enable_file_log_writer" yaml:"enable_file_log_writer"`
	LogFilename         string `mapstructure:"log_filename" json:"log_filename" yaml:"log_filename"`
}

type Kafka struct {
	Topic       string `mapstructure:"topic" json:"topic" yaml:"topic"`
	Host        string `mapstructure:"host" json:"host" yaml:"host"`
	ChannelType string `mapstructure:"channel_type" json:"channel_type" yaml:"channel_type"`
}
type Mysql struct {
	Driver              string `mapstructure:"driver" json:"driver" yaml:"driver"`
	Host                string `mapstructure:"host" json:"host" yaml:"host"`
	Port                int    `mapstructure:"port" json:"port" yaml:"port"`
	Database            string `mapstructure:"database" json:"database" yaml:"database"`
	UserName            string `mapstructure:"username" json:"username" yaml:"username"`
	Password            string `mapstructure:"password" json:"password" yaml:"password"`
	Charset             string `mapstructure:"charset" json:"charset" yaml:"charset"`
	MaxIdleConns        int    `mapstructure:"max_idle_conns" json:"max_idle_conns" yaml:"max_idle_conns"`
	MaxOpenConns        int    `mapstructure:"max_open_conns" json:"max_open_conns" yaml:"max_open_conns"`
	LogMode             string `mapstructure:"log_mode" json:"log_mode" yaml:"log_mode"`
	EnableFileLogWriter bool   `mapstructure:"enable_file_log_writer" json:"enable_file_log_writer" yaml:"enable_file_log_writer"`
	LogFilename         string `mapstructure:"log_filename" json:"log_filename" yaml:"log_filename"`
}
type Redis struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}

type Jwt struct {
	Secret                  string `mapstructure:"secret" json:"secret" yaml:"secret"`
	JwtTtl                  int64  `mapstructure:"jwt_ttl" json:"jwt_ttl" yaml:"jwt_ttl"`                                                          // token 有效期（秒）
	JwtBlacklistGracePeriod int64  `mapstructure:"jwt_blacklist_grace_period" json:"jwt_blacklist_grace_period" yaml:"jwt_blacklist_grace_period"` // 黑名单宽限时间（秒）
}

type Storage struct {
	// Default storage.DiskName `mapstructure:"default" json:"default" yaml:"default"` // local本地 oss阿里云 kodo七牛云
	Disks Disks `mapstructure:"disks" json:"disks" yaml:"disks"`
}

type Disks struct {
	AliOss       Alioss       `mapstructure:"ali_oss" json:"ali_oss" yaml:"ali_oss"`
	LocalStorage LocalStorage `mapstructure:"local_storage" json:"local_storage" yaml:"local_storage"`
}

type LocalStorage struct {
	RootFileDir  string `mapstructure:"root_file_dir" json:"root_file_dir" yaml:"root_file_dir"`
	RootImageDir string `mapstructure:"root_image_dir" json:"root_image_dir" yaml:"root_image_dir"`
	RootVideoDir string `mapstructure:"root_video_dir" json:"root_video_dir" yaml:"root_video_dir"`
	AppUrl       string `mapstructure:"app_url" json:"app_url" yaml:"app_url"`
}

type Alioss struct {
	AccessKeyId     string `mapstructure:"access_key_id" json:"access_key_id" yaml:"access_key_id"`
	AccessKeySecret string `mapstructure:"access_key_secret" json:"access_key_secret" yaml:"access_key_secret"`
	Bucket          string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	Endpoint        string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	IsSsl           string `mapstructure:"is_ssl" json:"is_ssl" yaml:"is_ssl"`
	IsPrivate       string `mapstructure:"is_private" json:"is_private" yaml:"is_private"`
}

type Swagger struct {
	Title    string `mapstructure:"title" json:"title" yaml:"title"`
	Desc     string `mapstructure:"desc" json:"desc" yaml:"desc"`
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	BasePath string `mapstructure:"base_path" json:"base_path" yaml:"base_path"`
}

type Cluster struct {
	ClusterIp      string `mapstructure:"cluster_ip" json:"cluster_ip" yaml:"cluster_ip"`
	ClusterPort    string `mapstructure:"cluster_port" json:"cluster_port" yaml:"cluster_port"`
	ClusterSslPort string `mapstructure:"cluster_ssl_port" json:"cluster_ssl_port" yaml:"cluster_ssl_port"`
}
