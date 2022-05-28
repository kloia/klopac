import logging
from rich.logging import RichHandler

logging.basicConfig(
    format="%(message)s", datefmt="[%X]", level=logging.INFO, handlers=[RichHandler()]
)
logger = logging.getLogger(__name__)
