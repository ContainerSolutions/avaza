{ pkgs ? (import <unstable> {}) }:
with pkgs;
stdenv.mkDerivation rec {
  name = "avaza";
  buildInputs = [ go ];

  shellHook = ''
    export GOPATH=`pwd`/.gopath/
    export PATH=$GOPATH/bin:$PATH
    export GO111MODULE=on
  '';
}
