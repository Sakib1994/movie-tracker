export const API = {
  baseURL: "/api",
  getTopMovies: async () => {
    return await API.fetch("movies/top");
  },
  getRandomMovies: async () => {
    return await API.fetch("movies/random");
  },
  getMovieById: async (id) => {
    return await API.fetch(`movies/${id}`);
  },
  searchMovies: async (q, order, genre) => {
    return await API.fetch(`movies/search`, { q, order, genre })
  },
  getGenres: async () => {
    return await API.fetch("genres");
  },
  getFavorites: async () => {
    try {
      return await API.fetch("account/favorites");
    } catch (e) {
      app.Router.go("/account/")
    }
  },
  getWatchlist: async () => {
    try {
      return await API.fetch("account/watchlist");
    } catch (e) {
      app.Router.go("/account/")
    }

  },
  fetch: async (ServiceURL, args) => {
    try {
      const queryString = args ? new URLSearchParams(args).toString() : ""
      const response = await fetch(`${API.baseURL}/${ServiceURL}?${queryString}`, {
        headers: {
          "Authorization": app.Store.jwt ? `Bearer ${app.Store.jwt}` : null
        }
      });
      const results = await response.json()
      return results;
    } catch (error) {
      console.error(error);
      app.showError();
    }

  },
  register: async (name, email, password) => {
    return await API.send("/account/register/", { name, email, password })
  },
  authenticate: async (email, password) => {
    return await API.send("/account/authenticate/", { email, password })
  },
  send: async (service, args) => {
    try {
      const response = await fetch(API.baseURL + service, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Authorization": app.Store.jwt ? `Bearer ${app.Store.jwt}` : null
        },
        body: JSON.stringify(args)
      });
      const result = await response.json();
      return result;
    } catch (e) {
      console.error(e);
      app.showError();
    }
  },
  saveToCollection: async (movie_id, collection) => {
    return await API.send("/account/save-to-collection/", {
      movie_id, collection
    });
  }
}