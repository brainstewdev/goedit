# .PHONY all
# all : build/ build/config/ build/edit	
# 	./build/
# 	./build/config/
# 	./build/edit
.PHONY all: build/edit build/config/ 

clean:
	rm build/ -r
./build/config/: ./src/config/
	cp ./src/config ./build/config -r
./build/edit: ./src/edit.go
	go build  -o ./build/edit ./src/edit.go
