env GOOS=darwin GOARCH=386 go build -o build/mac/vrf-proof vrf-proof.go
env GOOS=darwin GOARCH=386 go build -o build/mac/vrf-verify vrf-verify.go
env GOOS=darwin GOARCH=386 go build -o build/mac/retweets-parser retweets-parser.go

env GOOS=linux GOARCH=386 go build -o build/linux/vrf-proof vrf-proof.go
env GOOS=linux GOARCH=386 go build -o build/linux/vrf-verify vrf-verify.go
env GOOS=linux GOARCH=386 go build -o build/linux/retweets-parser retweets-parser.go

env GOOS=windows GOARCH=386 go build -o build/windows/vrf-proof.exe vrf-proof.go
env GOOS=windows GOARCH=386 go build -o build/windows/vrf-verify.exe vrf-verify.go
env GOOS=windows GOARCH=386 go build -o build/windows/retweets-parser.exe retweets-parser.go

zip build/release-mac-$(git describe --tags).zip build/mac/*
zip build/release-linux-$(git describe --tags).zip build/linux/*
zip build/release-windows-$(git describe --tags).zip build/windows/*