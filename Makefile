# .PHONY all
# all : build/ build/config/ build/edit	
# 	./build/
# 	./build/config/
# 	./build/edit
.PHONY all: build/goedit build/config/ 
.PHONY install: all copy
.PHONY remove: remove_configuration remove_binary

clean:
	rm build/ -r
./build/config/: ./src/config/
	cp ./src/config ./build/config -r
./build/goedit: ./src/goedit.go
	go build  -o ./build/goedit ./src/goedit.go
copy: ./build/config ./build/goedit
	@echo "installing"
	cp ./build/goedit /usr/local/bin/goedit
	mkdir /usr/local/bin/config
	cp ./build/config/ /usr/local/bin/config -r
remove_configuration:
	@echo "removing configuration files"
	rm /usr/local/bin/config -r
remove_binary:
	@echo "removing binary files"
	rm /usr/local/bin/goedit
