from .command import Command
from typing import Optional, List


class UpgradeCommand(Command):

    def __init__(self, args: Optional[List[str]] = None):
        self._args = args

    def run(self):
        print(f"Running Upgrade {self._args}")
