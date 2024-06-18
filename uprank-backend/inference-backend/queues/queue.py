from abc import ABC, abstractmethod
from typing import List
from .types import RankingRequest


class Queue(ABC):
    @abstractmethod
    def poll_ranker_queue(self) -> List[RankingRequest]:
        pass

    @abstractmethod
    def delete_message(self, message):
        pass


