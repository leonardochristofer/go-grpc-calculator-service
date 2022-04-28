### Local command
Use the following `make` command to operate locally
make init           # Init go module
```

```
make reinit         # Re-init go module
```

```
make tidy           # Tidy up go module
```

```
make proto-gen      # Generate protobuf
```

```
### Running Gateway and Service
This service uses Go and gRPC, so I assume that you have Go, gRPC and protobuf generator installed.
```

```
Steps:
1. Make sure you are using unix or WSL 2 and have `make` installed to do the shortcut `make` command. However, if you are using windows, refer to Makefile and just copy-paste the command directly to your command prompt.
2. Change directory to the current service, type make init, make tidy, make proto-gen.
3. Open 2 terminals.
4. `go run gateway.go` for running gateway.
5. `go run service.go` for running service.
6. Hit API on `localhost:9100/api/prime/<index>` to find nth prime number based on index.
Refer to https://en.wikipedia.org/wiki/List_of_prime_numbers
7. Hit API on `localhost:9100/api/prime/palindrome/<index>` to find nth prime palindrome number based on index.
Refer to https://en.wikipedia.org/wiki/Palindromic_prime