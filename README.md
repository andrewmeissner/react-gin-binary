# Gin-React-Binary

This is some scaffolding of how to package your react app inside a go binary

Steps to reproduce:
```bash
touch main.go

# edit main.go

npx create-react-app ui
cd ui

#  add "homepage": "http://localhost/ui" to package.json

npm run build # or yarn build
cd ..
go-bindata-assetfs -o ui/bindata.go -pkg ui ui/build/...
go build -o singleBinary .
./singleBinary
```
