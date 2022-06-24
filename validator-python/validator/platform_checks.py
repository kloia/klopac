import logging
from typing import List
from validator.error import ErrorCollection
from validator.validator_check import ValidatorCheck


class PlatformChecks(ValidatorCheck):
    @ErrorCollection
    def auth_type(self, auth_type: str):
        logging.info("[*] Checking auth type for the specified provider")
        if auth_type in ["user", "id"]:
            logging.info(self.pass_msg)
            return

        raise Exception('[!] Auth type should be "user" or "id"')

    @ErrorCollection
    def provider_name(self, provider_name: str):
        logging.info("[*] Checking provider name")
        if provider_name in ["azure", "aws"]:
            logging.info(self.pass_msg)
            return

        raise Exception('[!] Provider name should be "azure" or "aws"')

    @ErrorCollection
    def provider_type(self, provider_name: str, provider_type: str):
        logging.info("[*] Checking provider type")
        if provider_name == "aws":
            if provider_type in ["eks", "ec2"]:
                logging.info(self.pass_msg)
                return

            raise Exception('[!] Provider type should be "eks" or "ec2" for AWS')

        elif provider_name == "azure":
            if provider_type in ["vm", "vmss"]:
                logging.info(self.pass_msg)
                return

            raise Exception('[!] Provider type should be "vm" or "vmss" for Azure')

        raise NotImplementedError("[!] Unsupported provider name")

    @ErrorCollection
    def auth_azure(self, provider_name: str, subscription: str, tenant: str):
        if provider_name == "azure":
            logging.info("[*] Checking required auth fields for Azure")
            self.azure_subscription(subscription)
            self.azure_tenant(tenant)

    @ErrorCollection
    def azure_subscription(self, azure_subscription: str):
        logging.info("[*] Checking subscription field for Azure")
        if len(azure_subscription) > 0:
            logging.info(self.pass_msg)
            return

        raise Exception("[!] Subscription field is required for Azure")

    @ErrorCollection
    def azure_tenant(self, azure_tenant: str):
        logging.info("[*] Checking tenant field for Azure")
        if len(azure_tenant) > 0:
            logging.info(self.pass_msg)
            return

        raise Exception("[!] Tenant field is required for Azure")

    @ErrorCollection
    def auth_type_user(self, auth_type: str, user: str, password: str):
        if auth_type == "user":
            logging.info('[*] Checking required fields for auth type "user"')
            self.auth_user(user)
            self.auth_password(password)

    @ErrorCollection
    def auth_user(self, user: str):
        logging.info('[*] Checking "user" field for auth type "user"')
        if len(user) > 0:
            logging.info(self.pass_msg)
            return

        raise Exception('[!] User field is required for auth type "user"')

    @ErrorCollection
    def auth_password(self, password: str):
        logging.info('[*] Checking "password" field for auth type "user"')
        if len(password) > 0:
            logging.info(self.pass_msg)
            return

        raise Exception('[!] Password field is required for auth type "user"')

    def run_checks(self, platform) -> List[Exception]:
        self.auth_type(auth_type=platform.auth_type)
        self.provider_name(provider_name=platform.provider_name)
        self.provider_type(
            provider_type=platform.provider_type, provider_name=platform.provider_name
        )
        self.auth_azure(
            provider_name=platform.provider_name,
            subscription=platform.azure_subscription,
            tenant=platform.azure_tenant,
        )
        self.auth_type_user(
            auth_type=platform.auth_type,
            user=platform.auth["user"],
            password=platform.auth["password"],
        )

        return self.exceptions
