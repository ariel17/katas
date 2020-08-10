#!/usr/bin/env python
# -*- coding: utf-8 -*-

import unittest

from binary_search import binary_search


class TestBinarySearch(unittest.TestCase):
    def test_kata(self):
        self.assertEqual(-1, binary_search(3, []))
        self.assertEqual(-1, binary_search(3, [1]))
        self.assertEqual(0,  binary_search(1, [1]))

        self.assertEqual(0,  binary_search(1, [1, 3, 5]))
        self.assertEqual(1,  binary_search(3, [1, 3, 5]))
        self.assertEqual(2,  binary_search(5, [1, 3, 5]))
        self.assertEqual(-1, binary_search(0, [1, 3, 5]))
        self.assertEqual(-1, binary_search(2, [1, 3, 5]))
        self.assertEqual(-1, binary_search(4, [1, 3, 5]))
        self.assertEqual(-1, binary_search(6, [1, 3, 5]))

        self.assertEqual(0,  binary_search(1, [1, 3, 5, 7]))
        self.assertEqual(1,  binary_search(3, [1, 3, 5, 7]))
        self.assertEqual(2,  binary_search(5, [1, 3, 5, 7]))
        self.assertEqual(3,  binary_search(7, [1, 3, 5, 7]))
        self.assertEqual(-1, binary_search(0, [1, 3, 5, 7]))
        self.assertEqual(-1, binary_search(2, [1, 3, 5, 7]))
        self.assertEqual(-1, binary_search(4, [1, 3, 5, 7]))
        self.assertEqual(-1, binary_search(6, [1, 3, 5, 7]))
        self.assertEqual(-1, binary_search(8, [1, 3, 5, 7]))
