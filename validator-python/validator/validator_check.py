import logging
from abc import ABC, abstractmethod
from typing import List, Callable
from dataclasses import dataclass, field


def error_call(err, exceptions, *args, **kwargs):
    logging.error(err)
    logging.error(f"{args}, {kwargs}")
    exceptions.append(err)


@dataclass
class ValidatorCheck(ABC):
    exceptions: List[Exception] = field(default_factory=list)
    # if this is not a staticmethod the self argument is still passed onto the error_call function
    errorcall: Callable = staticmethod(error_call)
    pass_msg: str = f"[*] PASS"

    @abstractmethod
    def run_checks() -> List[Exception]:
        pass
