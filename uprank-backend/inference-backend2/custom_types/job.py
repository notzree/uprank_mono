from __future__ import annotations

from dataclasses import dataclass, field
from datetime import datetime
from typing import Any


@dataclass
class WorkHistory:
    id: int
    title: str
    client_feedback: str
    overall_rating: float
    freelancer_earnings: float
    start_date: datetime
    end_date: datetime
    description: str
    budget: float
    client_country: str
    client_total_spend: float
    client_total_hires: int
    client_active_hires: int
    edges: dict[str, Any] = field(default_factory=dict)

@dataclass
class UpworkFreelancer:
    id: str
    name: str
    title: str
    description: str
    city: str
    country: str
    timezone: str
    cv: str
    fixed_charge_amount: float
    fixed_charge_currency: str
    hourly_charge_amount: float
    hourly_charge_currency: str
    photo_url: str
    recent_hours: float
    total_hours: float
    total_portfolio_items: int
    total_portfolio_v2_items: int
    upwork_total_feedback: float
    upwork_recent_feedback: float
    upwork_top_rated_status: bool
    upwork_job_success_score: int
    skills: list[str]
    average_recent_earnings: float
    combined_average_recent_earnings: float
    combined_recent_earnings: float
    combined_total_earnings: float
    combined_total_revenue: float
    recent_earnings: float
    total_revenue: float
    uprank_updated_at: datetime
    edges: dict[str, list[WorkHistory]] = field(default_factory=dict)

@dataclass
class Job:
    id: str
    title: str
    created_at: datetime
    location: str
    description: str
    skills: list[str]
    experience_level: str
    hourly: bool
    hourly_rate: list[int]
    edges: dict[str, list[UpworkFreelancer]] = field(default_factory=dict)
