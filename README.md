# go-star-sim
Star system simulation in Go

## Build

```
go install ./cmd/star-calc/
go install ./cmd/star-sim/
```
or
```
go build ./cmd/star-calc/
go build ./cmd/star-sim/
```

## Run

Example:
```
./star-calc -stars 1000 -steps 500 -file /tmp/simdata
./star-sim -file /tmp/simdata
```
