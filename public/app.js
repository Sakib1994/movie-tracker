import { HomePage } from "./components/HomePage.js";
import { MovieDetailsPage } from "./components/MovieDetailsPage.js";
import { API } from "./services/API.js";
import './components/YouTubeEmbed.js'
import Router from "./services/Router.js";
import Store from "./services/Store.js";


window.app = {
    search: (event) => {
        event.preventDefault();
        const keywords = document.querySelector("input[type=search]").value;
        app.Router.go(`/movies?q=${keywords}`)
    },
    api: API,
    Router,
    Store,
    showError: (message = "There was an error loading the page", goToHome = false) => {
        document.querySelector("#alert-modal").showModal()
        document.querySelector("#alert-modal p").textContent = message
        if (goToHome) app.Router.go("/")
        return
    },
    closeError: () => document.querySelector("#alert-modal").close(),
    searchOrderChange: (order) => {
        const urlParams = new URLSearchParams(window.location.search);
        const q = urlParams.get("q");
        const genre = urlParams.get("genre") ?? "";
        app.Router.go(`/movies?q=${q}&order=${order}&genre=${genre}`);
    },
    searchFilterChange: (genre) => {
        const urlParams = new URLSearchParams(window.location.search);
        const q = urlParams.get("q");
        const order = urlParams.get("order") ?? "";
        app.Router.go(`/movies?q=${q}&order=${order}&genre=${genre}`);
    },
    register: async (event) => {
        event.preventDefault();
        let errors = [];
        const form = event.target;
        const name = form.querySelector('[id="register-name"]').value;
        const email = form.querySelector('[id="register-email"]').value;
        const password = form.querySelector('[id="register-password"]').value;
        const passwordConfirm = form.querySelector('[id="register-password-confirm"]').value;

        if (name.length < 4) errors.push("Enter your complete name");
        if (email.length < 8) errors.push("Enter your complete email");
        if (password.length < 6) errors.push("Enter a password with 6 characters");
        if (password != passwordConfirm) errors.push("Passwords don't match");
        if (errors.length == 0) {
            const response = await API.register(name, email, password);
            if (response.success) {
                app.Store.jwt = response.jwt
                app.Router.go("/account/")
            } else {
                app.showError(response.message, false);
            }
        } else {
            app.showError(errors.join(". "), false);
        }
    },
    login: async (event) => {
        event.preventDefault();
        let errors = [];
        const email = document.getElementById("login-email").value;
        const password = document.getElementById("login-password").value;

        if (email.length < 8) errors.push("Enter your complete email");
        if (password.length < 6) errors.push("Enter a password with 6 characters");
        if (errors.length == 0) {
            const response = await API.authenticate(email, password);
            if (response.success) {
                app.Store.jwt = response.jwt
                app.Router.go("/account/")
            } else {
                app.showError(response.message, false);
            }
        } else {
            app.showError(errors.join(". "), false);
        }
    },
    logout:()=>{
        Store.jwt = null;
        app.Router.go("/")
    }
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