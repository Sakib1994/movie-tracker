import { HomePage } from "../components/HomePage.js";
import { MovieDetailsPage } from "../components/MovieDetailsPage.js";
import { MoviePage } from "../components/MoviePage.js";

export const routes = [
    {
        path: "/",
        component: HomePage
    },
    {
        path: "/movies",
        component: MoviePage
    },
    {
        path: /\/movies\/(\d+)/,
        component: MovieDetailsPage
    },
//     {
//         path: "/account/register",
//         component: RegisterPage
//     },
//     {
//         path: "/account/login",
//         component: LoginPage
//     },     
//     {
//         path: "/account/",
//         component: AccountPage
//     },
//     {
//         path: "/account/favorites",
//         component: FavoritesPage
//     },	
// {
//         path: "/account/watchlist",
//         component: WatchlistPage
//     },	
]