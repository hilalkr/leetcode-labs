# Problem: Merge Sorted Array (LeetCode 88)

In this problem, we need to merge two sorted arrays into one sorted array.  
The array `nums1` has extra empty spaces (these are zeros) at the end, so the merged result should be placed directly inside `nums1`.

## My Thought Process

At first, I thought about merging from the beginning,  
but I realized that I would have to shift elements inside `nums1`, which would make it messy.  
Then I understood that starting from the end makes much more sense.

If I place the largest element at the end of `nums1`,  
I can use the empty space and merge the arrays in-place without moving everything around.

## Solution Logic

- `i` → last real element in `nums1` (`m - 1`)
- `j` → last element in `nums2` (`n - 1`)
- `k` → last position in `nums1` (`m + n - 1`)

At each step, I compare `nums1[i]` and `nums2[j]`.  
Whichever is larger goes into `nums1[k]`.  
Then I move the pointer backward.  
I keep doing this until all elements from `nums2` are placed.

## Time and Space Complexity

- Time: O(m + n)
- Space: O(1) (no extra array used)
