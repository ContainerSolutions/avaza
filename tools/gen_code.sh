# Version of go-swagger to use.
version=v0.18.0

TOOLS_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
DIR=$TOOLS_DIR/..

mkdir -p $DIR/client
curl https://api.avaza.com/swagger/docs/v1 > $DIR/client/avaza-swagger.json

SWAGGER=$TOOLS_DIR/swagger-${version}

if [ ! -f $SWAGGER ]; then
  curl -o $SWAGGER -L'#' https://github.com/go-swagger/go-swagger/releases/download/$version/swagger_$(echo `uname`|tr '[:upper:]' '[:lower:]')_amd64
  chmod +x $SWAGGER
fi

set -e

(cd $DIR; $SWAGGER generate client --spec=client/avaza-swagger.json --target=client)
