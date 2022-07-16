# Random Images List

## Usage
```markdown
Make sure that the `image` folder is created, 
and that it holds the pictures.
```

## Building
```shell
git clone https://github.com/Sakura1943/RandomImages_GoLang.git ./random_images
cd ./random_images
go build -o random_images ./main.go
chmod +x ./random_images
```
or
```shell
wget https://github.com/Sakura1943/RandomImages_GoLang/releases/download/Packages/random_images.tar.gz
mkdir -p ./random_images
tar -zxf ./random_images.tar.gz -C ./random_images
cd ./random_images
chmod +x ./random_images
```

## Running
```shell
./random_images
```