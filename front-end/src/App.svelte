<script>
    let weather;
    let city = "";
    async function search(){
      fetch(`http://localhost:3000/query/${city}`).then(async(data)=>{
        weather = await data.json();
        weather = weather.data;
      })
    }
</script>
<svelte:head>
  <link rel="stylesheet" href="https://site-assets.fontawesome.com/releases/v6.4.0/css/all.css">
</svelte:head>
<main>
<div class="flex">
  <input bind:value={city}>
  <button on:click={search}><i class="fa-regular fa-magnifying-glass"></i></button>
</div>
{#if weather}
    <div class="weather">
      <h1>{city.toUpperCase()}</h1>
      <div class="container">
        <img width="100px" height="100px" src="http://openweathermap.org/img/w/{weather.weather[0].icon}.png">
        <b class="temp">{weather.main.temp}<b class="d">°C</b></b>
      </div>
    </div>
{/if}
</main>