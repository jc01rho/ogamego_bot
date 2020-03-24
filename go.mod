module github.com/jc01rho/ogamego_bot

require (
	bitbucket.org/jc01rho/ogame v0.0.0
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/emirpasic/gods v1.12.0
	github.com/jc01rho/gocron v0.0.0-20200317110215-1a3ca30ce0f3
	github.com/labstack/echo v3.3.10+incompatible
	github.com/sirupsen/logrus v1.4.2
	github.com/teamwork/reload v1.3.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/tucnak/telebot.v2 v2.0.0-20191225234705-baa616bc00d5
	gopkg.in/urfave/cli.v2 v2.0.0-20180128182452-d3ae77c26ac8
	gopkg.in/yaml.v2 v2.2.7 // indirect
)

replace bitbucket.org/jc01rho/ogame => ../ogame

go 1.14
