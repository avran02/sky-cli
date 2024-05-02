package tpl

var Linter = `run:
	tests: false
	skip-files:
	  - \.pb\.go$
	  - \.pb\.validate\.go$
	  - \.pb\.gw\.go$
  
  linters-settings:
	dupl:
	  threshold: 500
	# commented linter rules should be uncommented when corresponding linter will be enabled
	# funlen:
	#   lines: 160
	#   statements: 60
	gci:
	  local-prefixes: github.com/golangci/golangci-lint
	# revive:
	#   min-confidence: 0
	# gomnd:
	#   checks:
	#     - argument
	#     - case
	#     - condition
	#     - return
	lll:
	  line-length: 180
  
  linters:
	# please, do not use 'enable-all': it's deprecated and will be removed soon.
	# inverted configuration with 'enable-all' and 'disable' is not scalable during updates of golangci-lint
	# Disabled linters should be enabled later
	disable-all: true
	enable:
	  - asciicheck
	  - bidichk
	  - bodyclose
	  - containedctx
	  # - deadcode
	  - decorder
	  # - depguard
	  - dogsled
	  - dupl
	  - durationcheck
	  - errcheck
	  - errchkjson
	  # - errname
	  - errorlint
	  - execinquery
	  - exportloopref
	  # - forbidigo
	  - forcetypeassert
	  # - funlen
	  - gocognit
	  - goconst
	  # - gocritic
	  - gocyclo
	  - gofmt
	  - gofumpt
	  - goheader
	  # - goimports
	  # - gomnd
	  - gomodguard
	  - goprintffuncname
	  - gosec
	  - gosimple
	  - govet
	  - grouper
	  # - ifshort
	  - importas
	  - ineffassign
	  - lll
	  # - maintidx
	  - makezero
	  # - misspell
	  - nakedret
	  - nestif
	  - nilerr
	  - nilnil
	  - noctx
	  # - nolintlint
	  - nosprintfhostport
	  - paralleltest
	  # - prealloc
	  # - predeclared
	  - promlinter
	  # - revive
	  - rowserrcheck
	  - sqlclosecheck
	  # - structcheck
	  # - stylecheck
	  # - tagliatelle
	  - tenv
	  - testpackage
	  - tparallel
	  - typecheck
	  - unconvert
	  - unparam
	  - unused
	  # - varcheck
	  - wastedassign
	  - whitespace
  
  issues:
	# Excluding configuration per-path, per-linter, per-text and per-source
	exclude-rules:
	  - path: _test\.go
		linters:
		  - gomnd
	  - text: "File is not 'goimports'-ed with -local github.com/golangci/golangci-lint"
		linters:
		  - goimports
`
