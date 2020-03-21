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
./star-calc -numbStars 1000 -numbSteps 500 -outputFile /tmp/output
./star-sim -inputFile /tmp/output
```
