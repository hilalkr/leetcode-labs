# Problem: Binary Tree Inorder Traversal (LeetCode 94)

In this problem, we are asked to traverse a **binary tree** in **inorder** order.  
**Inorder** means:  
**Left → Node → Right**

So for each node, we first visit the **left** subtree,  
then the **node itself**,  
and finally the **right** subtree.

---

## Thought Process

At first, I wasn’t sure how to “walk” inside a tree.  
Then I realized that inorder traversal simply follows the same rule for every node:  
**first left, then the node, then right**.

We can do this in two different ways:

- **Recursive (easier to understand)** – we let the computer handle the function calls
- **Iterative (using a stack)** – we do the same process manually

---

## 1. Recursive Solution (simpler)

**Idea:**  
For each node, do the same three steps:  
go left → record the node → go right.

- **Time Complexity:** O(n)
- **Space Complexity:** O(h) (h = height of the tree, because of recursion calls)

---

## 2. Iterative Solution (using a stack)

**Idea:**  
Go all the way to the left, pushing each node to a stack.  
When there’s no more left node, pop one node from the stack (go back),  
record its value, and then move to the right.

- **Time Complexity:** O(n)
- **Space Complexity:** O(h) (the stack keeps track of the path)
