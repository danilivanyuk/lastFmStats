<script lang='ts'>
    import { userNickname } from "../../store";
    import { mockedData as data } from "../../../static/mockData"
    let nickname = $userNickname
    // TODO: if there is no nickname return back to "/"
    const uniqueArtist = data.recenttracks.track.reduce((accumulator, currentValue) => (
    accumulator.includes(currentValue.artist["#text"]) ? accumulator : [...accumulator, currentValue.artist["#text"]]
  ), [] as {})
    console.log(uniqueArtist)
</script>
<div class="w-[80%] mx-auto">
    <h1>Hello {nickname}</h1>
    <!-- {uniqueArtist} -->
    {#each Object.entries(uniqueArtist) as [index, values]}
        <p>Artist: {values}</p>

    {/each}

</div>
