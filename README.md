# Sum to N

This project demonstrates three different methods for calculating the sum of integers from 1 to a given number `n`. The three approaches include a naive iterative method, a mathematical formula, and a divide-and-conquer recursive method.

## Methods

### 1. **Naive For Loop Method (O(n) time complexity)**

This method uses a simple for loop to iterate through every number from 1 to `n` and adds each number to a cumulative sum. The time complexity of this method is linear (O(n)) because the loop runs `n` times.

---

### 2. **Mathematical Formula Method (O(1) time complexity)**

This method utilizes the mathematical formula for the sum of the first `n` integers:
\[
S_n = \frac{n(n + 1)}{2}
\]
This formula provides the sum in constant time (O(1)), making it the most efficient approach.

---

### 3. **Divide-and-Conquer Recursive Method (O(n) time complexity)**

This method uses recursion to break down the sequence into smaller parts and calculates the sum of each half separately. Even though it divides the problem, it still visits each number, and its time complexity remains O(n).

---

## How to Run the Code

To run the Go program:

1. **Install Go**: Make sure you have Go installed on your machine. If you don't have it, download and install it from the [official Go website](https://golang.org/dl/).

2. **Clone the repository**: If you haven’t already, clone the repository to your local machine:

   ```bash
   git clone https://github.com/tannerlie/sum-to-n.git
   ```

3. **Navigate to the project folder**:

   ```
   cd sum-to-n
   ```

4. **Run the program**: Execute the following command to run the Go program and calculate the sum using the three methods:

   ```
   go run sumToN.go <value-of-n>
   ```

This will output the sum using each of the three methods.