


all : build


build:
	git pull
	cd ../ogame && git pull
	sed -i "s/var IsDevelopment = true/var IsDevelopment = false/g" Logger/Logger.go
	go build -o ogamebot main.go
	git reset --hard HEAD

