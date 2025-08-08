BINARY_NAME=gitstats
INSTALL_PATH=/usr/local/bin

.PHONY: build
build:
	go build -o $(BINARY_NAME)

.PHONY: install
install: build
	mv $(BINARY_NAME) $(INSTALL_PATH)/$(BINARY_NAME)

.PHONY: uninstall
uninstall:
	rm -f $(INSTALL_PATH)/$(BINARY_NAME)

.PHONY: clean
clean:
	go clean
	rm -f $(BINARY_NAME)
