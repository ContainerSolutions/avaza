# Avaza CLI tool

Using Avaza? With this tool you can update your hour registration, and receipts from the CLI.

# Development notes

* Assumes >= Go 1.11
* I use Nix to define the build environment; just run `nix-shell`.
* Uses go-swagger to generate the client from the Swagger/Open-API spec of Avaza. Please run `./tools/gen-code.sh` to both download the swagger spec and generate the client from it.
