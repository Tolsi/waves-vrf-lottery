env GOOS=darwin GOARCH=386 go build -o build/mac/vrf-create vrf-create.go
env GOOS=darwin GOARCH=386 go build -o build/mac/vrf-verify vrf-verify.go
env GOOS=darwin GOARCH=386 go build -o build/mac/retweets-parser retweets-parser.go

env GOOS=linux GOARCH=386 go build -o build/linux/vrf-create vrf-create.go
env GOOS=linux GOARCH=386 go build -o build/linux/vrf-verify vrf-verify.go
env GOOS=linux GOARCH=386 go build -o build/linux/retweets-parser retweets-parser.go

env GOOS=windows GOARCH=386 go build -o build/windows/vrf-create.exe vrf-create.go
env GOOS=windows GOARCH=386 go build -o build/windows/vrf-verify.exe vrf-verify.go
env GOOS=windows GOARCH=386 go build -o build/windows/retweets-parser.exe retweets-parser.go

zip release-mac-$(git describe --tags).zip build/mac/*
zip release-linux-$(git describe --tags).zip build/linux/*
zip release-windows-$(git describe --tags).zip build/windows/*