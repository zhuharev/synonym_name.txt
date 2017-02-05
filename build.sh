GOOS=windows GOARCH=386 go build -o build/synonym-win-32.exe ./src/cmd/synonym/*.go
GOOS=windows GOARCH=amd64 go build -o build/synonym-win-64.exe ./src/cmd/synonym/*.go
GOOS=darwin GOARCH=386 go build -o build/synonym-darwin-32 ./src/cmd/synonym/*.go
GOOS=darwin GOARCH=amd64 go build -o build/synonym-darwin-64 ./src/cmd/synonym/*.go
GOOS=linux GOARCH=386 go build -o build/synonym-lin-32 ./src/cmd/synonym/*.go
GOOS=linux GOARCH=amd64 go build -o build/synonym-lin-64 ./src/cmd/synonym/*.go

cd build
zip win32.zip  synonym-win-32.exe
cd .. 
zip build/win32.zip synonym_name.txt

cd build
zip win64.zip  synonym-win-64.exe
cd .. 
zip build/win64.zip synonym_name.txt

cd build
tar -zcvf lin64.tgz synonym-lin-64 ../synonym_name.txt
tar -zcvf lin32.tgz synonym-lin-32 ../synonym_name.txt
cd ..

cd build
zip darwin64.zip  synonym-darwin-64
cd .. 
zip build/darwin64.zip synonym_name.txt

cd build
zip darwin32.zip  synonym-darwin-32
cd .. 
zip build/darwin32.zip synonym_name.txt