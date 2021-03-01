module github.com/neoaggelos/juju_exporter

go 1.15

// Juju dependencies
replace github.com/altoros/gosigma => github.com/juju/gosigma v0.0.0-20200420012028-063911838a9e

replace gopkg.in/natefinch/lumberjack.v2 => github.com/juju/lumberjack v2.0.0-20200420012306-ddfd864a6ade+incompatible

replace gopkg.in/mgo.v2 => github.com/juju/mgo v2.0.0-20201106044211-d9585b0b0d28+incompatible

replace github.com/hashicorp/raft => github.com/juju/raft v2.0.0-20200420012049-88ad3b3f0a54+incompatible

replace gopkg.in/yaml.v2 => github.com/juju/yaml v0.0.0-20200420012109-12a32b78de07

replace github.com/dustin/go-humanize v1.0.0 => github.com/dustin/go-humanize v0.0.0-20141228071148-145fabdb1ab7

replace github.com/hashicorp/raft-boltdb => github.com/juju/raft-boltdb v0.0.0-20200518034108-40b112c917c5

// Temporarily pin goreleaser because of errors with io/fs
replace github.com/goreleaser/goreleaser v0.158.0 => github.com/goreleaser/goreleaser v0.157.0

require (
	github.com/Djarvur/go-err113 v0.1.0 // indirect
	github.com/golangci/golangci-lint v1.36.0 // indirect
	github.com/golangci/misspell v0.3.5 // indirect
	github.com/golangci/revgrep v0.0.0-20180812185044-276a5c0a1039 // indirect
	github.com/goreleaser/goreleaser v0.158.0
	github.com/gostaticanalysis/analysisutil v0.6.1 // indirect
	github.com/jirfag/go-printf-func-name v0.0.0-20200119135958-7558a9eaa5af // indirect
	github.com/juju/juju v0.0.0-20201113194435-339998a462e3
	github.com/juju/names/v4 v4.0.0-20200929085019-be23e191fee0
	github.com/matoous/godox v0.0.0-20200801072554-4fb83dc2941e // indirect
	github.com/prometheus/client_golang v1.7.1
	github.com/quasilyte/go-ruleguard v0.2.1 // indirect
	github.com/quasilyte/regex/syntax v0.0.0-20200805063351-8f842688393c // indirect
	github.com/spf13/afero v1.5.1 // indirect
	github.com/stretchr/objx v0.3.0 // indirect
	github.com/tdakkota/asciicheck v0.0.0-20200416200610-e657995f937b // indirect
	github.com/timakin/bodyclose v0.0.0-20200424151742-cb6215831a94 // indirect
	github.com/tomarrell/wrapcheck v0.0.0-20201130113247-1683564d9756 // indirect
	gopkg.in/yaml.v2 v2.4.0
)
