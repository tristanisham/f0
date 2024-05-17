# f0

## An implementation of the CVM algorithm for the distinct elements (f0) problem.

This package implements the [CVM](https://arxiv.org/abs/2301.10191) algoritm in
pure Go. CVM is an algorithm for the distinct elements problem. When you're
trying to find the unique elements in a slice, but the slice exceeds the amount
of available memory, use CVM to estimate the number of unique elements in said
slice.

```go
f0.Estimate(mySlice, 0.01, 0.001) // 329.4221234
```

### More Reading

- [Computer Scientists Invent an Efficient New Way to Count](https://www.quantamagazine.org/computer-scientists-invent-an-efficient-new-way-to-count-20240516/)

## Contributing

I **NEED** tests. The problem with this algorithm is that I need large sets of
comparable elements to test CVM on. Currently (below v1), there is no guarentee
the algorithm will work correctly. After we've developed a large and robust test
suite and can establish the algorithm is behaving as expected, this library will
be promoted to v1 and maintained.

The Bible is included for testing purposes. There's no copyright in the Kingdom of Heaven. 
