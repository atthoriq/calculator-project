# Project Calculator

A traditional calculator simulator that takes a string input contains arithmetic operation followed by a floating number. It works by doing the operation to the current result of previous computation. Initial current result is 0.

Commands that are supported can be seen through 'help' command:

```
> help
add <float>      : add <float> to current
subtract <float> : subtract <float> to current
multiply <float> : add <float> to current
divide <float>   : add <float> to current
neg              : make current to negative. equally multiplying -1 to current. it requires no <float>
abs              : make current to positive. it requires no <float>
sqrt             : compute sqrt of current
cbrt             : compute cbrt of current
sqr              : compute sqr of current
cube             : compute cube of current
repeat <float>   : repeating <float> steps behind
cancel           : cancel calculation which set the current to 0.
exit             : exit the calculator
help             : show the manual
```

There are 2 packages in the repository, main and calculator package. Handler is put in the main package to improve readability. However, I create a dedicated package for the calculator implementation so its private function remain private. Feedback are welcome for this structure!

## How to Run Locally

Using this command:
```
make run
```

## Requirement Limitation

1. If single command (i.e. neg, abs, sqrt, cbrt, etc.) is given a value or additional argument, it will return an error and exit the program.
2. Any complex arithmetic operator is done by golang built-in package called 'math' to ensure correctness.
3. Division by 0 will result NaN.
4. A result of NaN won't immediately exit the program. Instead user can cancel the calculation or restart the program.
