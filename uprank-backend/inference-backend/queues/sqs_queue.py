from typing import List

from mypy_boto3_sqs.client import SQSClient

from .queue import Queue
from .types import RankingRequest


class SQSQueue(Queue):
    sqs_client: SQSClient
    ranker_queue_url: str

    def __init__(self, sqs_client: SQSClient, ranker_queue_url: str) -> None:
        super().__init__()
        self.sqs_client = sqs_client
        self.ranker_queue_url = ranker_queue_url

    def poll_ranker_queue(self) -> List[RankingRequest]:
        data: List[RankingRequest] = []
        response = self.sqs_client.receive_message(
            QueueUrl=self.ranker_queue_url,
            WaitTimeSeconds=20,
            MessageAttributeNames=["job_id", "user_id", "platform", "platform_id", "short_lived_token"],
        )
        if "Messages" in response:
            for message in response["Messages"]:
                receipt_handle = message["ReceiptHandle"]
                body = message["Body"]

                # Extract message attributes
                message_attributes = message.get("MessageAttributes", {})
                job_id = message_attributes.get("job_id", {}).get("StringValue", "")
                user_id = message_attributes.get("user_id", {}).get("StringValue", "")
                platform = message_attributes.get("platform", {}).get("StringValue", "")
                platform_id = message_attributes.get("platform_id", {}).get(
                    "StringValue", "",
                )
                data.append(
                    RankingRequest(
                        body=body,
                        job_id=job_id,
                        user_id=user_id,
                        short_lived_token="",
                        platform=platform,
                        platform_id=platform_id,
                        receipt_handle=receipt_handle,
                    ),
                )
            return data
        return None

    def delete_message(self, receipt_handle: str) -> None:
        self.sqs_client.delete_message(
            QueueUrl=self.ranker_queue_url, ReceiptHandle=receipt_handle,
        )
