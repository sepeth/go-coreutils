import os
import shlex
import subprocess
import tempfile
import unittest


class CoreUtilsTestCase(unittest.TestCase):

    command = None

    def get_process_output(self, cmd):
        cmd = shlex.split('%s %s' % (self.command, cmd))
        process = subprocess.Popen(cmd, stdout=subprocess.PIPE,
                                   stderr=subprocess.PIPE)
        out, err = process.communicate()
        rc = process.returncode
        return out, err, rc

    def check_output(self, finput, foutput, args=None):
        """Saves the input to a temp file, runs the command,
        matches the output.

        """
        # The *delete* parameter should be False,
        # since we want to close the file and test it later.
        with tempfile.NamedTemporaryFile(dir='.', delete=False) as f:
            f.write(finput)
        try:
            if args is None:
                cmd = '%s' % f.name
            else:
                cmd = '%s %s' % (args, f.name)
            out, err, rc = self.get_process_output(cmd)

            self.assertEqual(out, foutput)
            if rc:
                self.assertEqual(err, '')
        finally:
            os.unlink(f.name)
