import Image from "next/image";
import { Inter } from "next/font/google";
const inter = Inter({ subsets: ["latin"] });
import Link from "next/link";
import { Button } from "@/components/ui/button";
import { CardHeader, CardContent, Card } from "@/components/ui/card";
import { SignOutButton } from "@clerk/nextjs";
import { useUser } from "@clerk/nextjs";
import  mux_user_type  from "@/utils/api_utils/mux_user_type";
type CardData = {
    title: string;
    content: string;
};

const FreelancerCardData: CardData[] = [
    {
        title: "Demand based subscription pricing",
        content:
            "Tired of undercutting your rate to win clients? Uprank uses demand based pricing to ensure you get paid what you're worth. No more inflated ratings. Your rate is based on demand for your services.",
    },
    {
        title: "F.I.R.S (Fast Iteration Request System)",
        content:
            "Get both feedback and requests sent to you in real-time, with a powerful annotation system allowing detailed and precise feedback. No more waiting for emails or scheduling calls.",
    },
    {
        title: "Short Application process",
        content:
            "We know your time is valuable. Our application process is short and sweet. No more long forms or waiting for approval. Get started in minutes and hear back within a week.",
    },
];
export default function Component() {
    const { user } = useUser();
    console.log(user);
    return (
        <div key="1" className="flex flex-col min-h-screen">
            <header className="flex items-center justify-between px-8 py-6 border-b">
                <div className="flex items-center">
                    <ReplyIcon className="w-6 h-6 mr-2" />
                    <Link className="text-lg font-bold" href="#">
                        Uprank
                    </Link>
                </div>
                <nav className="flex gap-4">
                    {user?.id ? (
                        <Link
                            className="text-sm font-medium hover:underline"
                            href = {mux_user_type(user?.unsafeMetadata?.type as string) + "/dashboard"}
                        >
                            Dashboard
                        </Link>
                    ) : (
                        <>
                            <Link
                                className="text-sm font-medium hover:underline"
                                href="/sign-up"
                            >
                                Sign up
                            </Link>
                            <p>or</p>
                            <Link
                                className="text-sm font-medium hover:underline"
                                href="/sign-in"
                            >
                                Sign in
                            </Link>
                        </>
                    )}
                </nav>
            </header>
            <main className="flex-1">
                <section className="py-24 px-10 text-center">
                    <h1 className="text-4xl font-bold">
                        Uprank lets freelance designers scale vertically to $1M+
                        ARR
                    </h1>
                    <div className="flex justify-center items-center">
                        <p className="mt-6 text-lg text-foreground w-[40rem] text-left"></p>
                    </div>
                </section>
                <section className="py-24 px-10 bg-background" id="features">
                    <h2 className="text-3xl font-bold text-center">Features</h2>
                    <div className="mt-10 grid gap-10 grid-cols-1 sm:grid-cols-2 lg:grid-cols-3">
                        {FreelancerCardData.map(
                            (card: CardData, index: any) => (
                                <Card key={index}>
                                    <CardHeader>
                                        <h3 className="text-lg font-bold">
                                            {card.title}
                                        </h3>
                                    </CardHeader>
                                    <CardContent>
                                        <p className="text-sm text-gray-600">
                                            {card.content}
                                        </p>
                                    </CardContent>
                                </Card>
                            )
                        )}
                    </div>
                </section>

                <section className="py-24 px-10 text-center">
                    <h2 className="text-3xl font-bold">Sign up for free</h2>
                    <p className="mt-6 text-lg text-gray-600">
                        Join thousands of freelance designers growing their
                        business with Remark
                    </p>
                    <Button className="mt-10">Sign Up for Free</Button>
                </section>
            </main>
            <footer className="py-6 px-10 text-center text-sm text-gray-600">
                © Uprank Design. All Rights Reserved.
            </footer>
        </div>
    );
}

function ReplyIcon(props: any) {
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
            <polyline points="9 17 4 12 9 7" />
            <path d="M20 18v-2a4 4 0 0 0-4-4H4" />
        </svg>
    );
}
