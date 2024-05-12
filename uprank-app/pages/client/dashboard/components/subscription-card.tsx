import {
    CardTitle,
    CardHeader,
    CardContent,
    CardFooter,
    Card,
} from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Separator } from "@/components/ui/separator";
import { Button } from "@/components/ui/button";
import { Subscription } from "@/types/subscription-type";
import { capitalizeFirstLetter } from "@/utils/client_utils/capitalize_first_letter";
import { useToast } from "@/components/ui/use-toast";
interface SubscriptionCardProps {
    subscription: Subscription;
}
import "dayjs/locale/es";
import dayjs from "dayjs";
import { cn } from "@/lib/utils";
import craft_api_url from "@/utils/api_utils/craft_api_url";
import { useRouter } from "next/router";
import { useUser } from "@clerk/nextjs";

export default function SubscriptionCard({
    subscription,
}: SubscriptionCardProps) {
    const { toast } = useToast();
    const router = useRouter();
    const { user } = useUser();

    const handleCancelSubscription = async (subscription_id: number) => {
        try {
            const response = await fetch(
                craft_api_url(
                    `/api/private/subscription/${subscription_id}/deactivate`
                ),
                {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify({
                        client_id: user?.id,
                    }),
                }
            );
            if (response.ok) {
                toast({
                    description:
                        "Successfully cancelled your subscription for the next biling cycle",
                });
                router.replace(router.asPath, undefined, { scroll: false });
            } else {
                const error = await response.json();
                console.log(error);
            }
        } catch (error) {
            alert(
                `Unable to cancel subscription due to: ${error}. Please email useUprankdesign@gmail.com`
            );
        }
    };

    return (
        <Card className="w-full max-w-md">
            <CardHeader className="flex items-center justify-between">
                <div className="space-y-1">
                    <CardTitle>
                        {capitalizeFirstLetter(
                            subscription.freelancer.first_name
                        ) +
                            " " +
                            capitalizeFirstLetter(
                                subscription.freelancer.last_name
                            )}
                    </CardTitle>
                    <Badge
                        variant="outline"
                        className={cn(
                            !subscription.end_date
                                ? ""
                                : "text-red-500"
                        )}
                    >
                        {subscription.freelancer_confirmation ? (subscription.end_date ? `Ending ${subscription.end_date}` : "Active" ): "Pending confirmation"}
                    </Badge>
                </div>
            </CardHeader>
            <CardContent className="grid gap-4">
                <div className="grid gap-1">
                    <div className="text-sm text-gray-500 dark:text-gray-400">
                        Subscription Details
                    </div>
                    <div className="flex items-center justify-between text-sm text-gray-500 dark:text-gray-400">
                        <span>Start Date</span>
                        <span>
                            {dayjs(subscription.start_date).format(
                                "MMMM D, YYYY"
                            )}
                        </span>
                    </div>
                    <div className="flex items-center justify-between text-sm text-gray-500 dark:text-gray-400">
                        <span>
                            {!subscription.end_date
                                ? "Next renewal date"
                                : "End Date"}
                        </span>
                        <span>
                            {!subscription.end_date
                                ? dayjs(subscription.start_date)
                                      .add(1, "month")
                                      .format("MMMM D, YYYY")
                                : subscription.end_date}
                        </span>
                    </div>
                    <div className="flex items-center justify-between text-sm text-gray-500 dark:text-gray-400">
                        <span>Renewal</span>
                        <span
                            className={cn(
                                !subscription.end_date ? "" : "text-red-500"
                            )}
                        >
                            {!subscription.end_date ? "Automatic" : "Suspended"}
                        </span>
                    </div>
                </div>
                <Separator />
                <div className="grid gap-1">
                    <div className="text-sm text-gray-500 dark:text-gray-400">
                        Payment Details
                    </div>
                    <div className="flex items-center justify-between">
                        <div className="flex items-center gap-2">
                            <CreditCardIcon className="h-5 w-5 text-gray-500 dark:text-gray-400" />
                            <span className="text-sm">Visa ending in 1234</span>
                        </div>
                        <div className="text-sm font-medium">$49.99</div>
                    </div>
                    <div className="flex items-center justify-between text-sm text-gray-500 dark:text-gray-400">
                        <span>Paid on</span>
                        <span>April 15, 2023</span>
                    </div>
                </div>
            </CardContent>
            <CardFooter className="flex justify-end gap-2">
                {(!subscription.end_date && subscription.freelancer_confirmation) && (
                    <Button
                        variant="outline"
                        onClick={() =>
                            handleCancelSubscription(subscription.id)
                        }
                    >
                        Cancel Subscription
                    </Button>
                )}
                <Button>Manage Subscription</Button>
            </CardFooter>
        </Card>
    );
}

function CreditCardIcon(props: any) {
    return (
        <svg
            {...props}
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
        >
            <rect width="20" height="14" x="2" y="5" rx="2" />
            <line x1="2" x2="22" y1="10" y2="10" />
        </svg>
    );
}

function WalletIcon(props: any) {
    return (
        <svg
            {...props}
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
        >
            <path d="M19 7V4a1 1 0 0 0-1-1H5a2 2 0 0 0 0 4h15a1 1 0 0 1 1 1v4h-3a2 2 0 0 0 0 4h3a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1" />
            <path d="M3 5v14a2 2 0 0 0 2 2h15a1 1 0 0 0 1-1v-4" />
        </svg>
    );
}
