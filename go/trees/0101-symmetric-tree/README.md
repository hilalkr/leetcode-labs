# Problem: Symmetric Tree (LeetCode 101)

We are given the root of a binary tree.  
**Goal:** Check if the tree is a mirror of itself (symmetric around the center).

---

## What does symmetric mean?

- Left subtree and right subtree are **mirror images**
- Shape matches
- Values match in those mirrored positions

So when you look from the center, left side and right side should look the same but flipped.

---

## Idea (Recursive)

Write a helper:

`isMirror(left, right)`

1. If both nodes are `nil` → they match here → return `true`
2. If one is `nil` and the other is not → not symmetric → return `false`
3. If values are different → not symmetric → return `false`
4. Otherwise:
   - Check `isMirror(left.Left, right.Right)`
   - And `isMirror(left.Right, right.Left)`
   - Both must be `true`

---

## Complexity

- **Time:** `O(n)` – we visit each node once
- **Space:** `O(h)` – recursion depth (tree height)
