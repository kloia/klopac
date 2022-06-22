import logging


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

    def run_checks(self):
        self.check_auth()
        self.check_provider_name()
        self.check_provider_type()
        self.check_auth_azure()

    def check_auth(self):
        logging.info("[*] Checking auth type for the specified provider")
        if self.auth_type in ["user", "id"]:
            logging.info("[*] PASS")
            return

        raise Exception('[!] Auth type should be "user" or "id"')

    def check_provider_name(self):
        logging.info("[*] Checking provider name")
        if self.provider_name in ["azure", "aws"]:
            logging.info("[*] PASS")
            return

        raise Exception('[!] Provider name should be "azure" or "aws"')

    def check_provider_type(self):
        logging.info("[*] Checking provider type")
        if self.provider_name == "aws":
            if self.provider_type in ["eks", "ec2"]:
                logging.info("[*] PASS")
                return

            raise Exception('[!] Provider type should be "eks" or "ec2" for AWS')

        elif self.provider_name == "azure":
            if self.provider_type in ["vm", "vmss"]:
                logging.info("[*] PASS")
                return

            raise Exception('[!] Provider type should be "vm" or "vmss" for Azure')

        raise NotImplementedError("[!] Unsupported provider name")

    def check_auth_azure(self):
        if self.provider_name == "azure":
            logging.info("[*] Checking required auth fields for Azure")
            self.check_azure_subscription()
            self.check_azure_tenant()

    def check_azure_subscription(self):
        logging.info("[*] Checking subscription field for Azure")
        if len(self.azure_subscription) > 0:
            logging.info("[*] PASS")
            return

        raise Exception("[!] Subscription field is required for Azure")

    def check_azure_tenant(self):
        logging.info("[*] Checking tenant field for Azure")
        if len(self.azure_tenant) > 0:
            logging.info("[*] PASS")
            return

        raise Exception("[!] Tenant field is required for Azure")
