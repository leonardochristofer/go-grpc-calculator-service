# Local Command
Use the following `make` command to operate locally
1. `make init`           # Init go module
2. `make reinit`         # Re-init go module on unix
3. `make reinits`        # Re-init go module on windows
4. `make tidy`           # Tidy up go module on unix
5. `make tidys`          # Tidy up go module on windows
6. `make proto-gen`      # Generate protobuf
7. `make run-grpc`       # Running server on unix
8. `make runs-grpc`      # Running server on windows
9. `make run`            # Running client on unix
10. `make runs`          # Running client on windows

# Prerequisite
This service uses Go and gRPC. I assume that you have everything installed on your local machine.

# Steps
1. Make sure you have `make` installed on your unix or windows to do the shortcut `make` command.
2. Change directory to the current service, init go module, tidy go module, make proto-gen.
3. Open 2 terminals.
4. `go run server/main.go --mode grpc` or `make run-grpc` for unix or `make runs-grpc` for windows to run server.
5. `go run client/main.go` or `make run` for unix or `make runs` for windows to run client.
6. Hit API on `localhost:9100/api/prime/<index>` to find nth prime number based on index.
Refer to https://en.wikipedia.org/wiki/List_of_prime_numbers
7. Hit API on `localhost:9100/api/prime/palindrome/<index>` to find nth prime palindrome number based on index.
Refer to https://en.wikipedia.org/wiki/Palindromic_prime