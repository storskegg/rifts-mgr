# Go parameters
PROJECTNAME=$(shell basename "$(PWD)")

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
DEPCMD=dep
DEPENSURE=$(DEPCMD) ensure -v
DEPUPDATE=$(DEPENSURE) -update -v
BINPATH=bin
BINARY=$(PROJECTNAME)

all: clean dep test build
build:
	    cd $(PROJECTNAME) && $(GOBUILD) -o ../$(BINPATH)/$(BINARY) -v
test:
		$(GOTEST) -v ./...
clean:
		$(GOCLEAN)
		rm -f $(BINPATH)/$(BINARY)
dep:
		$(DEPENSURE)
update:
		$(DEPUPDATE)
curdir:
		@echo $(PROJECTNAME)
