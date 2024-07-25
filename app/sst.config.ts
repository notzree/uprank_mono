/// <reference path="./.sst/platform/config.d.ts" />


export default $config({
    app(input) {
        return {
            name: "uprank-dev",
            removal: input?.stage === "production" ? "retain" : "remove",
            home: "aws",
        };
    },
    input: {
        console: {
            autodeploy: {
                target(event) {
                    if (event.type === "branch" && event.branch === "production" && event.action === "pushed") {
                        return { stage: "production" };
                    }
                    // if (event.type ==="branch" && event.branch ==="main" && event.action ==="pushed"){
                    //     return {stage: "dev"};
                    // }
                    return;
                }
            }
        }
    },
    async run() {
        new sst.aws.Nextjs("UprankApp");
    },
});
