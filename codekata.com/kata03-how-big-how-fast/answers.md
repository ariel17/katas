# Answers

## How big?

* Bits:

  + 1,000: 10 digits
  + 1,000,000: 20 digits
  + 1,000,000,000: 30 digits
  + 1,000,000,000,000: 40 digits
  + 8,000,000,000,000: 43 digits
  
  Every 1,000 numbers 10 bits are required to represent it.

* Residents' data:

  * 20,000 residents
  * Names: 30 characters
  * Address 100 characters
  * Phone number: 20 characters

  total = 20,000 * (30 + 100 + 20) = 3,000,000 bits = 375 KB

* Binary tree:

  * N = 1,000,000
  * The minimum height of a binary tree: `floor(ln N)` => 13 nodes
  * In a 32 bit architecture the memory address is 32 bit wide = 4 bytes.
  * The representation range for unsigned int is 0-4,294,967,296. Assuming the
    tree will store numbers in that range.
  * A node is represented by:

    + value
    + left pointer
    + right pointer

    total value = 3 * 4 bytes * 1,000,000 = 12,000,000 bytes = 11718.75 MB

## How fast?
