from .cli import CLI
import sys


def main() -> None:
    app = CLI(sys.argv)
    app.run()


if __name__ == "__main__":
    main()
