## Algorithm for CamelToSnakeCase problem:

### What's the Difference?

- **camelCase**: Words are joined together, each new word starts with a capital letter (except possibly the first word)
  - Examples: `helloWorld`, `HelloWorld`, `camelCase`

- **snake_case**: Words are joined with underscores, all lowercase
  - Example: `hello_world`, `camel_case`

### The Rules:

1. **If empty string** → return empty string `""`
2. **If NOT valid camelCase** → return the string unchanged
3. **If valid camelCase** → convert to snake_case

### What Makes Valid CamelCase?

The problem gives us rules:
-  No capital letter at the end (`CamelCasE`)
-  No two capitals in a row (`CamelCAse`)
-  No numbers or punctuation anywhere (`camelCase1`)
-  Can start with capital (UpperCamelCase) or lowercase (lowerCamelCase)

---

## Step-by-Step Algorithm

### Step 1: Check for empty string
```
IF length of string == 0 THEN
    RETURN ""
```

### Step 2: Validate camelCase format
```
FOR each character at position i from 0 to length-1:
    // Check for invalid characters (not A-Z or a-z)
    IF character is NOT a letter THEN
        RETURN original string
    
    // Check for two uppercase letters in a row
    IF i > 0 AND current character is uppercase 
       AND previous character is uppercase THEN
        RETURN original string

// Check if last character is uppercase
IF last character is uppercase THEN
    RETURN original string
```

### Step 3: Convert to snake_case (only if validation passed)
```
CREATE empty string result

FOR each character at position i from 0 to length-1:
    IF current character is uppercase THEN
        // Add underscore before uppercase letters (except first character)
        IF i > 0 THEN
            ADD "_" to result
        
        // Add the uppercase character (keep original case)
        ADD character to result
    ELSE
        // Add lowercase characters as-is
        ADD character to result

RETURN result
```

---

## Visual Examples

### Example 1: `"helloWorld"` (Valid)
```
Step 2 - Validation:
h e l l o W o r l d
✓ ✓ ✓ ✓ ✓ ✓ ✓ ✓ ✓ ✓  All are letters
No two capitals in a row
Last character 'd' is not uppercase → VALID

Step 3 - Conversion:
i=0: 'h' (lowercase) → result = "h"
i=1: 'e' (lowercase) → result = "he"
i=2: 'l' (lowercase) → result = "hel"
i=3: 'l' (lowercase) → result = "hell"
i=4: 'o' (lowercase) → result = "hello"
i=5: 'W' (uppercase, i>0) → add "_" + "W" → result = "hello_W"
i=6: 'o' (lowercase) → result = "hello_Wo"
i=7: 'r' (lowercase) → result = "hello_Wor"
i=8: 'l' (lowercase) → result = "hello_Worl"
i=9: 'd' (lowercase) → result = "hello_World"

RETURN "hello_World"
```

### Example 2: `"HelloWorld"` (Valid)
```
Step 3 - Conversion:
i=0: 'H' (uppercase, i=0) → NO underscore → result = "H"
i=1: 'e' (lowercase) → result = "He"
i=2: 'l' (lowercase) → result = "Hel"
i=3: 'l' (lowercase) → result = "Hell"
i=4: 'o' (lowercase) → result = "Hello"
i=5: 'W' (uppercase, i>0) → add "_" + "W" → result = "Hello_W"
i=6: 'o' (lowercase) → result = "Hello_Wo"
i=7: 'r' (lowercase) → result = "Hello_Wor"
i=8: 'l' (lowercase) → result = "Hello_Worl"
i=9: 'd' (lowercase) → result = "Hello_World"

RETURN "Hello_World"
```

### Example 3: `"CAMELtoSnackCASE"` (INVALID - two capitals in a row)
```
Step 2 - Validation:
C A M E L t o S n a c k C A S E
              ↑
Check position 5: 't' is lowercase
But earlier at positions 0-1: 'C' and 'A' are both uppercase
i=1: s[1]='A' is uppercase AND s[0]='C' is uppercase
     → TWO CAPITALS IN A ROW found!
     → RETURN original "CAMELtoSnackCASE"
```

### Example 4: `"hey2"` (INVALID - contains number)
```
Step 2 - Validation:
h e y 2
      ↑
Position 3: '2' is NOT a letter
→ RETURN original "hey2"
```

### Example 5: `"camelCase"` (Valid)
```
Step 3 - Conversion:
c a m e l C a s e
              ↑
i=5: 'C' uppercase, i>0 → add "_" + "C"
Result: "camel_Case"
```

-----------------------------------------------------------------------

## Algorithm for DigitLen problem:

Imagine you have a number and you want to know how many digits it would have if you wrote it in a different numbering system (like binary, hexadecimal, etc.).

For example:

- The number 100 in base 10 (decimal) is "100" → that's 3 digits

- The number 100 in base 2 (binary) is "1100100" → that's 7 digits

The function counts digits by repeatedly dividing the number by the base until it reaches zero.
Step-by-Step Algorithm

Let's trace through DigitLen(100, 10):

- Start with n = 100, count = 0

- Is n > 0? Yes → divide: 100 / 10 = 10, count = 1

- Is n > 0? Yes → divide: 10 / 10 = 1, count = 2

- Is n > 0? Yes → divide: 1 / 10 = 0, count = 3

- Is n > 0? No → stop, return count = 3

For DigitLen(-100, 16):

- Negative number → make it positive: 100

- Start with n = 100, count = 0

- 100 / 16 = 6, count = 1

- 6 / 16 = 0, count = 2

    Return 2
-----------------------------------------------------------------------------

Algorithm for FirstWord problem:

Think of this program as a "first word finder" - it takes a sentence and gives you back just the first word.
What is a "word" in this context?

- Any sequence of characters that is not a space

- A word ends when you hit a space or the end of the string

    So in "hello there", "hello" is the first word

    In "hello ......... bye", "hello" is still the first word

The Algorithm (Step by Step)

1. Check if we got EXACTLY ONE argument
   - If not: do nothing (exit)

2. Take that argument (the sentence) and:
   - Find where the first word ends
   - A word ends when we find a space OR reach the end

3. Print only the characters from start to where the word ends

Intuitive Way to Think About It

Imagine you have a sentence written on a strip of paper. Your task is to cut out the first word:

    START at the beginning of the paper

    Move your finger forward until you see a space or reach the end

    CUT the paper from the start to where your finger is

    That's your first word!
----------------------------------------------------------------------------

Algorithm for FishAndChips problem:

Think of this like a decision tree - we check different conditions about a number and return different messages based on what we find.
The Rules in Plain English:

    If the number is less than 0 → say "error: number is negative"

    If the number can be divided by BOTH 2 AND 3 → say "fish and chips"

    If the number can be divided by ONLY 2 → say "fish"

    If the number can be divided by ONLY 3 → say "chips"

    If the number can't be divided by 2 or 3 → say "error: non divisible"

The Algorithm Logic

Here's how to think about it in order:
text

START
  Step 1: Check if number is negative
    IF yes → return "error: number is negative"
    
  Step 2: Check if number is divisible by BOTH 2 and 3
    IF yes → return "fish and chips"
    
  Step 3: Check if number is divisible by ONLY 2
    IF yes → return "fish"
    
  Step 4: Check if number is divisible by ONLY 3
    IF yes → return "chips"
    
  Step 5: If none of the above are true
    return "error: non divisible"
END

Important Note About Order

The order matters! You must check for "both 2 and 3" FIRST. Why? Because if a number is divisible by both (like 6), it's also divisible by 2 alone. If you check for "divisible by 2" first, you'd get "fish" instead of "fish and chips".
