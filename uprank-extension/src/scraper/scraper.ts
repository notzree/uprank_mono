// Scraper is a class that scrapes some data and returns it as an array of objects
interface Scraper<T> {
    scrape(): Promise<T>;
}





