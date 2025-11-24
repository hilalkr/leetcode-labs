# Problem: Maximum Depth of Binary Tree (LeetCode 104)

We are given the root of a binary tree.  
**Goal:** Return the _maximum depth_ of the tree.

---

## What does depth mean?

- Depth = number of nodes on the **longest path**  
  from the root down to a leaf.
- Empty tree → depth `0`.

Example: `3 -> 20 -> 15` has depth `3` (3 nodes in the path).

---

## Idea (Recursive)

We write a function:

`maxDepth(root)`

1. If `root == nil` → tree is empty → return `0`
2. Otherwise:
   - Find depth of left subtree: `left := maxDepth(root.Left)`
   - Find depth of right subtree: `right := maxDepth(root.Right)`
   - Current node depth = `1 + max(left, right)`

The `+1` is for the current node itself.

---

## Complexity

- **Time:** `O(n)` – visit each node once
- **Space:** `O(h)` – recursion call stack (`h` = tree height)
