import { HomePage } from "./components/HomePage.js";
import { MovieDetailsPage } from "./components/MovieDetailsPage.js";
import { API } from "./services/API.js";
import './components/YouTubeEmbed.js'
import Router from "./services/Router.js";


window.app = {
    search: (event) => {
        event.preventDefault();
        const keywords = document.querySelector("input[type=search]").value;
        console.info(keywords)
    },
    api: API,
    Router
}
window.addEventListener("DOMContentLoaded", () => {
    // document.querySelector("main").appendChild(new HomePage())
    app.Router.init()
    /*
    document.querySelector("main").appendChild(new MovieDetailsPage())
    window.app = {
        ...window.app,
        api: API
    }*/
})