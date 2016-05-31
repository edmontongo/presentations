TARGETS := client server

.PHONY: all clean 

export GOPATH=$(CURDIR)

all: $(TARGETS)

%: %.go aux.go
	go build -o $@ $^

clean:
	$(RM) -vf $(TARGETS) ? *~
