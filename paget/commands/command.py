from abc import abstractmethod


class Command:

    @abstractmethod
    def run(self):
        pass

    def set_args(self, args):
        self._args = args
