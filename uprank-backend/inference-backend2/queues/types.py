class RankingRequest:
    def __init__(
        self,
        body: str,
        job_id: str,
        user_id: str,
        short_lived_token: str,
        receipt_handle: str,
        platform: str,
        platform_id: str,
    ) -> None:
        self.body = body
        self.job_id = job_id
        self.user_id = user_id
        self.short_lived_token = short_lived_token
        self.platform = platform
        self.platform_id = platform_id
        self.receipt_handle = receipt_handle
