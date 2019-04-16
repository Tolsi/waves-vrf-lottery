{ pkgs ? import <nixpkgs> {} }:

with pkgs;

buildGoPackage {
  name = "curve";
  src = ./.;
  goPackagePath = "github.com/Tolsi/vrf-lottery/curve";
}
