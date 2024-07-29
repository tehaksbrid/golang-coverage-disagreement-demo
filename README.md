# golang-coverage-disagreement-demo
Demonstrates a disagreement between official coverage analysis tools for the Golang team

To observe this inconsistency:
1. To see this inconsistency, run `make coverage-browser`. 
2. Observe that stdout shows 75% coverage for package `foo`.
3. Observe that the HTML coverage browser shows 100% coverage for all files within package `foo`

To see how this has changed since 1.21.12:
1. Switch Go version to 1.21.12 in `go.mod`
2. Ensure your machine is running this version
3. Repeat above steps
4. Observe that both methods are in agreement that package `foo` is 100% covered