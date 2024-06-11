// Locator is a class that makes it easy and reliable to locate elements on a page.


export class Locator {
    public timeout;
    constructor(timeout?: number) {
        this.timeout = timeout ?? 2000;
    }

    async locate<T>(
        selector: string,
        parent?: Element,

    ): Promise<Element | null> {
        return new Promise((resolve, reject) => {
            const intervalId = setInterval(() => {
                const element = parent
                    ? this.nestedSelector(selector, parent)
                    : this.querySelector(selector);
                if (element) {
                    if (
                        (element instanceof NodeList && element.length !== 0) ||
                        element instanceof Element
                    ) {
                        clearInterval(intervalId);
                        clearTimeout(timeoutId);
                        resolve(element);
                    }
                }
            }, 100);

            const timeoutId = setTimeout(() => {
                clearInterval(intervalId);
                console.log(
                    `Element with selector "${selector}" not found within ${this.timeout}ms`
                );
                resolve(null);
            }, this.timeout);
        });
    }

    async locateAll<T>(
        selector: string,
        parent?: Element,

    ): Promise<NodeList | null> {
        return new Promise((resolve, reject) => {
            const intervalId = setInterval(() => {
                const element = parent
                    ? this.nestedSelectorAll(selector, parent)
                    : this.querySelectorAll(selector);
                if (element) {
                    if (
                        (element instanceof NodeList && element.length !== 0) ||
                        element instanceof Element
                    ) {
                        clearInterval(intervalId);
                        clearTimeout(timeoutId);
                        resolve(element);
                    }
                }
            }, 100);

            const timeoutId = setTimeout(() => {
                clearInterval(intervalId);
                console.log(
                    `Element with selector "${selector}" not found within ${this.timeout}ms`
                );
                resolve(null);
            }, this.timeout);
        });
    }

    async locateMany<T extends Element | NodeList>(
        selectorQueryMap: SelectorQueryMap<T>,
    ): Promise<{ [key: string]: T }> {
        let found = {};
        let number_of_found_elements = 0;
        for (let [selector, [query, parent]] of Object.entries(
            selectorQueryMap
        )) {
            found[selector] = null;
        }

        return new Promise((resolve, reject) => {
            const intervalId = setInterval(() => {
                if (
                    number_of_found_elements ===
                    Object.keys(selectorQueryMap).length
                ) {
                    clearInterval(intervalId);
                    clearTimeout(timeoutId);
                    resolve(found);
                }

                for (let [selector, [query, parent]] of Object.entries(
                    selectorQueryMap
                )) {
                    if (found[selector]) {
                        continue;
                    }
                    const element = parent
                        ? query(selector, parent)
                        : query(selector);
                    if (
                        (element &&
                            element instanceof NodeList &&
                            element.length !== 0) ||
                        element instanceof Element
                    ) {
                        found[selector] = element;
                        number_of_found_elements += 1;
                    }
                }
            }, 100);

            const timeoutId = setTimeout(() => {
                clearInterval(intervalId);
                for (let [selector, [query, parent]] of Object.entries(
                    selectorQueryMap
                )) {
                    if (!found[selector]) {
                        console.warn(
                            `Element with selector "${selector}" not found within ${this.timeout}ms`
                        );
                    }
                }
                resolve(found);
            }, this.timeout);
        });
    }

    async waitForClose(selector: string): Promise<void> {
        return new Promise<void>((resolve) => {
            const intervalId = setInterval(() => {
                const closeButton = this.querySelector(selector);
                if (!closeButton) {
                    clearInterval(intervalId);
                    resolve();
                }
            }, 100);
        });
    }

    public querySelector(input: string): Element | null {
        return document.querySelector(input);
    }

    public querySelectorAll(input: string): NodeList | null {
        return document.querySelectorAll(input);
    }

    public nestedSelector(input: string, parent: Element): Element | null {
        return parent.querySelector(input);
    }

    public nestedSelectorAll(input: string, parent: Element): NodeList | null {
        return parent.querySelectorAll(input);
    }

}

type SelectorQueryMap<T> = {
    [key: string]: [(input: string, parent?: Element) => T, Element?];
};