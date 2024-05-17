# CVM
## An implementation of the CVM algorithm for the distinct elements problem.

This package implements the [CVM](https://arxiv.org/abs/2301.10191) algoritm in pure Go. CVM is an algorithm for the distinct elements problem.
When you're trying to find the unique elements in a slice, but the slice exceeds the amount of available memory, use CVM to estimate the number of unique elements in
said slice. 

```go
cvm.Count(s, 100) // 1500
```

### More Reading
- [Computer Scientists Invent an Efficient New Way to Count](https://www.quantamagazine.org/computer-scientists-invent-an-efficient-new-way-to-count-20240516/)

## Contributing

I **NEED** tests. The problem with this algorithm is that I need large sets of comparable elements to test CVM on. 