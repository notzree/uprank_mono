/// <reference path="./.sst/platform/config.d.ts" />


const scret = new sst.Secret("base_backed_url")


export default $config({
    app(input) {
        return {
            name: "uprank-dev",
            removal: input?.stage === "production" ? "retain" : "remove",
            home: "aws",
        };
    },
    async run() {
        new sst.aws.Nextjs("MyWeb", {
            link: [scret]
        });
    },
});
