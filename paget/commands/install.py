from .command import Command
from typing import Optional, List


class InstallCommand(Command):

    def __init__(self, args: Optional[List[str]] = None):
        self._args = args

    def run(self):
        print(f"Installing {self._args}")
