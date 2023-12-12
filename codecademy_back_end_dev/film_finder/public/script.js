const tmdbKey = '412b20c9c4c714cb2e219d935d252381';
const tmdbBaseUrl = 'https://api.themoviedb.org/3';
const playBtn = document.getElementById('playBtn');

const getGenres = async () => {
  const genreRequestEndpoint = '/genre/movie/list';
  const requestParams = `?api_key=${tmdbKey}`;
  const urlToFetch = `${tmdbBaseUrl}${genreRequestEndpoint}${requestParams}`;
  console.log(urlToFetch);

  try{
    const response = await fetch(urlToFetch);
    //console.log(response);

    if(response.ok){
      //console.log('Response evaluates to a truthy value!');
      const jsonResponse = await response.json();
      //console.log(jsonResponse);
      const genres = jsonResponse.genres;
      //console.log(genres);
      return genres;
    }
    else{
      throw new Error(`HTTP error! Status: ${response.status}`);
    }
  } catch (error){
     console.log('Fetch error: ', error);
  };

};

const getMovies = async () => {
  const selectedGenre = getSelectedGenre();
  const discoverMovieEndpoint = '/discover/movie';
  const requestParams = `?api_key=${tmdbKey}&with_genres=${selectedGenre}`;
  const urlToFetch = `${tmdbBaseUrl}${discoverMovieEndpoint}${requestParams}`;
  //console.log(urlToFetch);

  try {
    const response = await fetch(urlToFetch);
    if(response.ok) {
      const jsonResponse = await response.json();
      //console.log(jsonResponse);
      const movies = jsonResponse.results;
      console.log(movies);
      return movies;
    } else {
      throw new Error(`HTTP error! Status: ${response.status}`);
    }
  } catch (error) {
    console.log('Fetch error:', error);
  }
};

const getMovieInfo = async (movie) => {
  const movieId = movie.id;
  const movieEndpoint = `/movie/${movieId}`;
  const requestParams = `?api_key=${tmdbKey}`;
  const urlToFetch = `${tmdbBaseUrl}${movieEndpoint}${requestParams}`;
  console.log(urlToFetch);

  try {
    const response = await fetch(urlToFetch);
    if(response.ok) {
      const movieInfo = await response.json();
      return movieInfo;
    }
  } catch (error) {
    console.log('Fetch error:', error);
  }
};

// Gets a list of movies and ultimately displays the info of a random movie from the list
const showRandomMovie = async () => {
  const movieInfo = document.getElementById('movieInfo');
  if (movieInfo.childNodes.length > 0) {
    clearCurrentMovie();
  };
  const movies = await getMovies();
  const randomMovie = getRandomMovie(movies);
  const info = getMovieInfo(randomMovie);
  displayMovie(info);
};

getGenres().then(populateGenreDropdown);
playBtn.onclick = showRandomMovie;