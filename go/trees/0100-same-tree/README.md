# Problem: Same Tree (LeetCode 100)

We have two binary trees: **p** and **q**.  
**Goal:** Decide if the two trees are exactly the same.

## What does same mean?

- **Same shape** (nodes exist or don’t exist at the same positions)
- **Same values** at the same positions

Think of walking the two trees side by side and comparing the node at each position.

## How to check

1. If both nodes are `nil`: they match at this position (continue)
2. If one is `nil` and the other is not: different → return `false`
3. If both exist but their values differ: different → return `false`
4. If both exist and values are equal: compare their left children, then their right children

## Solution 1: Recursive

**Idea:**  
Compare the nodes at the same position. If they match, recursively compare left children and right children.

- Time: O(n) (n is total number of nodes)
- Extra space: O(h) (h is tree height due to recursion depth)

## Solution 2: Iterative (queue or stack)

**Idea:**  
Keep **pairs** of nodes (one from each tree) in a queue or stack.  
Process one pair at a time using the rules above.  
If they match, push the pair of left children and the pair of right children.

- Time: O(n)
- Extra space: O(w) (w is the maximum number of nodes at any level)
