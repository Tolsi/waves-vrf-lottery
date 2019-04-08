rm -rf build/

env GOOS=darwin GOARCH=amd64 go build -o build/mac/pick-winner pick-winner.go
env GOOS=darwin GOARCH=amd64 go build -o build/mac/verify-winner verify-winner.go
env GOOS=darwin GOARCH=amd64 go build -o build/mac/create-keys create-keys.go
env GOOS=darwin GOARCH=amd64 go build -o build/mac/retweets-parser retweets-parser.go
env GOOS=darwin GOARCH=amd64 go build -o build/mac/outgoing-waves-tx-recipients outgoing-waves-tx-recipients.go
env GOOS=darwin GOARCH=amd64 go build -o build/mac/waves-height-tools waves-height-tools.go

env GOOS=linux GOARCH=386 go build -o build/linux/pick-winner pick-winner.go
env GOOS=linux GOARCH=386 go build -o build/linux/verify-winner verify-winner.go
env GOOS=linux GOARCH=386 go build -o build/linux/create-keys create-keys.go
env GOOS=linux GOARCH=386 go build -o build/linux/retweets-parser retweets-parser.go
env GOOS=linux GOARCH=386 go build -o build/linux/outgoing-waves-tx-recipients outgoing-waves-tx-recipients.go
env GOOS=linux GOARCH=386 go build -o build/linux/waves-height-tools waves-height-tools.go

env GOOS=windows GOARCH=386 go build -o build/windows/pick-winner pick-winner.go
env GOOS=windows GOARCH=386 go build -o build/windows/verify-winner verify-winner.go
env GOOS=windows GOARCH=386 go build -o build/windows/create-keys create-keys.go
env GOOS=windows GOARCH=386 go build -o build/windows/retweets-parser retweets-parser.go
env GOOS=windows GOARCH=386 go build -o build/windows/outgoing-waves-tx-recipients outgoing-waves-tx-recipients.go
env GOOS=windows GOARCH=386 go build -o build/windows/waves-height-tools waves-height-tools.go

zip -j build/release-mac-$(git describe --tags).zip build/mac/*
zip -j build/release-linux-$(git describe --tags).zip build/linux/*
zip -j build/release-windows-$(git describe --tags).zip build/windows/*