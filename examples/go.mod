module github.com/pulumi/pulumi-docker/examples/v2

go 1.14

require (
	github.com/onsi/ginkgo v1.12.0 // indirect
	github.com/onsi/gomega v1.9.0 // indirect
	github.com/pulumi/pulumi-docker/sdk/v2 v2.2.3
	github.com/pulumi/pulumi/pkg/v2 v2.9.3-0.20200901032843-632995149920
	github.com/pulumi/pulumi/sdk/v2 v2.9.3-0.20200901032843-632995149920
	github.com/stretchr/testify v1.6.1
	gopkg.in/airbrake/gobrake.v2 v2.0.9 // indirect
	gopkg.in/gemnasium/logrus-airbrake-hook.v2 v2.1.2 // indirect
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
