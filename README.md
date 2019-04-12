# Gin-React-Binary

This is some scaffolding of how to package your react app inside a go binary

Steps to reproduce:
```bash
touch main.go

# edit main.go

npx create-react-app ui
cd ui
npm run build
cd ..
go-bindata-assetfs ui/build/...
go build -o singleBinary .
```

this uses contrib/static and implements a `BinaryFileSystem` to serve up the binary represented artifacts.