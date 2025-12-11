{
	description = "Label utilities";

	inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-25.11";
		flake-utils.url = "github:numtide/flake-utils";
	};

	outputs = { self, nixpkgs, flake-utils, ... }@inputs: flake-utils.lib.eachDefaultSystem (system:
		let
			pkgs = import nixpkgs { inherit system; };
		in
		{
			packages.textlabel = pkgs.buildGoModule {
				pname = "textlabel";
				version = "0.1.0";

				src = ./.;

				vendorHash = "sha256-ejc+3Gcp1Ax8z6CExigGgl6aXnyOiFSmA+eKWhxLDb0=";
				subPackages = [ "./textlabel" ];
			};

			packages.printserver = pkgs.buildGoModule {
				pname = "printserver";
				version = "0.1.0";

				src = ./.;

				vendorHash = "sha256-ejc+3Gcp1Ax8z6CExigGgl6aXnyOiFSmA+eKWhxLDb0=";
				subPackages = [ "./printserver" ];
			};
		}
	);
}
