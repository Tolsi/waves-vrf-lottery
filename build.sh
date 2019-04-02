rm -rf build/

env GOOS=darwin GOARCH=386 go build -o build/mac/vrf-prove pick-winner.go
env GOOS=darwin GOARCH=386 go build -o build/mac/vrf-verify verify-winner.go
env GOOS=darwin GOARCH=386 go build -o build/mac/create-keys create-keys.go

env GOOS=linux GOARCH=386 go build -o build/linux/vrf-prove pick-winner.go
env GOOS=linux GOARCH=386 go build -o build/linux/vrf-verify verify-winner.go
env GOOS=linux GOARCH=386 go build -o build/linux/create-keys create-keys.go

env GOOS=windows GOARCH=386 go build -o build/windows/vrf-prove pick-winner.go
env GOOS=windows GOARCH=386 go build -o build/windows/vrf-verify verify-winner.go
env GOOS=windows GOARCH=386 go build -o build/windows/create-keys create-keys.go

zip -j build/release-mac-$(git describe --tags).zip build/mac/*
zip -j build/release-linux-$(git describe --tags).zip build/linux/*
zip -j build/release-windows-$(git describe --tags).zip build/windows/*