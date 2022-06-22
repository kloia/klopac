from validator.checks import *


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
        check_auth(auth_type=self.auth_type)
        check_provider_name(provider_name=self.provider_name)
        check_provider_type(
            provider_type=self.provider_type, provider_name=self.provider_name
        )
        check_auth_azure(
            provider_name=self.provider_name,
            subscription=self.azure_subscription,
            tenant=self.azure_tenant,
        )
        check_auth_type_user(
            auth_type=self.auth_type,
            user=self.auth["user"],
            password=self.auth["password"],
        )
