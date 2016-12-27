A slice is a segment of an array.  Like arrays, slices have an index and a length. 
Unlike arrays, the length is allowed to change.
(by dynmically updating the underlying array to accomodate the length of the slice)

slices are always associated with an array.  They can never be longer than their
underlying array but they can be smaller.

Array with a capcity of 10
_____________________
| | | | | | | | | | | 
 ^-------^
     ^
 Slice with a capacity of 5

 Append
 we can append an element on to the end of a slice, assuming the underlying array has enough space left.
 If the underlying array is not big enough, the array is destroyed, a new array is created and all the 
 data is copied to the new array, then the additional element(s) is(are) appended to the end of the new array 

 Copy
 Copy takes two arguments destination (dst) and source (src).  Every entry is src is copied to dst 
 OVERWRITING whatever is there.  If the length of the slices are not the same, the smaller of the two
 is used. 


Usage:

Typically, slices are used to represent lists of items.