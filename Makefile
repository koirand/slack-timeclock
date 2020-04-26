OUTPUT := slack-timeclock
ZIPFILE := lambda.zip

.PHONY: build
build: clean
	GOOS=linux GOARCH=amd64 go build -o $(OUTPUT) main.go
	zip $(ZIPFILE) $(OUTPUT) && rm $(OUTPUT)

.PHONY: clean
clean:
	rm -f $(ZIPFILE)
