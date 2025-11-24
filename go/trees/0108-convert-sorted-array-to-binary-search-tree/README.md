# Problem: Convert Sorted Array to Binary Search Tree (LeetCode 108)

We are given a sorted integer array `nums` (ascending order).  
**Goal:** Build a height-balanced **Binary Search Tree (BST)** from this array and return the root.

---

## What does height-balanced mean?

- Left and right subtree heights are as close as possible  
  (tree is not skewed to one side).
- Using the middle element as root helps keep the tree balanced.

---

## Idea (Recursive)

Since `nums` is sorted:

- Pick the **middle element** as root
- Left part of the array → becomes **left subtree**
- Right part of the array → becomes **right subtree**

We write a helper:

`build(nums, left, right)`

1. If `left > right` → no numbers left in this range → return `nil`
2. `mid := (left + right) / 2`
3. Create node with value `nums[mid]`
4. Recursively build:
   - `node.Left = build(nums, left, mid-1)`
   - `node.Right = build(nums, mid+1, right)`
5. Return `node`

Main function:

`sortedArrayToBST(nums)` just calls `build(nums, 0, len(nums)-1)`.

---

## Complexity

- **Time:** `O(n)` – each element is used once to create a node
- **Space:** `O(h)` – recursion depth (`h` = tree height)
