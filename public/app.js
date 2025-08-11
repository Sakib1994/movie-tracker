import { HomePage } from "./components/HomePage.js";
import { MovieDetailsPage } from "./components/MovieDetailsPage.js";
import { API } from "./services/API.js";
import './components/YouTubeEmbed.js'
import Router from "./services/Router.js";


window.app = {
    search: (event) => {
        event.preventDefault();
        const keywords = document.querySelector("input[type=search]").value;
        app.Router.go(`/movies?q=${keywords}`)
    },
    api: API,
    Router,
    showError: (message="There was an error loading the page", goToHome=true)=>{
        document.querySelector("#alert-modal").showModal()
        document.querySelector("#alert-modal p").textContent = message
        if (goToHome) app.Router.go("/")
        return
    },
    closeError:()=>document.querySelector("#alert-modal").close()
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