from abc import ABC, abstractmethod
from typing import Any, List

from .types import RankingRequest


class Queue(ABC):
    @abstractmethod
    def poll_ranker_queue(self) -> List[RankingRequest]:
        pass

    @abstractmethod
    def delete_message(self, message):
        pass

    @abstractmethod
    def send_ranking_response(self, ranking_response: Any):
        pass
