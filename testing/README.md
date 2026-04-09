`go test` only executes test in current folder/package

to start all tests of a module, run
```shell
go test ./...
```
from the root folder.


There are various frameworks for mocking out there (e.g. [gomock](https://github.com/uber-go/mock)).
Simple mock implementations can suffice though.
