import { Homepage } from "./components/HomePage.js";
import { API } from "./services/API.js";

window.app = {
    search: (event) =>{
        event.preventDefault();
        const keywords = document.querySelector("input[type=search]").value;
        console.info(keywords)
    },
    api:API
}
window.addEventListener("DOMContentLoaded", () => {
    document.querySelector("main").appendChild(new Homepage())
    /*window.app = {
        ...window.app,
        api: API
    }*/
})