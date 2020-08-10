# Answers

## How big?

* Bits:

  + 1,000: 10 digits
  + 1,000,000: 20 digits
  + 1,000,000,000: 30 digits
  + 1,000,000,000,000: 40 digits
  + 8,000,000,000,000: 43 digits
  
  Every 4 decimal numbers, 10 bits are required for representation.

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

* Modem transmission:

  56 Kbauds = 115,200 bps (bits per second).
  1,200 pages 
  Average letters in a book page: 1,800

  (1,800 * 1,200 * 8 bits per character) / 115,200 bps =
  17,280,000 b / 115,200 bps = 150 seconds

* Binary search:

  Algorithm is O(log n) in average, best/worst case is O(1).
  An increment of 10x in array size takes 1.5 ms more time to resolve => 10,000,000 will take 10 ms (1.5 ms * 2 + logarithmical delta).

* Unix passwords:

  16 positions ^ 96 possible characters * 1 ms to generate hash = 39402006196394479212279040100143613805079739270465446667948293404245721771497210611414266254884915640806627990306816 ms
  .
  .
  .
  ~1249429420230672222611588029558080092753670068190811982114037715761216443794305257845454916758146741527353.75413200202942668696 ***YEARS***
  
  0.o
  I don't think it is possible.
