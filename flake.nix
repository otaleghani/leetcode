{
  description = "leetcode";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
  };

  outputs =
    { self, nixpkgs }:
    let
      system = "x86_64-linux";
      pkgs = nixpkgs.legacyPackages.${system};
    in
    {
      devShells.${system}.default = pkgs.mkShell {
        packages = with pkgs; [
          nodejs_24
          pnpm
          zsh
          typescript-language-server
          nodePackages.prettier
          go
          gopls
          mdformat
        ];
        shellHook = ''
          tmux new-session -s leetcode -d 'nvim'
          exec tmux attach-session -t leetcode
        '';
      };
    };
}
