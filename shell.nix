with import <nixpkgs> {};

let

in stdenv.mkDerivation {
  name = "kindle-notes";
  buildInputs = [
    go
  ];
  src = null;
}
