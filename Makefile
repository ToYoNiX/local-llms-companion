BUILD_DIR := "./bin"

build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR) ./...

clean:
	rm -rf $(BUILD_DIR)
