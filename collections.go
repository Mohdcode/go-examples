//Collection,Fixed Size?,Reference Semantics?,Use Case
// Array,Yes,No (Copies),Low-level hardware/fixed buffers.
// Slice,No,Yes (Header),95% of all Go code.
// Map,No,Yes,"Quick lookups, caches, sets."

// Slice Header (The "Trio")
// A slice is not the data itself; it is a struct that describes a piece of an underlying array. When you pass a slice to a function, you are actually passing a copy of this 24-byte header (on a 64-bit system).

// The Internal Structure:

// Go
// type slice struct {
//     Data unsafe.Pointer // 8 bytes: Address of the first element
//     Len  int            // 8 bytes: Current number of elements
//     Cap  int            // 8 bytes: Max elements before reallocating
// }
// Data Pointer: This points specifically to the index where the slice starts. If you do b := a[2:], the pointer in b just starts 16 bytes (2 ints) further down the same array.

// Length: What you see when you call len(). It’s the "active" zone.

// Capacity: What you see with cap(). It’s the "reserved" zone.

// Why this matters for Performance:

// When you pass a slice to a function, Go copies these 24 bytes. It does not copy the underlying array. This is why Go can handle millions of items efficiently—it’s just passing small "windows" around.

// 2. The String Header (The "Duo")
// Strings in Go are even simpler. They are immutable, so they don't need a capacity. They are just a pointer and a length.

// The Internal Structure:

// Go
// type string struct {
//     Data unsafe.Pointer // 8 bytes: Address of the bytes
//     Len  int            // 8 bytes: Number of bytes
// }
// The SDE Trap: Since strings and slices of bytes ([]byte) have very similar headers, you can convert between them. However, standard conversion string(myBytes) creates a full copy of the data to ensure immutability. In high-performance backend code (like parsing massive JSONs), engineers often use "unsafe" tricks to switch headers without copying the data.

// 3. The Map Header (The "Hmap")
// Maps are the most complex. A map variable is actually a pointer to a struct called hmap.

// The Internal Structure (Simplified):

// Go
// type hmap struct {
//     count     int    // Number of elements
//     buckets   unsafe.Pointer // Array of 2^B buckets
//     oldbuckets unsafe.Pointer // Used during growing/evacuation
//     // ... other internal flags
// }
// Buckets: Go maps use an array of buckets. Each bucket holds up to 8 key-value pairs.

// Hashing: Go hashes your key, uses the low bits to pick a bucket, and high bits to find the specific slot inside that bucket.

// Evacuation: When a map gets too full, Go allocates a new bucket array (double the size) and gradually moves data over. This "incremental" move prevents the program from freezing up during a large resize.

// 4. How append manipulates the Header
// When you call slice = append(slice, value), Go performs a "check-and-update" logic on the header:

// Check: Is Len + 1 <= Cap?

// If Yes: Just increment the Len in the header and write the data to the pointer + offset. (Very Fast)

// If No:

// Allocate a new, larger array (usually 2x bigger).

// Copy the old data to the new array.

// Update the Pointer to the new address.

// Update Len and Cap. (Slower - triggers GC work later)

// 5. Summary: Headers in your "Mental Model"
// Collection	Header Size	Contains	Semantics
// Array	N/A	The actual data	Value (full copy)
// Slice	24 bytes	Pointer, Len, Cap	Reference-like (header copy)
// String	16 bytes	Pointer, Len	Value (immutable)
// Map	8 bytes	Pointer to hmap	Pure Reference

package main

import (
	"fmt"
)

func collection() {

	arr := [3]int{1, 2, 3} //initializing on the go
	// var newArr [3]int
	// newArr[0]=1
	fmt.Println("from the arrays")
	for _, v := range arr {
		fmt.Println(v)
	}
	// slices := []int{1, 2, 3, 4, 5, 6, 7}
	slice1 := make([]int, 0, 10) //type ,len,cap
	//need copy function to really copy the slice otherwise it would be just sharing the memory address with variables

	slice1 = append(slice1, 34)
	// 	"slices"
	// "maps" use these inbuilt function by go to use its feature

	//SLICING EXPRESSION USAGE
	// 	nums[2:] (Missing high): Starts at index 2 and goes all the way to the end.

	// nums[:3] (Missing low): Starts at the very beginning (index 0) and stops before index 3.

	// nums[:] (Missing both): Points to the entire array.

	fmt.Println("from the slices")
	for _, v := range slice1 {
		fmt.Println(v)

	}

	maps1 := make(map[int]string)

	// m := map[string]int{"Alice": 25, "Bob": 30}//direct initializing

	fmt.Println("from the maps")
	maps1[1] = "one"
	maps1[2] = "two"
	delete(maps1, 1)

	for i, v := range maps1 { //copying the data in range
		//if map and slice is empty loop never enters
		fmt.Println(i, v)
	}

	// 	State,Variable,Header Len / count,Does it Loop?
	// Nil,var s []int,0,No
	// Empty,s := []int{},0,No

	// 	A Nil slice has a Data pointer of nil.

	// An Empty slice (created with {} or make) has a Data pointer to a special "zerobase" memory address, but the Len is still 0.

}

// Function,Slices,Maps,What it touches in the Header
// make,Yes,Yes,Initializes all fields/buckets.
// len,Yes,Yes,Reads Len or count.
// cap,Yes,No,Reads Cap.
// append,Yes,No,"Updates Len, possibly Data & Cap."
// delete,No,Yes,Updates count and bucket data.
// copy,Yes,No,Reads Data address and Len.
// clear,Yes,Yes,Resets entries/elements to zero.

// new(T): Allocates zeroed memory and returns a pointer (*T). If you use new([]int), you get a pointer to a nil header. You can't append to it easily.

// make(T, ...): Only for slices, maps, and channels. It returns an initialized (non-nil) header.

// Rule of Thumb: Always use make for slices and maps.
