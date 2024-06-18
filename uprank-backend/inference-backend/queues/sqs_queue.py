from .queue import Queue
from .types import RankingRequest
from mypy_boto3_sqs.client import SQSClient
import json
from typing import List


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
            WaitTimeSeconds=5,
            MessageAttributeNames=['job_id', 'user_id']
        )
        if 'Messages' in response:
            for message in response['Messages']:
                receipt_handle = message['ReceiptHandle']
                body = (message['Body'])

                # Extract message attributes
                message_attributes = message.get('MessageAttributes', {})
                job_id = message_attributes.get('job_id', {}).get('StringValue', '')
                user_id = message_attributes.get('user_id', {}).get('StringValue', '')
                data.append({
                    'body': body,
                    'job_id': job_id,
                    'user_id': user_id,
                    'receipt_handle': receipt_handle
                })
            return data

    def delete_message(self, receipt_handle: str) -> None:
        self.sqs_client.delete_message(
            QueueUrl=self.ranker_queue_url,
            ReceiptHandle=receipt_handle
        )

