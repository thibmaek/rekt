{
  description = "rekt";
  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  outputs = { self, nixpkgs }: {
    devShells.x86_64-linux.default =
      let pkgs = nixpkgs.legacyPackages.x86_64-linux;
      in pkgs.mkShell {
        packages = [ pkgs.go_1_25 ];
      };
    devShells.aarch64-darwin.default =
      let pkgs = nixpkgs.legacyPackages.aarch64-darwin;
      in pkgs.mkShell {
        packages = [ pkgs.go_1_25 ];
      };
  };
}
