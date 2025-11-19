# Problem: Keep Multiplying Found Values by Two (LeetCode 2154)

In this problem, we are given an array **nums** and an integer **original**.  
The task is:

- If **original** exists in the array → multiply it by **2**
- Then search for the new value again in the array
- Continue this process as long as the value is found
- Once the value is no longer found return it

---

## Thought Process

At first, I thought about checking the array every time to see if  
`original` exists. But this would be inefficient because we might scan  
the whole array multiple times.

Then I realized a simpler approach:

If I store all elements from the array in a **set** (`map[int]bool` in Go),  
I can check existence in **O(1)** time.

---

## Solution Logic

1. Insert all elements of `nums` into a map to act as a set.  
   Fast existence checks.

2. While `original` exists in the set:  
   → `original *= 2`

3. When `original` is no longer found, stop and return it.

---

## Complexity

**Time Complexity:** O(n)  
 We loop through the array once to build the set, and all lookups are O(1).

**Space Complexity:** O(n)  
 The set stores all values from the array.
