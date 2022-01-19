from typing import Optional, List
from .command import Command


class UpdateCommand(Command):

    def __init__(self, args: Optional[List[str]] = None):
        self._args = args

    def run(self):
        print("Running Upgrade")
