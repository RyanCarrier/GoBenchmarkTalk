# GoBenchmarkTalk

[Store Credit, Code Jam](https://code.google.com/codejam/contest/351101/dashboard)

## A-reason-to-benchmark

Contains an interesting example of go which made me wonder what the best way to get around it is.

## Benchmarking

Examples of different ways to solve the [Store Credit, Code Jam](https://code.google.com/codejam/contest/351101/dashboard) problem.

**NOTE: Be sure to `go run main.go` from CreateCases before trying to benchmark.**
### Method A

Use the difference of half the coin as the index of a map (with values of original position);

ie; For aim of 100, 10 would give 10, 75; 25 and 25;25.

This would cause 75 and 25 to collide, finding our sum to 100.
##### Method A2

Same as generic method A but using an array instead of map.
### Method B

Sort the array of coins, then work from the outside in until you find a match.

(Note: Due to the CreateCases being designed lazily, this method is ALWAYS worst case scenario)

### Method C

The obvious method, for each item, go through every item to see if they add to the aim.

##### Method C2

This method is similar, but takes out already compared items, so if the first loop is indexed by i, the second loop will start at i+1.

## CreateCases

Contains golang method to create the [Store Credit, Code Jam](https://code.google.com/codejam/contest/351101/dashboard) cases to be benchmarked against.

A cheat option can be toggled, where the cheat will place the correct coins at the very end of the list, forcing worst case scenario on every solving method.
