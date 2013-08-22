import unittest

from . import CoreUtilsTestCase


class HeadTest(CoreUtilsTestCase):

    command = './head'

    def test_head_lines(self):
        self.check_output('', '')
        self.check_output('1\n2\n3\n4\n5\n', '1\n2\n3\n4\n5\n')
        self.check_output('1\n2\n3\n4\n5\n', '1\n', args='-n 1')
        # missing last line
        self.check_output('1x\n2x\n3\n4\n5x', '1x\n2x\n3\n4\n5x')
        # -n -2 check
        self.check_output('1\n2\n3\n4\n5\n', '1\n2\n3\n', args='-n -2')

    def test_head_bytes(self):
        self.check_output('12345', '12', args='-c 2')
        self.check_output('12345', '123', args='-c -2')

if __name__ == '__main__':
    unittest.main()
