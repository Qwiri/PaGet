from typing import List, Literal
from rich import print

from . import help_str
from .commands.command import Command
from .commands.install import InstallCommand
from .commands.help import HelpCommand
from .commands.upgrade import UpgradeCommand
from .commands.update import UpdateCommand


class CLI:

    def __init__(self, args: List[str]) -> None:
        self._args = args

    def __clear_first_element(self):
        self._args = self._args[1:]

    def parse_arguments():
        pass

    def parse_command(self) -> Command:
        # Keine Argumente Update und Upgrade laufen
        self.__clear_first_element()

        if len(self._args) == 0:
            return UpgradeCommand()

        # Argumente mit Command, Command mit Argumente ausführen
        command = self._args[0].lower()
        if command == "update":
            command = UpdateCommand()
            self.__clear_first_element()
        elif command == "remove":
            command = UpgradeCommand()
            self.__clear_first_element()
        elif command == "help":
            return HelpCommand()
        else:
            command = InstallCommand()

        command.set_args(self._args)

        return command

    def run(self) -> None:
        cmd = self.parse_command()
        cmd.run()

    @staticmethod
    def print_help() -> None:
        print(help_str)
