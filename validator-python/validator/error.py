def ErrorCollection(func):
    def collect_error_wrapper(self, *args, **kwargs):
        try:
            return func(self, *args, **kwargs)
        except Exception as err:
            self.errorcall(err, self.exceptions, *args, **kwargs)

    return collect_error_wrapper
