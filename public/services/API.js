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
  fetch: async (ServiceURL, args) => {
    try {
      const queryString = args ? new URLSearchParams(args).toString() : ""
      const response = await fetch(`${API.baseURL}/${ServiceURL}?${queryString}`);
      const results = await response.json()
      return results;
    } catch (error) {
      console.error(error);
      // app.showError();
    }

  }
}