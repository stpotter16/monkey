shell:
	nix develop -c $$SHELL

test:
	go test ./...
