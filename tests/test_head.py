# this is the test implementation for head command,
# gnu-coreutils uses perl but its really complicated to run
# gnu's test suite and did i mention it uses perl.
# 
# anyways so i wrote a simple one in python.
#
import unittest
import os
from subprocess import Popen, PIPE
from tempfile import NamedTemporaryFile

class TestHead(unittest.TestCase):

    def get_process_output(self, *args, **kwargs):
        if 'exit_code' not in kwargs:
            kwargs['exit_code'] = 0
        process = Popen(args, stdout=PIPE)
        exit_code = os.waitpid(process.pid, kwargs['exit_code'])
        return process.communicate()[0]

    def setUp(self):
        pass

    def check_output(self, case):
        """
        saves the input to a temp file, runs the command, 
        matches the output
        """
        # this should be false, since we want to close the
        # file and test it later.
        f = NamedTemporaryFile(dir=".", delete=False)
        f.write(case['in'])
        f.close()
        #print f.name
        has_params = 'params' in case
        if 'params' in case:
            args = []
            args.append('../head')
            for param in case['params']:
                args.append(param)
            args.append(f.name)
            out = self.get_process_output(*args)
        else:
            out = self.get_process_output('../head', f.name)
        
        #print out
        assert out == case['out']
        # then we delete the tmp file
        os.unlink(f.name)

    def test_head_lines(self):
        self.check_output({'in':"", 'out':""})
        self.check_output({'in':"1\n2\n3\n4\n5\n", 'out':"1\n2\n3\n4\n5\n"})
        # missing last line
        self.check_output({'in':"1x\n2x\n3\n4\n5x", 'out':"1x\n2x\n3\n4\n5x"})
        # 
        case = {'params':['-n', '1'],'in':"1\n2\n3\n4\n5\n", 'out':"1\n"}
        self.check_output(case)
        # -n -2 check
        case = {'params':['-n', '-2'],'in':"1\n2\n3\n4\n5\n", 'out':"1\n2\n3\n"}
        self.check_output(case)

    def test_head_bytes(self):
        self.check_output({'params':['-c', '2'], 'in':"12345", 'out':"12"})
        self.check_output({'params':['-c', '-2'], 'in':"12345", 'out':"123"})



if __name__ == '__main__':
    unittest.main() 