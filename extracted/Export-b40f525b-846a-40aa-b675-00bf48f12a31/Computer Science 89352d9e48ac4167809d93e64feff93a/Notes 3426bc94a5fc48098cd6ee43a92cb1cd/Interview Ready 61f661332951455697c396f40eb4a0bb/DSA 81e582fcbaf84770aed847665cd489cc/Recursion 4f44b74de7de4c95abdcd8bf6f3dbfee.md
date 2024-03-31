# Recursion

[How to use Recursion - DeriveIt](https://deriveit.org/coding/recursion/how-to-use-recursion-215)

[](https://medium.com/@daniel.oliver.king/getting-started-with-recursion-f89f57c5b60e)

[](https://medium.com/@daniel.oliver.king/learning-to-think-with-recursion-part-2-887bd4c41274#.miz09qwck)

[](https://web.stanford.edu/class/archive/cs/cs106b/cs106b.1126/lectures/08/Slides08.pdf)

[Lecture 8 | Programming Abstractions (Stanford)](https://www.youtube.com/watch?v=gl3emqCuueQ)

[How to Code Combinations Using Recursion](https://www.youtube.com/watch?v=NA2Oj9xqaZQ)

---

Defining a problem in terms of itself

- Base case
- Work towards base case - make the problem smaller
- Recursive call

The base case is far easier to identify, do not worry about it, it will all make sense once you write the recursive calls.

## Golden rules for solving recursive calls

- Break the problem into a problem that is one step simpler, with emphasis on the one-step simpler
- Assume that your function works
- Ask yourself: since I know how to solve the simpler problem, how would I solve the more complex problem?

## Write a function that reverses a string.

- One step simpler
    - reversing a string with one character shorter
- Assume the function works and can reverse a string that is one letter shorter than the original string.
- first_char + func(string without first_char)

Don't think how you would break the problem all the way down to the base case. Think about one step simpler to the actual problem,.

Recursion implicitly uses the stack, hence all-recursive approaches can be rewritten using a stack.

***Combination, Permutation, trees = Recursion***

```python
def get_permutations(str):
    if len(str) <= 1:
        return [str]
    first = str[0]
    last = str[1:]

    rest_permutations = get_permutations(last)  # I assume that this function returns
    # me the list of permutations

    permutations = []

    for perm in rest_permutations:
        for i in range(len(perm) + 1):
            temp = perm[:i] + first + perm[i:]
            permutations.append(temp)

    return permutations

input_string = "abc"
permutations = get_permutations(input_string)
print(permutations)
```

[The Recursive Book of Recursion](https://inventwithpython.com/recursion/)

[Recursion in Programming - Full Course](https://www.youtube.com/watch?si=hTP07XNhJi0eMFox&v=IJDJ0kBx2LM&feature=youtu.be)