import logging


def check_auth(auth_type: str):
    logging.info("[*] Checking auth type for the specified provider")
    if auth_type in ["user", "id"]:
        logging.info("[*] PASS")
        return

    raise Exception('[!] Auth type should be "user" or "id"')


def check_provider_name(provider_name: str):
    logging.info("[*] Checking provider name")
    if provider_name in ["azure", "aws"]:
        logging.info("[*] PASS")
        return

    raise Exception('[!] Provider name should be "azure" or "aws"')


def check_provider_type(provider_name: str, provider_type: str):
    logging.info("[*] Checking provider type")
    if provider_name == "aws":
        if provider_type in ["eks", "ec2"]:
            logging.info("[*] PASS")
            return

        raise Exception('[!] Provider type should be "eks" or "ec2" for AWS')

    elif provider_name == "azure":
        if provider_type in ["vm", "vmss"]:
            logging.info("[*] PASS")
            return

        raise Exception('[!] Provider type should be "vm" or "vmss" for Azure')

    raise NotImplementedError("[!] Unsupported provider name")


def check_auth_azure(provider_name: str, subscription: str, tenant: str):
    if provider_name == "azure":
        logging.info("[*] Checking required auth fields for Azure")
        check_azure_subscription(subscription)
        check_azure_tenant(tenant)


def check_azure_subscription(azure_subscription: str):
    logging.info("[*] Checking subscription field for Azure")
    if len(azure_subscription) > 0:
        logging.info("[*] PASS")
        return

    raise Exception("[!] Subscription field is required for Azure")


def check_azure_tenant(azure_tenant: str):
    logging.info("[*] Checking tenant field for Azure")
    if len(azure_tenant) > 0:
        logging.info("[*] PASS")
        return

    raise Exception("[!] Tenant field is required for Azure")


def check_auth_type_user(auth_type: str, user: str, password: str):
    if auth_type == "user":
        logging.info('[*] Checking required fields for auth type "user"')
        check_auth_user(user)
        check_auth_password(password)


def check_auth_user(user: str):
    logging.info('[*] Checking "user" field for auth type "user"')
    if len(user) > 0:
        logging.info("[*] PASS")
        return

    raise Exception('[!] User field is required for auth type "user"')


def check_auth_password(password: str):
    logging.info('[*] Checking "password" field for auth type "user"')
    if len(password) > 0:
        logging.info("[*] PASS")
        return

    raise Exception('[!] Password field is required for auth type "user"')
