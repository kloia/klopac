class Platform:
    def __init__(self, platform: dict) -> None:
        self._raw = platform
        self._data = platform["platform"]

    @property
    def data(self):
        return self._data

    @property
    def raw(self):
        return self._raw
