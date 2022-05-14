#!/bin/bash

NAME="NotenrechnerWindows-x86"

cat > "${NAME}.rc" << EOL
id ICON "./Img/icon.ico"
GLFW_ICON ICON "./Img/icon.ico"
EOL

x86_64-w64-mingw32-windres "${NAME}.rc" -O coff -o "${NAME}.syso"
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ HOST=x86_64-w64-mingw32 \
    go build -ldflags "-s -w -H=windowsgui -extldflags=-static" -p 4 -v -o "${NAME}.exe" src/Notenrechner.go

rm "${NAME}.rc" "${NAME}.syso"
