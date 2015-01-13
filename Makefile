LIBS = fbp
PROJECTS = brainfuck
BINDIR = $(shell pwd)/bin
PKGDIR = $(shell pwd)/pkg/linux_amd64

PKGS = $(patsubst %, $(PKGDIR)/%.a, $(LIBS))
BINS = $(patsubst %, $(BINDIR)/%, $(PROJECTS))

GOPATH = $(shell pwd)
GOBIN = $(BINDIR)

GO = /usr/local/go/bin/go
GOFMT = /usr/local/go/bin/gofmt
RM = /bin/rm

all: ${PKGS} ${BINS}

${PKGDIR}/fbp.a: src/fbp/*go
	env GOPATH=${GOPATH} GOBIN=${GOBIN} ${GO} install fbp

${BINDIR}/brainfuck: src/brainfuck/*.go ${PKGS}
	env GOPATH=${GOPATH} GOBIN=${GOBIN} ${GO} install brainfuck

test:
	env GOPATH=${GOPATH} ${GO} test ${PROJECTS}

format:
	${GOFMT} -w src/**/*.go

clean:
	${RM} ${BINS} ${PKGS}

.PHONY: all clean test
