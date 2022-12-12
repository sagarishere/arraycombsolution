# Question

AND is a standard bitwise operation. For example, given K = 12 (binary representation 01100)
and L = 21 (binary representation 10101) we obtain:
01100 AND
10101 =
00100
The AND operation can be extended to N integers, for example:
01100 AND 10101 AND 00100 = 00100
Because AND of 01100 (first argument) and 10101 (second argument) is 00100,
and AND of this number with 00100 (third argument) is also 00100.

---

## Write a function

func Solution (A []int) int
that, given an array A consisting of N integers, returns the size of the largest possible subset of A
such that AND of all its elements is greater than 0.

---

### Examples

1. Given A = [13, 7, 2, 8, 3] your function should return 3.
   We can pick subset 13, 7 and 3. AND of these elements is positive and it is the largest possible subset of numbers that fulfills the criteria.
   1101 (13) AND
   0111 (7) AND
   0011 (3) =
   0001 (1)
2. Given A = [1, 2, 4, 8] your function should return 1. The AND of each pair from the array is equal to 0.
3. Given A = [16, 16] your function should return 2. The AND of both numbers is 16.
   Write an efficient algorithm for the following assumptions:
   • N is an integer within the range [1..100,000];
   each element of array A is an integer within the range [1..1,000,000,000].

---

## Solution

The solution contained in the project has been originally developed and is copyright of [Sagar Yadav](https://www.linkedin.com/in/sagaryadav/)

The solution shall not be used for direct or indirect commercial purposes without the written consent of the author.
