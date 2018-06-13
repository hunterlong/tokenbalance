#!/usr/bin/env bash
mkdir build

APP="tokenbalance"

#xgo --targets=*/* --dest=build -ldflags="-X main.VERSION=$VERSION" ./

xgo --targets=darwin/amd64 --dest=build -ldflags="-X main.VERSION=$VERSION" ./
xgo --targets=darwin/386 --dest=build -ldflags="-X main.VERSION=$VERSION" ./

xgo --targets=linux/amd64 --dest=build -ldflags="-X main.VERSION=$VERSION" ./
xgo --targets=linux/386 --dest=build -ldflags="-X main.VERSION=$VERSION" ./

xgo --targets=windows-6.0/amd64 --dest=build -ldflags="-X main.VERSION=$VERSION" ./
xgo --targets=windows-6.0/386 --dest=build -ldflags="-X main.VERSION=$VERSION" ./

xgo --targets=linux/arm-7 --dest=build -ldflags="-X main.VERSION=$VERSION" ./
xgo --targets=linux/arm64 --dest=build -ldflags="-X main.VERSION=$VERSION" ./

#xgo --targets=android-16/arm --dest=build -ldflags="-X main.VERSION=$VERSION" ./
#xgo --targets=ios/arm-7 --dest=build -ldflags="-X main.VERSION=$VERSION" ./

mv build/$APP-darwin-10.6-amd64 build/$APP-osx-x64
mv build/$APP-darwin-10.6-386 build/$APP-osx-x32
mv build/$APP-linux-amd64 build/$APP-linux-x64
mv build/$APP-linux-386 build/$APP-linux-x32
mv build/$APP-windows-6.0-amd64.exe build/$APP-windows-x64.exe
mv build/$APP-windows-6.0-386.exe build/$APP-windows-x32.exe
mv build/$APP-linux-arm-7 build/$APP-linux-arm7

#tar -czvf build/$APP-android-arm.tar.gz build/$APP-android-16-arm build/$APP-android-16-aar
#tar -czvf build/$APP-ios-arm.tar.gz build/$APP-ios-5.0-armv7 build/$APP-ios-5.0-framework

