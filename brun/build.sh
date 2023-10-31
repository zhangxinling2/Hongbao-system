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
  mv ${tname} ${TARGET_FILE_NAME}${EXT}
  env GOOS=$GOOS GOARCH=$GOARCH
  go build -o ${tname} -v ${SOURCE_FILE_NAME}.go
  chmod +x ${tname}
  if [ ${GOOS} == "windows" ]; then
    zip {tname}.zip ${TARGET_FILE_NAME}${EXT} config.ini *.sh ../public
  else
    tar -czvf {tname}.tar.gz ${TARGET_FILE_NAME}${EXT} config.ini *.sh ../public -C ./.
  fi
  mv ${TARGET_FILE_NAME}${EXT} ${tname}
}
CGO_ENABLED=0
GOOS=darwin
GOARCH=amd64
build