import { API } from "../services/API.js";
import { MovieItemComponent } from "./MovieItem.js";

export class Homepage extends HTMLElement {
    async render() {
        const topMovies = await API.getTopMovies()
        const randomMovies = await API.getRandomMovies()

        const top10MoviesEl = document.getElementById("top-10")
        const randomMoviesEl = document.getElementById("random")
        renderMoviesInList(topMovies, top10MoviesEl.querySelector("ul"))
        renderMoviesInList(randomMovies, randomMoviesEl.querySelector("ul"))
        function renderMoviesInList(movies, ul) {
            ul.innerHTML = '';
            movies.forEach(movie => {
                const li = document.createElement("li")
                li.appendChild(new MovieItemComponent(movie));
                ul.appendChild(li);
            });
        }
    }
    connectedCallback() {
        const template = document.getElementById("template-home");
        const content = template.content.cloneNode(true);
        this.appendChild(content);

        this.render();
    }
}

customElements.define("home-page", Homepage)