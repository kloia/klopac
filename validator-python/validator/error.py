class ErrorCollection:
    def __init__(self, errorcall):
        self.errorcall = errorcall

    def __call__(self, func):
        def collect_error_wrapper(*args, **kwargs):
            try:
                return func(*args, **kwargs)
            except Exception as err:
                self.errorcall(err, *args, **kwargs)

        return collect_error_wrapper
