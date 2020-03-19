


all : build


build:
	cd ../ogame && git pull
	sed -i "s/var IsDevelopment = true/var IsDevelopment = false/g" Logger/Logger.go
	go build main.go ogamebot
