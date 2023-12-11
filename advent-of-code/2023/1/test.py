import unittest
from main import part1, part2

class Test(unittest.TestCase):
    def test_part1(self):
        result = part1("sample-1.txt")
        self.assertEqual(result, 142)
        
    def test_part2(self):
        result = part2("sample-2.txt")
        self.assertEqual(result, 281)

if __name__ == '__main__':
    unittest.main()