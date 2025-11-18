# Problem: 1-bit and 2-bit Characters (LeetCode 717)

We have an array like this: `[]int{0,1,0,0,...}` and it always ends with `0`.  
`0` → a single **1-bit character**  
`10` or `11` → together a **2-bit character**

The array is basically all the bits of the characters written one after another.  
What we want to answer is: when we read this array from left to right using these rules,  
**is the last character always a single-bit `0`?**

---

## My Thought Process and Solution Logic

We use an index `i`.

If `bits[i] == 0`:

- This is a 1-bit character → `i += 1`

If `bits[i] == 1`:

- This is the start of a 2-bit character → `bits[i]` and `bits[i+1]` form one character
- So we do `i += 2`

The array already ends withh `0`.  
We want to know if this last `0` is alone or part of a 2-bit character.

To do that, we loop while `i < n-1`.  
So we move the pointer forward **without stepping onto the last element**.

---

## Finally

When the loop finishes:

- If `i == n-1` → the pointer stopped on the last element  
  → the last `0` is a standalone character → return **`true`**

- If `i == n` → the pointer went past the array  
  → the last `0` was already used as part of a 2-bit character → return **`false`**

---

## Time and Space Complexity

- Time: **O(n)** — we scan the array once
- Extra space: **O(1)** — we only use a few variables
