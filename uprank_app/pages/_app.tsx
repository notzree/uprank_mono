import "@/styles/globals.css";
import type { AppProps } from "next/app";
import { ClerkProvider } from "@clerk/nextjs";
import { Toaster } from "@/components/ui/toaster";
export default function App({
    Component,
    pageProps: { session, ...pageProps },
}: AppProps) {
    return (
        <ClerkProvider>
                <Component {...pageProps} />
            <Toaster/>
        </ClerkProvider>
    );
}
