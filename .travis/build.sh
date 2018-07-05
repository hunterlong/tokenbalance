#!/usr/bin/env bash
mkdir build

APP="tokenbalance"

rm -rf build

xgo -go 1.10.x --targets=darwin/amd64 --dest=build -ldflags="-X main.VERSION=$VERSION" ./
xgo -go 1.10.x --targets=darwin/386 --dest=build -ldflags="-X main.VERSION=$VERSION" ./
xgo -go 1.10.x --targets=linux/amd64 --dest=build -ldflags="-X main.VERSION=$VERSION" ./
xgo -go 1.10.x --targets=linux/386 --dest=build -ldflags="-X main.VERSION=$VERSION" ./
xgo -go 1.10.x --targets=windows-6.0/amd64 --dest=build -ldflags="-X main.VERSION=$VERSION" ./
xgo -go 1.10.x --targets=windows-6.0/386 --dest=build -ldflags="-X main.VERSION=$VERSION" ./
xgo -go 1.10.x --targets=linux/arm-7 --dest=build -ldflags="-X main.VERSION=$VERSION" ./
xgo -go 1.10.x --targets=linux/arm64 --dest=build -ldflags="-X main.VERSION=$VERSION" ./
xgo -go 1.10.x --targets=linux/amd64 --dest=build -ldflags="-X main.VERSION=$VERSION -linkmode external -extldflags -static" -out alpine ./

cd build
ls

sudo mv alpine-linux-amd64 $APP
sudo tar -czvf $APP-linux-alpine.tar.gz $APP && rm -f $APP

sudo mv $APP-darwin-10.6-amd64 $APP
sudo tar -czvf $APP-osx-x64.tar.gz $APP && rm -f $APP

sudo mv $APP-darwin-10.6-386 $APP
sudo tar -czvf $APP-osx-x32.tar.gz $APP && rm -f $APP

sudo mv $APP-linux-amd64 $APP
sudo tar -czvf $APP-linux-x64.tar.gz $APP && rm -f $APP

sudo mv $APP-linux-386 $APP
sudo tar -czvf $APP-linux-x32.tar.gz $APP && rm -f $APP

sudo mv $APP-windows-6.0-amd64.exe $APP.exe
sudo zip $APP-windows-x64.zip $APP.exe  && rm -f $APP.exe

sudo mv $APP-windows-6.0-386.exe $APP.exe
sudo zip $APP-windows-x32.zip $APP.exe  && rm -f $APP.exe

sudo mv $APP-linux-arm-7 $APP
sudo tar -czvf $APP-linux-arm7.tar.gz $APP && rm -f $APP

sudo mv $APP-linux-arm64 $APP
sudo tar -czvf $APP-linux-arm64.tar.gz $APP && rm -f $APP

#tar -czvf build/$APP-android-arm.tar.gz build/$APP-android-16-arm build/$APP-android-16-aar
#tar -czvf build/$APP-ios-arm.tar.gz build/$APP-ios-5.0-armv7 build/$APP-ios-5.0-framework

