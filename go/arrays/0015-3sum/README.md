# Problem: 3Sum (LeetCode 15)

From an integer array, we are asked to find triplets whose sum is 0.
We must not repeat the same triplet.

## How I Thought About It

At first, I said “let me try all triplets” (brute force) but this is **O(n³)** and slow in practice.
Then I realized this:

If the array is **sorted**, I can use **two pointers** moving from both ends.

In a sorted array:

- If the sum is **small**, I move the **left** pointer to the right to reach a larger number.
- If the sum is **big**, I move the **right** pointer to the left to reach a smaller number.
- I can easily skip duplicates from the same value (because after sorting they stand next to each other).

## Solution Logic

- **Sort** `nums`.
- I traversed **i** from left to right and in each round I said `nums[i]` should be the first element.
- If `i > 0` and `nums[i] == nums[i-1]`, it’s the **same start**, we **skip** it.
- I set up **two pointers**: `l = i + 1` and `r = n - 1`.
- I checked the sum `s = nums[i] + nums[l] + nums[r]`.
- If `s == 0`, I **added** the triplet and then did `l++`, `r--` and **skipped** the same `l/r` values.
- If `s < 0`, the sum is small → `l++`.
- If `s > 0`, the sum is big → `r--`.
- When all `i`s are done, we can return the **unique** results.

## Complexity

- **Time:** `O(n²)`
- **Extra Space:** `O(1)` (excluding the output)
