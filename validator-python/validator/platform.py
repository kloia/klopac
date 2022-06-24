class Platform:
    def __init__(self, platform: dict) -> None:
        self._raw = platform
        self._data = platform["platform"]
        self.auth = self.data["provider"]["auth"]
        self.auth_type = self.auth["type"]
        self.azure_subscription = self.auth["subscription"]
        self.azure_tenant = self.auth["tenant"]
        self.provider_name = self.data["provider"]["name"]
        self.provider_type = self.data["provider"]["type"]

    @property
    def data(self):
        return self._data

    @property
    def raw(self):
        return self._raw
