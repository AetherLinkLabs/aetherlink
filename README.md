# aetherlink

1. Call `registerName` to bind the name with the user address
```
./dist/aether contract register --name wenliang
```

2. Call `registerFileCID` to bind the file name and CID `baxxxxx`
```
./dist/aether contract bind --cid baxxxxx --name hello.html
```

3. Call `getFileCID` to query the file CID by specifying the file name
```
./dist/aether contract resolve --name hello.html --address 0xXXXXXXXX
```

4. Upload the file
```
./dist/aether upload --path ./resources/hello.html --filename hello.html
```