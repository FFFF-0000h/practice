Algorithm for CheckNumber program:

Think of this like checking if there's any digit hidden inside a word or sentence.
Step-by-step thinking process:

1. We get a string - Imagine it's like a chain of characters: "Hello1" is like H → e → l → l → o → 1

2.  We need to check each character - Like looking at each bead on a necklace, one by one

3.  For each character, ask: "Is this a number?" - Check if it's 0, 1, 2, 3, 4, 5, 6, 7, 8, or 9

4. If we find ANY number - Stop looking and say "YES" (return true)

5. If we finish checking all characters - And found no numbers, say "NO" (return false)

----------------------------------------------------------------------------

Algorithm for CountAlpha program:

Think of this as counting only the letters of the alphabet (A-Z and a-z) in a string, ignoring everything else!
Step-by-step thinking process:

1. We get a string - Like "Hello world" which has letters, spaces, maybe numbers or symbols

2. We need a counter - Start at 0, and increase it every time we find a letter

3. Go through each character - Look at them one by one

4. For each character, ask: "Is this a letter?"

    Is it between 'a' and 'z'? OR

    Is it between 'A' and 'Z'?

5. If YES → Add 1 to our counter

6. After checking all characters → Return the final count

--------------------------------------------------------------------------

Algorithm for CountCharacter program:

Think of this as counting how many times a specific letter (or character) appears in a word or sentence - like counting how many times the letter 'l' appears in "Hello World"!
Step-by-step thinking process:

1. We get two things:

        A string to search through (like "Hello World")

        A specific character to look for (like 'l')

2. We need a counter - Start at 0, and increase it every time we find a match

3. Go through each character in the string, one by one

4. For each character, ask: "Does this match the character we're looking for?"

5. If YES → Add 1 to our counter

6. After checking all characters → Return the final count

Important Note About Types:

In Go, the function uses rune for the character. A rune is just Go's way of representing a single character (like 'l', '5', or ' ').

-----------------------------------------------------------------------------

Algorithm for PrintIf problem:

This function checks the length of a string and returns different things based on that length:

1. If string length > 3 → Return "G\n" (letter G + newline)

2. If string length ≤ 3 (but not empty) → Return "Invalid Input\n"

3. If string is empty (length = 0) → Return "G\n" (special case!)

Wait - I notice the example output shows "Invalid Output$" but the instructions say "Invalid Input". Looking at the test output, it seems the function should return "Invalid Output", not "Invalid Input". I'll use what the test expects.
Step-by-Step Algorithm

1. Check if the string is empty

        If yes → return "G\n"

2. If not empty, check the length

        If length > 3 → return "G\n"

        If length ≤ 3 → return "Invalid Output\n"

--------------------------------------------------------------------------

Algorithm for PrintIfNot problem:
    
    Rules:
    If string length < 3 → Return "G\n" (letter G + newline)

    If string length ≥ 3 → Return "Invalid Output\n"

    If string is empty (special case) → Return "G\n"

Step-by-Step Algorithm

Check the length of the string using len(str)

    If length < 3 OR length == 0 → return "G\n"

    Otherwise (length ≥ 3) → return "Invalid Output\n"
----------------------------------------------------------------------------

Algorithm for RectPerimeter problem:

What is Perimeter?

The perimeter of a rectangle is the distance around the outside. You calculate it by:

    Adding up all 4 sides

    Or using the formula: Perimeter = 2 × (width + height)

The Rules:

1. Normal case: Calculate perimeter = 2 × (w + h)

2. Error case: If ANY argument is negative → return -1

3. Return type: The function returns an int (integer)

Step-by-Step Algorithm

1. Check for negative numbers

        If width < 0 OR height < 0 → return -1

2. If both numbers are valid (0 or positive)

        Calculate perimeter = 2 × (width + height)

        Return the result
---------------------------------------------------------------------------

Algorithm for RetainFirstHalf problem:

What "First Half" Means:

    "Hello World" (length 11) → First half is positions 0-4 → "Hello"

    "A" (length 1) → First half is just "A"

    "" (empty) → First half is ""

The Rules:

    Normal case: Take the first half of the string

    If length is odd: Round DOWN (e.g., length 11 → take first 5 characters)

    If empty string: Return empty string ""

    If length is 1: Return the string itself

Step-by-Step Algorithm

1. Get the length of the string with len(str)

2. Calculate the halfway point:

        For even length (10): halfway = 5

        For odd length (11): halfway = 5 (round down)

        Formula: halfway := len(str) / 2 (Go integer division automatically rounds down!)

3. Return the first half using slicing: str[:halfway]
