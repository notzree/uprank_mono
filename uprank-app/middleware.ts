import { authMiddleware, redirectToSignIn } from "@clerk/nextjs";
import { NextResponse } from "next/server";
export default authMiddleware({
    ignoredRoutes: ["/", "/sign-in", "/sign-up", "/api/public/(.*)"],
    apiRoutes: ["/api/private/(.*)"],
    afterAuth(auth, req, evt) {
        const res = NextResponse.next();
        console.log("path" + req.nextUrl.pathname);
        if (auth.isApiRoute) {
            //todo: Change this so that it only allows access to the api if the user is authenticated
            return res;
        }

        if (!auth.userId && !auth.isPublicRoute) {
            return redirectToSignIn({ returnBackUrl: req.url });
        }
        if (req.nextUrl.pathname.includes("/api")) {
            return res;
        }

        // Attempt to retrieve user-specific metadata
        if (auth.userId && auth.sessionClaims) {
            let unsafe_metadata: any;
            try {
                unsafe_metadata = auth.sessionClaims?.unsafe_metadata;
            } catch (e) {
                console.log("Error retrieving user metadata:", e);
                // Optionally handle the error more gracefully here
            }
            const completed_onboarding = unsafe_metadata?.completed_onboarding;
            // Redirect to onboarding if the user hasn't completed it and isn't currently on the onboarding form
            if (completed_onboarding === false) {
                if (
                    req.nextUrl.pathname.includes("/onboarding") ||
                    auth.isApiRoute
                ) {
                    return res;
                } else {
                    const onboarding_url = new URL(
                        "/client/onboarding/onboarding-form",
                        req.url
                    );
                    return NextResponse.redirect(onboarding_url);
                }
            } else if (
                completed_onboarding === true &&
                req.nextUrl.pathname.includes("/onboarding")
            ) {
                const onboarding_url = new URL("/client/dashboard", req.url);
                return NextResponse.redirect(onboarding_url);
            }
        }

        return res;
    },
});

export const config = {
    matcher: ["/((?!.+\\.[\\w]+$|_next).*)", "/", "/(api|trpc)/private/(.*)"],
};
