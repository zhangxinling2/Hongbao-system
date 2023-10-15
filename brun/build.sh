SOURCE_FILE_NAME=main
TARGET_FILE_NAME=reskd
cpath=`pwd`
PROJECT_PATH=${cpath%src*}
export GOPATH=$GOPATH:${PROJECT_PATH}
# shellcheck disable=SC1123
# shellcheck disable=SC1073
# shellcheck disable=SC1055
build(){
  echo $GOOS $GOARCH
  tname=${TARGET_FILE_NAME}_${GOOS}_${GOARCH}${EXT}
  env GOOS=$GOOS GOARCH=$GOARCH
  go build -o ${tname} -v ${SOURCE_FILE_NAME}.go
  chmod +x ${tname}
}
GOOS=darwin
GOARCH=amd64
build