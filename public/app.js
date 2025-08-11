import { Homepage } from "./components/HomePage.js";
import { MovieDetailsPage } from "./components/MovieDetailsPage.js";
import { API } from "./services/API.js";
import './components/YouTubeEmbed.js'


window.app = {
    search: (event) => {
        event.preventDefault();
        const keywords = document.querySelector("input[type=search]").value;
        console.info(keywords)
    },
    api: API
}
window.addEventListener("DOMContentLoaded", () => {
    document.querySelector("main").appendChild(new Homepage())
    /*
    document.querySelector("main").appendChild(new MovieDetailsPage())
    window.app = {
        ...window.app,
        api: API
    }*/
})