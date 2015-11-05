### WTF?

This repository contains solution for printing some integer streams in order. Details below:


### Problem 

There are streams(could be unlimited in size). 
Each returns an integer which is greater than the last integer from its own stream and each stream is independent of the other stream.
Print the values in an increasing order from the all the streams combined. eg stream.read() // returns ant int. 
The next time read() is called, it will return an integer >= to the previous value.

Stream 1:
```
1, 3, 8, 10 ... 
```

Stream 2:
```
1, 4, 5, 6, 12 ... 
```

Stream 3: 
```
3, 7, 8, 9 ... 
```

Expected result:
```
1, 1, 3, 3, 4, 5, 6, 7, 8, 8, 9, 10, 12 ...
```

### Solution

```GO
    stream0 := stream.NewStream([]int{1, 3, 8, 10})
	stream1 := stream.NewStream([]int{1, 4, 5, 6, 12})
	stream2 := stream.NewStream([]int{3, 7, 8, 9})

	for val := range stream.Order(stream0, stream1, stream2) {
		fmt.Printf("%d, ", val)
	}
```

Or run in terminal:
```
go run main.go
```