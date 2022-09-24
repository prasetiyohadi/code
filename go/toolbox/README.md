# Toolbox

Command line tool using cobra and viper library.

## Development

```
# Install cobra-cli
go install github.com/spf13/cobra-cli@latest

# Create new project
cobra-cli --help
cobra-cli init toolbox
cd toolbox

# Create with viper support
cobra-cli init toolbox --viper
cd toolbox

# Add subcommands
cobra-cli add --help
cobra-cli add ping
cobra-cli add net

# Test
go run main.go net
go run main.go net ping --url google.com
go build .
./toolbox -h
./toolbox net -h
./toolbox net ping -h

# Add another subcommands
cobra-cli add info
cobra-cli add disk-usage

# Get third party library
go get github.com/ricochet2200/go-disk-usage/du

# Test
go mod tidy
go run main.go
go run main.go info
go run main.go info -h
go run main.go info diskUsage -h
go run main.go info diskUsage
go run main.go net -h
go run main.go net ping -h
go run main.go net ping --url google.com
```
