import validator.checks as checks


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
        self.errors = []

    @property
    def data(self):
        return self._data

    @property
    def raw(self):
        return self._raw

    def run_checks(self):
        checks.auth_type(auth_type=self.auth_type)
        checks.provider_name(provider_name=self.provider_name)
        checks.provider_type(
            provider_type=self.provider_type, provider_name=self.provider_name
        )
        checks.auth_azure(
            provider_name=self.provider_name,
            subscription=self.azure_subscription,
            tenant=self.azure_tenant,
        )
        checks.auth_type_user(
            auth_type=self.auth_type,
            user=self.auth["user"],
            password=self.auth["password"],
        )

        if len(checks.exceptions) > 0:
            raise Exception("There were one or more error(s) in your configuration")
