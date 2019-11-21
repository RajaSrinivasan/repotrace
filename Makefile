SOURCES=$(wildcard *.go)
EXEC=repotrace
BINARIES=../bin
all: $(EXEC)

$(EXEC):
	go build -o $(BINARIES)/$(EXEC)

clean:
	$(RM) $(BINARIES)/$(EXEC)

generate:	
ifneq ("$(wildcard $(BINARIES)/$(EXEC))","")
	cd tests; ../$(BINARIES)/$(EXEC) --language go --output versions; mv versions.go ../versions/genversions.go
endif

dependencies:
	go get -u -v github.com/akamensky/argparse