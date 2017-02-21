# GO File Stat Collect 

A Go program to troll a file system for file stats and send it to a MongoDB document store. 

## Benchmark 

An example test on my home directory with print verbosity turned off. 

```
99250 items found.
go run walk.go ~/  0.57s user 0.22s system 114% cpu 0.686 total
```
And with verbosity turned on. 

```
time go run walk.go ~/ 
[majority of output omitted]
99282 items found.
go run walk.go ~/  1.81s user 2.30s system 69% cpu 5.957 total
```
