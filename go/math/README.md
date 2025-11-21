# Problem: Excel Sheet Column Title (LeetCode 168)

In this problem, we are given a **positive integer**, and we must convert it into its corresponding **Excel column title**.

---

## Idea

Excel columns behave like a **base-26 number system**,  
but with one important difference:

- **A = 1, B = 2, …, Z = 26**

Because of this, the conversion follows these steps:

1. **Decrease the number by 1** to shift it into the range 0–25
2. Compute **number % 26** to determine the current letter
3. Append this letter to the result
4. Divide the number by **26** and continue
5. Since characters are collected in reverse order, **reverse the final string**

---

## Time and Space Complexity

- **Time Complexity:**  
  **O(log₍₂₆₎ n)** — the number is repeatedly divided by 26

- **Space Complexity:**  
  **O(k)** — where _k_ is the length of the resulting column title
