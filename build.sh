rm -rf build/

env GOOS=darwin GOARCH=386 go build -o build/mac/vrf-prove vrf-prove.go
env GOOS=darwin GOARCH=386 go build -o build/mac/vrf-verify vrf-verify.go
env GOOS=darwin GOARCH=386 go build -o build/mac/retweets-parser retweets-parser.go

env GOOS=linux GOARCH=386 go build -o build/linux/vrf-prove vrf-prove.go
env GOOS=linux GOARCH=386 go build -o build/linux/vrf-verify vrf-verify.go
env GOOS=linux GOARCH=386 go build -o build/linux/retweets-parser retweets-parser.go

env GOOS=windows GOARCH=386 go build -o build/windows/vrf-prove.exe vrf-prove.go
env GOOS=windows GOARCH=386 go build -o build/windows/vrf-verify.exe vrf-verify.go
env GOOS=windows GOARCH=386 go build -o build/windows/retweets-parser.exe retweets-parser.go

zip -j build/release-mac-$(git describe --tags).zip build/mac/* scripts/unix/* scripts/mac/*
zip -j build/release-linux-$(git describe --tags).zip build/linux/* scripts/unix/* scripts/linux/*
zip -j build/release-windows-$(git describe --tags).zip build/windows/*