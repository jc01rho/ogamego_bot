module github.com/jc01rho/ogamego_bot

require (
	github.com/0xE232FE/ogame.mod v0.0.0
	github.com/SeitArt/autorestart v0.0.0-20200118174117-09bb20ae362c
	github.com/emirpasic/gods v1.12.0
	github.com/jc01rho/gocron v0.0.0-20200317110215-1a3ca30ce0f3
	github.com/labstack/echo v3.3.10+incompatible
	github.com/sirupsen/logrus v1.4.2
	github.com/slayer/autorestart v0.0.0-20170706172704-7bc8d250279b // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/tucnak/telebot.v2 v2.3.5
	gopkg.in/urfave/cli.v2 v2.0.0-20180128182452-d3ae77c26ac8
	gopkg.in/yaml.v2 v2.2.7 // indirect
)

replace github.com/0xE232FE/ogame.mod => ../ogame

go 1.14
