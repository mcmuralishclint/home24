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
- Improve Error handling
- Cache already analyzed webpages
### Setup instructions
- Option 1
    ```
    go mod download
    go run main.go
    ```
    - goto localhost:3001 on a webbrowser
    - Input the webpage URL and click on submit
- Option 2
    ```
    docker pull mcmuralishclint/docker-home24:latest
    docker run -p 3001:3001 -d  mcmuralishclint/docker-home24
    ```
    - goto localhost:3001 on a webbrowser
    - Input the webpage URL and click on submit
