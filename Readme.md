# Home24 WebScraper Assignment
### Features Implemented
- What HTML version has the document?
- What is the page title?
- How many headings of what level are in the document?
- How many internal and external links are in the document? Are there any inaccessible links and how many?
    - Was not able to complete this. Currently, all the links on the webpage are displayed along with the link count.
- Does the page contain a login form?
### Improvements
- Differentiate internal and external links
- Reuse the colly.NewCollector object efficiently without initializing it for each method
- Display the response neatly on the UI (Currently the response is logged to the terminal)