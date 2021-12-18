# .PHONY all
# all : build/ build/config/ build/edit	
# 	./build/
# 	./build/config/
# 	./build/edit
.PHONY all: build/goedit build/config/ 

clean:
	rm build/ -r
./build/config/: ./src/config/
	cp ./src/config ./build/config -r
./build/goedit: ./src/goedit.go
	go build  -o ./build/goedit ./src/goedit.go
