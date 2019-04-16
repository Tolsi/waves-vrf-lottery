{ pkgs ? import <nixpkgs> {} }:

with pkgs;

buildGoPackage {
  name = "curve25519-go";
  src = ./.;
  goPackagePath = "github.com/Tolsi/vrf-lottery/curve";
}
