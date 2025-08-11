export class MoviePage extends HTMLElement {
    connectedCallback() {
        const template = document.getElementById("template-home");
        const content = template.content.cloneNode(true);
        this.appendChild(content);
    }
}
customElements.define("movie-page", MoviePage)